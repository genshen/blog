package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gensh.me/blog/controllers/admin"
)

func intiFilter() {
	var FilterUser = func(ctx *context.Context) {
		if ctx.Request.RequestURI != adminAuthUri {
			_, ok := ctx.Input.Session(admin.UserId).(string)
			if !ok {
				ctx.Redirect(302, adminAuthUri)
			}
		}
	}
	beego.InsertFilter(adminPrefix + "/*", beego.BeforeRouter, FilterUser)
}