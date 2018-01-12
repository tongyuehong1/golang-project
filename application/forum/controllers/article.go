package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/tongyuehong1/golang-project/application/forum/common"
	"github.com/tongyuehong1/golang-project/application/forum/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
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

func (this *ArticleController) Change() {
	article := models.Article{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &article)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.Change(article)
		if err != nil {
			logger.Logger.Error("Change ", err)

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

// 推荐文章
func (this *ArticleController) Recommend() {
	var Title struct {
		Title string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Title)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.Recommend(Title.Title)
		if err != nil {
			logger.Logger.Error("Recommend ", err)

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

// 根据类别显示
func (this *ArticleController) GetArticle() {
	var Category struct {
		Category string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Category)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		show, err := models.ArticleServer.GetArticle(Category.Category)
		if err != nil {
			logger.Logger.Error("GetArticle ", err)

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: show}
		}
	}

	this.ServeJSON()
}

// 显示所有文章
func (this *ArticleController) AllArticle() {
	show, _ := models.ArticleServer.AllArticle()
	this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: show}

	this.ServeJSON()
}

// 显示搜索文章
func (this *ArticleController) SearchArticle() {
	var Title struct {
		Title string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Title)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		show, err := models.ArticleServer.SearchArticle(Title.Title)
		if err != nil {
			logger.Logger.Error("SearchArticle ", err)

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: show}
		}
	}

	this.ServeJSON()
}


// 收藏文章
func (this *ArticleController) Collect() {
	var User struct {
		User  string
		Title string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &User)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		userId, err := models.UserServer.GetUserId(User.User)
		if err != nil {
			logger.Logger.Error("getuserid", err)
		} else {
			error := models.ArticleServer.Collect(User.Title, userId)
			if error != nil {
				logger.Logger.Error("GetArticle ", error)

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			} else {

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		}
	}

	this.ServeJSON()

}

// 显示收藏文章
func (this *ArticleController) ShowCollection() {
	var User struct {
		User string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &User)
	if err != nil {
		logger.Logger.Error("Unmarshal ", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		userId, err := models.UserServer.GetUserId(User.User)
		if err != nil {
			logger.Logger.Error("getuserid", err)
		} else {
			articles, error := models.ArticleServer.ShowCollection(userId)
			if error != nil {
				logger.Logger.Error("GetArticle ", error)

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			} else {

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: articles}
			}
		}
	}

	this.ServeJSON()

}
