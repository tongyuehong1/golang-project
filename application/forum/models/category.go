package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Category))
}

type CategoryServerProvider struct {
}

var CategoryServer *CategoryServerProvider

type Category struct {
	CategoryId   uint32 `orm:"column(id);pk"`
	ArticleId    uint64 `orm:"column(articleId)"`
	CategoryName string `orm:"column(categoryName)"`
	Status       uint8  `orm:"column(status)"`
}

func (this *CategoryServerProvider) GetCgyId(category string) (uint32, error) {
	var cgyid uint32
	o := orm.NewOrm()
	_, err := o.Raw("SELECT id FROM  WHERE categoryName=? AND status=?", category, true).QueryRows(&cgyid)
	return cgyid, err
}
