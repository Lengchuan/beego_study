package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "home.html"

	this.Data["Website"] = "beegotest"
	this.Data["Email"] = "lishuijun1992@gmail.com"

	this.Data["TrueCond"] = true
	this.Data["FalseCond"] = false

	type u struct {
		Name string
		Age  int
		Sex  string
	}

	user := &u{
		"lengchuan",
		26,
		"male",
	}

	this.Data["user"] = user

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	this.Data["nums"] = nums

	this.Data["tplVar"] = "hey guys"

	//模板函数
	this.Data["html"] = "" +
		"</div>" +
		"html" +
		"</div>"

	this.Data["pipe"] = "<div> hello beego </div>"

}
