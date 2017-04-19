package routers

import (
	"gensh.me/blog/controllers"
	"github.com/astaxie/beego"
	"gensh.me/blog/controllers/admin"
)

var adminStaticPrefix, adminApi string

func init() {
	adminApi = beego.AppConfig.String("admin_api")
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
	beego.Router(admin.AdminAuthUri, &admin.AuthController{}, "get,post:SignIn")
	if beego.BConfig.RunMode == beego.DEV {
		//can register a admin user in Dev Mode
		beego.Router(admin.AdminAuthUri, &admin.AuthController{}, "get,post:SignUp")
	}

	beego.Router(admin.AdminPrefix + "/", &admin.PanelController{}, "get:Get")

	beego.Router(admin.AdminPrefix + adminApi + "/post", &admin.PostsController{}, "get:List")
	beego.Router(admin.AdminPrefix + adminApi + "/post/add", &admin.PostsController{}, "post:Add")
	beego.Router(admin.AdminPrefix + adminApi + "/upload_token", &admin.PostsController{}, "get:UploadToken")
	//beego.Router(admin.AdminPrefix+"/post/delete", &admin.PostsController{}, "post:Del")
	beego.Router(admin.AdminPrefix + adminApi + "/categories", &admin.CategoryController{}, "get:Get")
	beego.Router(admin.AdminPrefix + adminApi + "/category/add", &admin.CategoryController{}, "post:CategoryAdd")
	beego.Router(admin.AdminPrefix + adminApi + "/sub_category/add", &admin.CategoryController{}, "post:SubCategoryAdd")
}