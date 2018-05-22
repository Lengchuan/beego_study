package routers

import (
	"github.com/astaxie/beego"
	"github.com/Lengchuan/beego_study/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/", &controllers.HelloWorldController{})

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.Logincontroller{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	//beego自动路由
	beego.AutoRouter(&controllers.TopicController{})
}
