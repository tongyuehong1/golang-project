package init

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/tongyuehong1/golang-project/application/blog/common"
)

func InitSql() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:111111@tcp(127.0.0.1:3306)/forum?charset=utf8")
	orm.RegisterDataBase(DBForum, "mysql", "root:111111@tcp(127.0.0.1:3306)/forum?charset=utf8")
}
