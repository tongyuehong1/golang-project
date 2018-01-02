package models

import (
	"github.com/astaxie/beego/orm"

	"github.com/tongyuehong1/golang-project/application/blog/utility"
)

type AdminServiceProvider struct {
}

var AdminServer *AdminServiceProvider

type Manager struct {
	ID   uint64 `orm:"column(id)"`
	Name string
	Pass string
}

func (this *AdminServiceProvider) Login(name string, password string) (bool, error) {
	o := orm.NewOrm()
	var pass string

	err := o.Raw("SELECT pass FROM article.admin WHERE name=? LIMIT 1 LOCK IN SHARE MODE", name).QueryRow(&pass)
	if err != nil {
		return false, err
	} else if !utility.CompareHash([]byte(pass), password) {
		return false, nil
	}

	return true, nil
}
