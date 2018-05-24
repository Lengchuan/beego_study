package models

import (
	"time"
	"github.com/Unknwon/com"
	"os"
	"path"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	CreateTime      time.Time `orm:"index;auto_now_add"` //tag
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index;auto_now"`
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
	ReplytTime      time.Time `orm:"index;auto_now"`
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
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)

}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)

	if err == nil {
		return err
	}

	_, err = o.Insert(cate)

	if err != nil {
		return err
	}

	return nil
}

func GetAllCateGory() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)

	return cates, err
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

		return err
	}
	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)

	return err

}

func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := Topic{Title: title, Content: content, CreateTime: time.Now(), UpdateTime: time.Now()}

	_, err := o.Insert(&topic)

	return err
}

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")

	var err error
	if !isDesc {
		qs.OrderBy("-createTime")
		_, err = qs.All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func GetTopic(id string) (*Topic, error) {

	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	topic := new(Topic)
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tid, ).One(topic)

	if err != nil {

		return nil, err
	}

	//增加浏览次数
	topic.Views++
	_, err = o.Update(topic)

	return topic, err

}

func ModifyTopic(id, title, content string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	topic := Topic{Id: tid}
	if o.Read(&topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.UpdateTime = time.Now()
		_, err := o.Update(&topic)
		if err != nil {
			return err
		}
	}

	return nil
}
