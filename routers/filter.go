package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gensh.me/blog/controllers/admin"
)

func intiFilter() {
	var FilterUser = func(ctx *context.Context) {
		if ctx.Request.RequestURI != adminAuthUri {
			is_auth, ok := ctx.Input.Session(admin.UserId).(bool)
			if !ok || !is_auth {
				ctx.Redirect(302, adminAuthUri)
			}
		}
	}
	beego.InsertFilter(adminPrefix + "/*", beego.BeforeRouter, FilterUser)
}