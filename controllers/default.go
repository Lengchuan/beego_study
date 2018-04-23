package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beegotest"
	//c.Data["Email"] = "lishuijun1992@gmail.com"
	//c.TplName = "index.tpl"

	c.Ctx.WriteString("appname:" + beego.AppConfig.String("appname"))
	c.Ctx.WriteString("\nhttpport:" + beego.AppConfig.String("httpport"))
	c.Ctx.WriteString("\nrunmode:" + beego.AppConfig.String("runmode"))

	c.Ctx.WriteString("\n-----------------------")

	beego.Info("test1")
	beego.Trace("test1")

	beego.SetLevel(beego.LevelInformational)

	beego.Info("test2")
	beego.Trace("test2")
}
