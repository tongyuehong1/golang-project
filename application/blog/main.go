package main

import (
	_ "github.com/tongyuehong1/golang-project/application/blog/routers"
	"github.com/astaxie/beego"
	mysql "github.com/tongyuehong1/golang-project/application/blog/init"
)

func main() {
	mysql.InitSql()
	beego.Run()
}
