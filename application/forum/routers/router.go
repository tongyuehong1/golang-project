package routers

import (
	"github.com/tongyuehong1/golang-project/application/forum/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/user/create", &controllers.UserController{}, "post:Create")
    beego.Router("/user/login", &controllers.UserController{},"post:Login")
	beego.Router("/article/insert", &controllers.ArticleController{},"post:Insert")
}
