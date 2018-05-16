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
}
