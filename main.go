package main

import (
	_ "github.com/beego_study/routers"
	"github.com/astaxie/beego"
	"github.com/beego_study/models"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDb()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)


	beego.Run()
}
