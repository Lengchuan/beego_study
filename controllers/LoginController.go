package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type Logincontroller struct {
	beego.Controller
}

func (this *Logincontroller) Get() {
	this.TplName = "login.html"

}

func (this *Logincontroller) Post() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	name := this.Input().Get("name")
	pwd := this.Input().Get("passwd")
	configpasswd := beego.AppConfig.String("passwd")
	configname := beego.AppConfig.String("name")

	remerberme := this.Input().Get("remerberme") == "true"

	if configname == name && beego.AppConfig.String("passwd") == configpasswd {
		maxAge := 0
		if remerberme {

			maxAge = 1<<32 - 1
		}

		this.Ctx.SetCookie("name", name, maxAge, "/")
		this.Ctx.SetCookie("passwd", pwd, maxAge, "/")
	}

	this.Redirect("/", 301)

	return
}

func checkAccount(ctx *context.Context) bool {

	name := ctx.GetCookie("name")
	if name == "" {
		return false
	}

	pwd := ctx.GetCookie("passwd")
	if pwd == "" {
		return false
	}

	return beego.AppConfig.String("name") == name && beego.AppConfig.String("passwd") == pwd
}
