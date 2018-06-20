package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func intiFilter() {
	//beego.InsertFilter(admin.AdminPagesPrefix+"/*", beego.BeforeRouter, FilterPagesAuth)
	//beego.InsertFilter(admin.AdminApiPrefix+"/*", beego.BeforeRouter, FilterApiAuth)
	// in dev mode, we ignore access control.
	if beego.BConfig.RunMode == beego.DEV {
		beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
			ctx.Output.Header("Access-Control-Allow-Origin", "*")
		})
	}
}
