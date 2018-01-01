package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/tongyuehong1/golang-project/application/blog/common"
	"github.com/tongyuehong1/golang-project/application/blog/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
	"fmt"
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
func (this *ArticleController) UpdateTitle() {
	var Title struct {
		Title       string
		ChangeTitle string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Title)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.UpdateTitle(Title.Title, Title.ChangeTitle)
		if err != nil {
			if err == common.ErrNotFound {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlNotFound}
			} else {
				logger.Logger.Error("UpdateTitle ", err)

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}
func (this *ArticleController) UpdateArticle() {
	var title struct {
		title         string
		ChangeArticle string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &title)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.UpdateArticle(title.title, title.ChangeArticle)
		if err != nil {
			logger.Logger.Error("UpdateArticle ", err)

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}
func (this *ArticleController) UpdateBrief() {
	var title struct {
		title       string
		changebrief string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &title)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.UpdateBrief(title.title, title.changebrief)
		if err != nil {
			logger.Logger.Error("Updatebrief ", err)

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

func (this *ArticleController) Delete() {
	var title struct {
		title string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &title)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.Delete(title.title)
		if err != nil {
			fmt.Println("ArticleServer.Delete err:", err)
			logger.Logger.Error("delete ", err)

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}
