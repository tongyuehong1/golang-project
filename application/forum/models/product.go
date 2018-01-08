package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Product))
}

type ProductServiceProvider struct {
}

var ProductServer *UserServiceProvider

type Product struct{
	Id         uint64 		`orm:"column(id);pk"  json:"id"`
	Title  	   string 		`orm:"column(title)"    json:"title"`
	Classes    string		`orm:"column(classes)"    json:"classes"`
    Content    string		`orm:"column(content)"    json:"content"`
    Author     string		`orm:"column(author)"    json:"author"`
    Created    time.Time    `orm:"column(created)"    json:"created"`
    Brief      string		`orm:"column(brief)"    json:"brief"`
    Comment    string		`orm:"column(comment)"    json:"comment"`
    Status     bool			`orm:"column(status)"    json:"status"`
}

func (this *ProductServiceProvider) Insert(product Product){
	o := orm.NewOrm()

}