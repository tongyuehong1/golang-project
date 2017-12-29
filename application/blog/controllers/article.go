package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/tongyuehong1/golang-project/application/blog/models"
	"encoding/json"
	"github.com/tongyuehong1/golang-project/libs/logger"
	. "github.com/tongyuehong1/golang-project/application/blog/common"
)

type ArticleController struct {
	beego.Controller
}
func (this *ArticleController) Insert(){
	article := Article{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &article)
	if err != nil {
		logger.GlobalLogReporter.Error(err)
		this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrInvalidParam}
	} else {
		err := ArticleServer.Insert(article)
		if err != nil {
				logger.GlobalLogReporter.Error(err)

				this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrMysqlQuery}
			} else {

			this.Data["json"] = map[string]interface{}{RespKeyStatus: ErrSucceed}
		}
	}

	this.ServeJSON()
	}
