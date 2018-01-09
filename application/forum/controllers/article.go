package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/tongyuehong1/golang-project/application/forum/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
	"github.com/tongyuehong1/golang-project/application/forum/common"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Insert() {
	article := models.Article{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &article)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.Insert(article)
		if err != nil {
			logger.Logger.Error("Insert ", err)

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}