package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/tongyuehong1/golang-project/application/forum/common"
	"github.com/tongyuehong1/golang-project/application/forum/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Insert() {
	comment := models.Comment{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &comment)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.CommentServer.Comment(comment)
		if err != nil {
			logger.Logger.Error("Insert ", err)

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}
func (this *CommentController) GetComment() {
	var Comment struct {
		ArtTitle string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Comment)

	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		show, err := models.CommentServer.GetComment(Comment.ArtTitle)
		if err != nil {
			logger.Logger.Error("GetComment ", err)

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: show}
		}
	}

	this.ServeJSON()
}
