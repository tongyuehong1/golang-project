package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/tongyuehong1/golang-project/application/forum/filters"
	mysql "github.com/tongyuehong1/golang-project/application/forum/init"
	_ "github.com/tongyuehong1/golang-project/application/forum/routers"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LoginFilter)
	mysql.InitSql()
	beego.Run()
}
