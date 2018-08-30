package routers

import (
	"github.com/astaxie/beego"
	"github.com/genshen/blog/controllers"
	"github.com/genshen/blog/controllers/admin"
)

var blogPagePrefix string
var blogApiPrefix string

func init() {
	blogPagePrefix = beego.AppConfig.DefaultString("blog_pages_prefix", "")
	blogApiPrefix = beego.AppConfig.String("blog_api_prefix")
	initRouter()
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

	beego.Router(blogApiPrefix+"/category", &controllers.PostsController{}, "get:Category")

	beego.Router(blogApiPrefix+"/detail/:id([0-9A-Fa-f]{24,24})", &controllers.PostsController{}, "get:Detail")
	beego.Router(blogApiPrefix+"/comment/add", &controllers.CommentController{}, "post:Add")
	beego.Router(blogApiPrefix+"/comments/:post_id([0-9A-Fa-f]{24,24})/:start([0-9]+)", &controllers.CommentController{}, "get:Load")
	beego.Router(blogApiPrefix+"/reply/add", &controllers.CommentController{}, "post:ReplyAdd")

	//admin router
	//beego.Router(admin.AdminSignOutUri, &admin.AuthController{}, "get:SignOut") // todo remove signout variable
	beego.Router(admin.AdminSignInUri, &admin.AuthController{}, "get,post:SignIn")
	if beego.BConfig.RunMode == beego.DEV {
		//can register a admin user in Dev Mode
		beego.Router(admin.AdminPagesPrefix+"/dev/sign_up", &admin.AuthController{}, "get,post:SignUp")
	}

	beego.Router(admin.AdminPagesPrefix+"/", &admin.PanelController{}, "get:Get")

	if beego.AppConfig.DefaultBool("storage::EnableQiNiuCloud", false) {
		beego.Router(admin.AdminApiPrefix+"/upload_token", &admin.StorageController{}, "get:QiNiuCloudStorageUploadToken")
	} else {
		beego.Router(admin.AdminApiPrefix+"/upload_token", &admin.StorageController{}, "get:LocalStorageUploadToken")
		// Note: no adminApi prefix !!
		beego.Router(beego.AppConfig.DefaultString("storage::LocalStorageDomain", "/images/:hash"),
			&admin.LocalStorageHashController{}, "get:LocalStorageResource")
		// Note:urlFor is ues in function storage_controller.go#initStorage
		beego.Router(admin.AdminApiPrefix+"/upload", &admin.StorageController{}, "post:LocalUpload")
		admin.InitStorage()
	}
	beego.Router(admin.AdminApiPrefix+"/article", &admin.PostsController{}, "get:List")
	beego.Router(admin.AdminApiPrefix+"/article/publish", &admin.PostsController{}, "post:Add")
	//beego.Router(admin.AdminApiPrefix+"/post/delete", &admin.PostsController{}, "post:Del")
	beego.Router(admin.AdminApiPrefix+"/categories", &admin.CategoryController{}, "get:Get")
	beego.Router(admin.AdminApiPrefix+"/category/add", &admin.CategoryController{}, "post:CategoryAdd")
	beego.Router(admin.AdminApiPrefix+"/sub_category/add", &admin.CategoryController{}, "post:SubCategoryAdd")
}
