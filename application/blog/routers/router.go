package routers

import (
	"github.com/astaxie/beego"
	"github.com/tongyuehong1/golang-project/application/blog/controllers"
)

func init() {
	beego.Router("/insert", &controllers.ArticleController{}, "post:Insert")
}
