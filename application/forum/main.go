package main

import (
	_ "github.com/tongyuehong1/golang-project/application/forum/routers"
	mysql "github.com/tongyuehong1/golang-project/application/forum/init"
	"github.com/astaxie/beego"
)

func main() {
	mysql.InitSql()
	beego.Run()
}

