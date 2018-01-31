package routers

import (
	"github.com/astaxie/beego"
	"github.com/tongyuehong1/golang-project/application/blog/controllers"
)

func init() {
	beego.Router("/create", &controllers.AdminController{}, "post:Create")

	beego.Router("/login", &controllers.AdminController{}, "post:Login")
	beego.Router("/blog/insert", &controllers.ArticleController{}, "post:Insert")
	beego.Router("/blog/title/update", &controllers.ArticleController{}, "post:UpdateTitle")
	beego.Router("/blog/brief/update", &controllers.ArticleController{}, "post:UpdateBrief")
	beego.Router("/blog/article/update", &controllers.ArticleController{}, "post:UpdateArticle")
	beego.Router("/blog/delete", &controllers.ArticleController{}, "post:Delete")
	beego.Router("/blog/getall", &controllers.ArticleController{}, "post:GetAll")
	beego.Router("/blog/get", &controllers.ArticleController{}, "post:Get")

	beego.Router("/blog/update", &controllers.ArticleController{}, "post:Update")
}
