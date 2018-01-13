package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/tongyuehong1/golang-project/application/forum/utility"
	"fmt"
)

func init() {
	orm.RegisterModel(new(User),new(UserExtra))
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
type UserExtra struct {
	Id     uint64	`orm:"column(id);pk"  json:"id"`
	UserID uint64   `orm:"column(userid);"  json:"userid"`
	Key    string	`orm:"column(key);"  json:"key"`
	Value  string	`orm:"column(value);"  json:"value"`
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
	_, err = o.Insert(newuser)
	fmt.Println("er",err)
	return err
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

func (this *UserServiceProvider) GetUserId(name string) (uint64, error) {
	o := orm.NewOrm()
	var userId uint64

	err := o.Raw("SELECT id FROM forum.user WHERE name = ?", name).QueryRow(&userId)

	return userId, err
}
