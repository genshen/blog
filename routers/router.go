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
	beego.Router("/settings", &controllers.HomeController{}, "get:Settings")

	//admin router
	beego.Router(admin.AdminSignOutUri, &admin.AuthController{}, "get:SignOut")
	beego.Router(admin.AdminAuthUri, &admin.AuthController{}, "get,post:SignIn")

	beego.Router(admin.AdminPrefix, &admin.PanelController{}, "get:Get")

	beego.Router(admin.AdminPrefix + adminApi + "/post", &admin.PostsController{}, "get:List")
	beego.Router(admin.AdminPrefix + adminApi + "/post/add", &admin.PostsController{}, "post:Add")
	//beego.Router(admin.AdminPrefix+"/post/delete", &admin.PostsController{}, "post:Del")
}