package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"

	"github.com/tongyuehong1/golang-project/application/blog/filters"
	mysql "github.com/tongyuehong1/golang-project/application/blog/init"
	_ "github.com/tongyuehong1/golang-project/application/blog/routers"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:    []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LoginFilter)
	mysql.InitSql()
	beego.Run()
}
