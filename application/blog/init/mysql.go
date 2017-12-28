package init

import (
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego"
)

func initSql() {
	//article := beego.AppConfig.String("mysql::article")
	orm.RegisterDataBase("default", "mysql", "root:111111@/127.0.0.1:3306?charset=utf8")
}
