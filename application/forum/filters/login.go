package filters

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/tongyuehong1/golang-project/application/forum/common"
)

func LoginFilter(ctx *context.Context) {
	if _, ok := MapFilter[ctx.Request.RequestURI]; !ok {
		userID := ctx.Input.CruSession.Get(common.SessionUserID)

		fmt.Println("jl", userID)
		if userID == nil {
			ctx.Output.JSON(map[string]interface{}{common.RespKeyStatus: common.ErrLoginRequired}, false, false)
		}
	}
}
