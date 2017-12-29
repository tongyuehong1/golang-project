package controllers

import (
	"github.com/astaxie/beego"
	"github.com/tongyuehong1/golang-project/application/sample/models"
)

type Sample struct {
	beego.Controller
}
func (o *Sample) Hello() {
	o.Data["json"] = map[string]string{"content": models.Hello()}
	o.ServeJSON()
}
