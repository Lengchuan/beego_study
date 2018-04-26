package routers

import (
	"github.com/beego_study/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/", &controllers.HelloWorldController{})
	beego.Router("/", &controllers.HomeController{})
}
