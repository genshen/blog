package routers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gensh.me/blog/controllers/admin"
)

func intiFilter() {
	var FilterUser = func(ctx *context.Context) {
		var baseUri string
		index := strings.IndexByte(ctx.Request.RequestURI, '?')
		if index == -1 {
			index = len(ctx.Request.RequestURI)
		}
		baseUri = ctx.Request.RequestURI[0:index]
		if !strings.HasPrefix(baseUri, admin.AdminAuthUri) && baseUri != admin.AdminSignOutUri {
			_, ok := ctx.Input.Session(admin.UserId).(string)
			if !ok {
				ctx.Redirect(302, admin.AdminAuthUri + "?next=" + ctx.Request.RequestURI)
			}
		}
	}
	beego.InsertFilter(admin.AdminPrefix + "/*", beego.BeforeRouter, FilterUser)
}