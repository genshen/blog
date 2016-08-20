package controllers

import (
	"gensh.me/blog/components/context/settings"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	//this.ServeStaticView()
	this.Data["Email"] = "astaxie@gmail.com"
	this.EnableRender = true
	this.TplName = "home/index.html"
}

func (this  *HomeController)Settings() {
	this.Data["json"] = &settings.S
	this.ServeJSON()
}