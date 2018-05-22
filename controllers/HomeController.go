package controllers

import (
	"github.com/astaxie/beego"
	"github.com/Lengchuan/beego_study/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "home.html"
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	var err error
	this.Data["Topics"], err = models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}
}
