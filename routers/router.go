package routers

import (
	"gensh.me/blog/controllers"
	"github.com/astaxie/beego"
	"gensh.me/blog/controllers/admin"
)

var adminPrefix string
var adminAuthUri string

func init() {
	adminAuthUri = beego.AppConfig.String("admin_auth_path")
	adminPrefix = beego.AppConfig.String("admin_prefix")

	initRouter()
	intiFilter()
}

func initRouter() {
	beego.Router("/", &controllers.HomeController{}, "get:Get")
	beego.Router("/settings", &controllers.HomeController{}, "get:Settings")

	beego.Router(adminPrefix, &admin.PanelController{}, "get:Get")
	beego.Router(adminAuthUri, &admin.AuthController{}, "get:SignIn;post:SignIn_POST")
}