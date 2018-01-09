package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tongyuehong1/golang-project/application/forum/common"
	"github.com/tongyuehong1/golang-project/application/forum/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Create() {
	var user models.User

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)

	if err != nil {
		logger.Logger.Error("Unmarshal:", err)

		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.UserServer.Create(user)

		if err != nil {
			logger.Logger.Error("Unmarshal", err)

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {

			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}
func (this *UserController) Login() {
	var user models.User
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)

	if err != nil {
		logger.Logger.Error("Unmarshal:", err)

		this.Data["JSON"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		flag, err := models.UserServer.Login(user.Name, user.Pass)

		if err != nil {
			if err == orm.ErrNoRows {
				logger.Logger.Error("UserServer.Login", err)
				this.Data["JSON"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidUser}
			} else {
				logger.Logger.Error("UserServer.Login", err)
				this.Data["JSON"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
		} else {
			if !flag {
				logger.Logger.Debug("Wrong Pass!")

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrWrongPass}
			} else {
				this.SetSession(common.SessionUserID, user.Name)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		}
	}
	this.ServeJSON()
}
