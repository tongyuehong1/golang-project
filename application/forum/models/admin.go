package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/tongyuehong1/golang-project/application/forum/utility"
)

func init() {
	orm.RegisterModel(new(Admin))
}

type AdminServiceProvider struct {
}

var AdminServer *AdminServiceProvider

type Admin struct {
	ID   uint64 `orm:"column(id)"`
	Name string `orm:"column(name)"   json:"name"       valid:"Required`
	Pass string `orm:"column(pass)"   json:"pass"		valid:"MaxSize(16);MixSize(6)"`
}

// 添加管理员用户
func (this *AdminServiceProvider) Create(admin Admin) error {
	o := orm.NewOrm()

	// 哈希加密
	hash, err := utility.GenerateHash(admin.Pass)

	if err != nil {
		return err
	}
	password := string(hash)

	sql := "INSERT INTO forum.admin(name,pass) VALUES (?,?)"
	values := []interface{}{admin.Name, password}
	raw := o.Raw(sql, values)
	_, err = raw.Exec()

	return err
}

// 管理员用户登录
func (this *AdminServiceProvider) Login(name string, password string) (bool, error) {
	o := orm.NewOrm()
	var pass string

	err := o.Raw("SELECT pass FROM forum.admin WHERE name=? LIMIT 1 LOCK IN SHARE MODE", name).QueryRow(&pass)

	if err != nil {
		return false, err
	} else if !utility.CompareHash([]byte(pass), password) {
		return false, nil
	}

	return true, nil
}
