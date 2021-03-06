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
	this.Data["Topics"], err = models.GetAllTopics(false)
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
	tid := this.Input().Get("id")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, content)

	} else {
		err = models.ModifyTopic(tid, title, content)
	}
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)

}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"

	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = topic.Id
}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"
	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = topic.Id
}
