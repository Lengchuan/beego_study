package main

import (
	_ "github.com/Lengchuan/beego_study/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/Lengchuan/beego_study/models"
)

func init() {
	models.RegisterDb()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)


	beego.Run()
}
