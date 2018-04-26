package models

import (
	"time"
	"github.com/Unknwon/com"
	"os"
	"path"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type CateGory struct {
	Id              int64
	Title           string
	CreateTime      time.Time `orm:"index"` //tag
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string    `orm:"size(5000)"`
	Attachment      string
	CreateTime      time.Time `orm:"index"`
	UpdateTime      time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	ReplytTime      time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId string    `orm:"index"`
}

func RegisterDb() {
	//创建文件夹
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	//创建模型
	orm.RegisterModel(new(CateGory), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)

}
