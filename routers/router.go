package routers

import (
	"github.com/astaxie/beego"
	"github.com/genshen/blog/controllers"
)

var blogPagePrefix string
var blogApiPrefix string

func init() {
	blogPagePrefix = beego.AppConfig.DefaultString("blog_pages_prefix", "")
	blogApiPrefix = beego.AppConfig.String("blog_api_prefix")
	initRouter()
	initAdminRoute()
	intiFilter()
	// in dev mode, we ignore xsrf.
	if beego.BConfig.RunMode == beego.DEV {
		beego.BConfig.WebConfig.EnableXSRF = false
	}
}

func initRouter() {
	beego.Router(blogPagePrefix+"/", &controllers.HomeController{}, "get:Get")
	beego.Router(blogPagePrefix+"/detail/:id([0-9A-Fa-f]{24,24})", &controllers.HomeController{}, "get:Get")
	beego.Router(blogPagePrefix+"/settings", &controllers.HomeController{}, "get:Settings")
	beego.Router(blogPagePrefix+"/auth/callback", &controllers.AuthController{}, "get:Callback")

	beego.Router(blogApiPrefix+"/categories", &controllers.CategoryController{}, "get:Get")
	beego.Router(blogApiPrefix+"/list", &controllers.PostsController{}, "get:List")

	beego.Router(blogApiPrefix+"/detail/:id([0-9A-Fa-f]{24,24})", &controllers.PostsController{}, "get:Detail")
	beego.Router(blogApiPrefix+"/comment/add", &controllers.CommentController{}, "post:Add")
	beego.Router(blogApiPrefix+"/comments/:post_id([0-9A-Fa-f]{24,24})/:start([0-9]+)", &controllers.CommentController{}, "get:Load")
	beego.Router(blogApiPrefix+"/reply/add", &controllers.CommentController{}, "post:ReplyAdd")
}
