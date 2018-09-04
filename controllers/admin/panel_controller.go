package admin

import "github.com/genshen/blog/components/utils"

type PanelController struct {
	BaseAuthController
}

func (p *PanelController) OnUnAuth() {
	p.Redirect(utils.CustomConfig.Api.AdminSignInPath+"?next="+p.Ctx.Request.RequestURI, 302)
}

func (this *PanelController) Get() {
	this.Data["Email"] = "astaxie@gmail.com"
	this.EnableRender = true
	this.TplName = "admin/panel/index.html"
}
