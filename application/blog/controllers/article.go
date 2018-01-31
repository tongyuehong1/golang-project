package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/tongyuehong1/golang-project/application/blog/common"
	"github.com/tongyuehong1/golang-project/application/blog/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Insert() {
	article := models.Article{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &article)
	logger.Logger.Info("articleeee:", article)
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
	var Title struct {
		Title         string
		ChangeArticle string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Title)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.UpdateArticle(Title.Title, Title.ChangeArticle)
		if err != nil {
			if err == common.ErrNotFound {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlNotFound}
			} else {
				logger.Logger.Error("UpdateArticle ", err)

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

func (this *ArticleController) UpdateBrief() {
	var Title struct {
		Title       string
		ChangeBrief string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Title)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.UpdateBrief(Title.Title, Title.ChangeBrief)
		if err != nil {
			if err == common.ErrNotFound {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlNotFound}
			} else {
				logger.Logger.Error("UpdateArticle ", err)

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

func (this *ArticleController) Delete() {
	var Title struct {
		Title string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Title)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.Delete(Title.Title)
		if err != nil {
			if err == common.ErrNotFound {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlNotFound}
			} else {
				logger.Logger.Error("Delete ", err)

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

func (this *ArticleController) GetAll() {
	var Classes struct {
		Classes string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Classes)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		show, err := models.ArticleServer.GetAll(Classes.Classes)
		if err != nil {
			if err == common.ErrNotFound {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlNotFound}
			} else {
				logger.Logger.Error("GetAll ", err)

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: show}
		}
	}

	this.ServeJSON()
}
func (this *ArticleController) Get() {
	var Title struct {
		Title string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Title)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		show, err := models.ArticleServer.Get(Title.Title)
		if err != nil {
			if err == common.ErrNotFound {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlNotFound}
			} else {
				logger.Logger.Error("Get ", err)

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: show}
		}
	}

	this.ServeJSON()
}

func (this *ArticleController) Update() {
	var Title struct {
		Title   string
		Article models.Article
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Title)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.Update(Title.Title, Title.Article)
		if err != nil {
			if err == common.ErrNotFound {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlNotFound}
			} else {
				logger.Logger.Error("Update ", err)

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}
