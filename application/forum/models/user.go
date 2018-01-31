package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/tongyuehong1/golang-project/application/forum/common"
	"github.com/tongyuehong1/golang-project/application/forum/utility"
	"github.com/tongyuehong1/golang-project/libs/logger"
	"time"
)

func init() {
	orm.RegisterModel(new(User), new(UserExtra))
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
	Id     uint64 `orm:"column(id);pk"  json:"id"`
	UserID uint64 `orm:"column(userid);"  json:"userid"`
	Key    string `orm:"column(key);"  json:"key"`
	Value  string `orm:"column(value);"  json:"value"`
}

func (u UserExtra) TableName() string {
	return "userextra"
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
	id, err := o.Insert(newuser)
	if err != nil {
		return err
	}
	err = this.InsertTime(id)
	if err != nil {
		return err
	}
	fmt.Println("err", err)
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

func (this *UserServiceProvider) GetLastTime(userid uint64) (time.Time, error) {
	o := orm.NewOrm()
	var lasttime time.Time
	o.Using("forum")
	user := UserExtra{UserID: userid, Key: common.KeyLastInsert}
	err := o.Read(&user, "UserID", "Key")
	if err != nil {
		logger.Logger.Error("GetLastTime:", err)
	} else {
		t := user.Value
		lasttime, _ = time.Parse("2006-01-02 15:04:05", t)
		return lasttime, err
	}
	return lasttime, err
}

func (this *UserServiceProvider) InsertTime(userid int64) error {
	o := orm.NewOrm()
	lasttime := "2006-01-02 15:04:05"
	sql := "INSERT INTO forum.userextra(userid,`key`,value) VALUES (?,?,?)"
	values := []interface{}{userid, common.KeyLastInsert, lasttime}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()

	return err
}
