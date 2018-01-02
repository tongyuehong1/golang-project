package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"

	"github.com/tongyuehong1/golang-project/application/blog/filters"
	mysql "github.com/tongyuehong1/golang-project/application/blog/init"
	_ "github.com/tongyuehong1/golang-project/application/blog/routers"
)

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LoginFilter)
	mysql.InitSql()
	beego.Run()
}
