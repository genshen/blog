package routers

import (
	"gensh.me/blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{},"get:Get")
    beego.Router("/settings", &controllers.HomeController{},"get:Settings")
}