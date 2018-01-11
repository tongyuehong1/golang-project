package routers

import (
	"github.com/astaxie/beego"
	"github.com/tongyuehong1/golang-project/application/forum/controllers"
)

func init() {
	beego.Router("/user/create", &controllers.UserController{}, "post:Create")
	beego.Router("/user/login", &controllers.UserController{}, "post:Login")

	beego.Router("/article/insert", &controllers.ArticleController{}, "post:Insert")
	beego.Router("/article/change", &controllers.ArticleController{}, "post:Change")
	beego.Router("/article/recommend", &controllers.ArticleController{}, "post:Recommend")
	beego.Router("/article/getarticle", &controllers.ArticleController{}, "post:GetArticle")
	beego.Router("/article/allarticle", &controllers.ArticleController{}, "get:AllArticle")

	beego.Router("/comment/insert", &controllers.CommentController{}, "post:Insert")
	beego.Router("/comment/getcomment", &controllers.CommentController{}, "post:GetComment")
}
