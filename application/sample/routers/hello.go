package routers

import (
	"github.com/astaxie/beego"
	"github.com/tongyuehong1/golang-project/application/sample/controllers"
)

func init() {
	beego.Router("/", &controllers.Sample{}, "get:Hello")
}
