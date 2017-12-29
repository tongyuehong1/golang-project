package init

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/tongyuehong1/golang-project/application/blog/common"
)

func initSql() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:111111@/127.0.0.1:3306?charset=utf8")
	orm.RegisterDataBase(DBArticle, "mysql", "root:111111@/127.0.0.1:3306?charset=utf8")
}
