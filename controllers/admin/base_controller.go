package admin

import "github.com/astaxie/beego"


var adminPrefix, adminAuthUri string

type BaseController struct {
	beego.Controller
}

func init(){
	adminAuthUri = beego.AppConfig.String("admin_auth_path")
	adminPrefix = beego.AppConfig.String("admin_prefix")
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