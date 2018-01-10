package main

import (
	"github.com/astaxie/beego"
	mysql "github.com/tongyuehong1/golang-project/application/forum/init"
	_ "github.com/tongyuehong1/golang-project/application/forum/routers"
)

func main() {
	mysql.InitSql()
	beego.Run()
}
