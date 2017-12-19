package admin

import "github.com/astaxie/beego"

var (
	AdminSignOutUri string
	AdminSignInUri  string
	AdminPrefix     string
)

type BaseController struct {
	beego.Controller
}

func init() {
	AdminSignInUri = beego.AppConfig.String("admin_sign_in_path")
	AdminSignOutUri = beego.AppConfig.String("admin_sign_out_path")
	AdminPrefix = beego.AppConfig.String("admin_prefix")
}

func (this *BaseController)isUserLogin() bool {
	return this.getUserId() != ""
}

func (this *BaseController) getUserId() string {
	u := this.GetSession(UserId)
	if u == nil {
		return ""
	}
	return u.(string)
}

func (this *BaseController) getUsername() string {
	name := this.GetSession(Username)
	if (name == nil) {
		return ""
	}
	return name.(string)
}