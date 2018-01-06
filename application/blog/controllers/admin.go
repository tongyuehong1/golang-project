package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/tongyuehong1/golang-project/application/blog/common"
	"github.com/tongyuehong1/golang-project/application/blog/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Login() {
	var admin models.Manager

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &admin)

	if err != nil {
		logger.Logger.Error("Unmarshal ", err)

		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		flag, err := models.AdminServer.Login(admin.Name, admin.Pass)

		if err != nil {
			if err == orm.ErrNoRows {
				logger.Logger.Error("Unmarshal ", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidUser}
			} else {
				logger.Logger.Error("Unmarshal ", err)

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
		} else {
			if !flag {
				logger.Logger.Debug("Wrong Pass!")

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrWrongPass}
			} else {
				this.SetSession(common.SessionUserID, admin.Name)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		}
	}

	this.ServeJSON()
}