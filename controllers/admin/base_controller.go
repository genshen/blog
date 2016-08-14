package admin

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
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