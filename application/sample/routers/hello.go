package routers

import (
	"github.com/tongyuehong1/golang-project/application/sample/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.Sample{},"get:Hello")
}
