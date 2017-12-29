package main

import (
	_ "github.com/tongyuehong1/golang-project/application/blog/routers"
	"github.com/astaxie/beego"
	"github.com/tongyuehong1/golang-project/application/blog/controllers"
)

func main() {
	beego.Router("/insert", &controllers.ArticleController{}, "post: Insert")
}

