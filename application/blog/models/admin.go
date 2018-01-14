package models

import (
	"github.com/astaxie/beego/orm"

	"github.com/tongyuehong1/golang-project/application/blog/utility"
)

type AdminServiceProvider struct {
}

var AdminServer *AdminServiceProvider

type Manager struct {
	ID   uint64 `orm:"column(id)"     json:"id"`
	Name string `orm:"column(name)"   json:"name"`
	Pass string `orm:"column(pass)"   json:"pass"`
}

func (this *AdminServiceProvider) Create(manager Manager) error {
	o := orm.NewOrm()

	// 哈希加密
	hash, err := utility.GenerateHash(manager.Pass)

	if err != nil {
		return err
	}
	password := string(hash)

	sql := "INSERT INTO article.admin(name,pass) VALUES (?,?)"
	values := []interface{}{manager.Name, password}
	raw := o.Raw(sql, values)
	_, err = raw.Exec()

	return err
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
