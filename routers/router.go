package routers

import (
	"github.com/astaxie/beego"
	"github.com/genshen/blog/controllers"
	"github.com/genshen/blog/controllers/admin"
)

var adminStaticPrefix, adminApi string

func init() {
	adminApi = admin.AdminPrefix + beego.AppConfig.String("admin_api")
	adminStaticPrefix = beego.AppConfig.String("admin_static_prefix")

	initRouter()
	intiFilter()
}

func initRouter() {
	beego.Router("/", &controllers.HomeController{}, "get:Get")
	beego.Router("/detail/:id([0-9A-Fa-f]{24,24})", &controllers.HomeController{}, "get:Get")
	beego.Router("/settings", &controllers.HomeController{}, "get:Settings")

	beego.Router("/at/category", &controllers.PostsController{}, "get:Category")

	beego.Router("/at/detail/:id([0-9A-Fa-f]{24,24})", &controllers.PostsController{}, "get:Detail")
	beego.Router("/at/comment/add", &controllers.CommentController{}, "post:Add")
	beego.Router("/at/comments/:post_id([0-9A-Fa-f]{24,24})/:start([0-9]+)", &controllers.CommentController{}, "get:Load")
	beego.Router("/at/reply/add", &controllers.CommentController{}, "post:ReplyAdd")

	beego.Router("/auth/callback", &controllers.AuthController{}, "get:Callback")

	//admin router
	beego.Router(admin.AdminSignOutUri, &admin.AuthController{}, "get:SignOut")
	beego.Router(admin.AdminSignInUri, &admin.AuthController{}, "get,post:SignIn")
	if beego.BConfig.RunMode == beego.DEV {
		//can register a admin user in Dev Mode
		beego.Router(adminApi+"/dev/sign_up", &admin.AuthController{}, "get,post:SignUp")
	}

	beego.Router(admin.AdminPrefix+"/", &admin.PanelController{}, "get:Get")

	if beego.AppConfig.DefaultBool("storage::EnableQiNiuCloud", false) {
		beego.Router(adminApi+"/upload_token", &admin.StorageController{}, "get:QiNiuCloudStorageUploadToken")
	} else {
		beego.Router(adminApi+"/upload_token", &admin.StorageController{}, "get:LocalStorageUploadToken")
		// Note: no adminApi prefix !!
		beego.Router(beego.AppConfig.DefaultString("storage::LocalStorageDomain", "/images/:hash"),
			&admin.StorageController{}, "get:LocalStorageResource")
		// Note:urlFor is ues in function storage_controller.go#initStorage
		beego.Router(adminApi+"/upload", &admin.StorageController{}, "post:LocalUpload")
		admin.InitStorage()
	}
	beego.Router(adminApi+"/article", &admin.PostsController{}, "get:List")
	beego.Router(adminApi+"/article/publish", &admin.PostsController{}, "post:Add")
	//beego.Router(admin.AdminPrefix+"/post/delete", &admin.PostsController{}, "post:Del")
	beego.Router(adminApi+"/categories", &admin.CategoryController{}, "get:Get")
	beego.Router(adminApi+"/category/add", &admin.CategoryController{}, "post:CategoryAdd")
	beego.Router(adminApi+"/sub_category/add", &admin.CategoryController{}, "post:SubCategoryAdd")
}
