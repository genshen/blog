package controllers

import (
	"github.com/astaxie/beego"
	"gensh.me/blog/components/context/settings"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	//c.Ctx.ResponseWriter.
	this.Data["Email"] = "astaxie@gmail.com"
	this.EnableRender = true
	this.TplName = "home/index.html"
}

func (this  *HomeController)Settings() {
	this.Data["json"] = &settings.S
	this.ServeJSON()
}