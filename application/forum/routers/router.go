package routers

import (
	"github.com/astaxie/beego"
	"github.com/tongyuehong1/golang-project/application/forum/controllers"
)

func init() {
	//beego.Router("/admin/create", &controllers.AdminController{}, "post:Create")
	//beego.Router("/admin/login", &controllers.AdminController{}, "post:Login")

	beego.Router("/user/create", &controllers.UserController{}, "post:Create")
	beego.Router("/user/login", &controllers.UserController{}, "post:Login")

	beego.Router("/article/insert", &controllers.ArticleController{}, "post:Insert")
	beego.Router("/article/change", &controllers.ArticleController{}, "post:Change")
	beego.Router("/article/recommend", &controllers.ArticleController{}, "post:Recommend")
	beego.Router("/article/getarticle", &controllers.ArticleController{}, "post:GetArticle")
	beego.Router("/article/searcharticle", &controllers.ArticleController{}, "post:SearchArticle")
	beego.Router("/article/allarticle", &controllers.ArticleController{}, "get:AllArticle")
	//beego.Router("/article/collect", &controllers.ArticleController{}, "post:Collect")
	beego.Router("/article/showcollection", &controllers.ArticleController{}, "post:ShowCollection")

	beego.Router("/comment/insert", &controllers.CommentController{}, "post:Insert")
	beego.Router("/comment/getcomment", &controllers.CommentController{}, "post:GetComment")
}
