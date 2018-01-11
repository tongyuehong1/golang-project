package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/tongyuehong1/golang-project/application/blog/utility"
)

func init() {
	orm.RegisterModel(new(User))
}

type UserServiceProvider struct {
}

var UserServer *UserServiceProvider

type User struct {
	UserId uint64 `orm:"column(id);pk"  json:"id"`
	Name   string `orm:"column(name)"   json:"name"`
	Pass   string `orm:"column(pass)"   json:"pass"`
	Phone  string `orm:"column(phone)"  json:"phone"`
}

// 用户相关信息
type UserExtraInfo struct {
	UserID uint64	`orm:"column(userid);pk"  json:"userid"`
	Key    string
	Value  string
}

func (this *UserServiceProvider) Create(user User) error {
	o := orm.NewOrm()
	hash, err := utility.GenerateHash(user.Pass)

	if err != nil {
		return err
	}
	password := string(hash)
	o.Using("User")
	newuser := new(User)
	newuser.Name = user.Name
	newuser.Pass = password
	newuser.Phone = user.Phone
	_, error := o.Insert(&newuser)
	return error
}

func (this *UserServiceProvider) Login(name string, pass string) (bool, error) {
	o := orm.NewOrm()
	o.Using("User")
	user := User{Name: name}
	err := o.Read(&user, "name")
	if err != nil {
		return false, err
	} else if !utility.CompareHash([]byte(user.Pass), pass) {
		return false, err
	}
	return true, nil
}
