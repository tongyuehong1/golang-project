package routers

import (
	"github.com/tongyuehong1/golang-project/application/blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
