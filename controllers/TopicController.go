package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/Lengchuan/beego_study/models"
	"github.com/Lengchuan/beego_study/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {

	this.TplName = "topic.html"
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	var err error
	this.Data["Topics"], err = models.GetAllTopics()
	if err != nil {
		beego.Error(err)
	}
}

func (this *TopicController) Add() {

	this.TplName = "topic_add.html"
}

func (this *TopicController) Post() {

	//身份验证
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")

	var err error
	err = models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)

}
