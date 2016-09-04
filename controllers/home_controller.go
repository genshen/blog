package controllers

import (
	"gensh.me/blog/components/context/settings"
	"gensh.me/blog/components/auth"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.Data["Email"] = "astaxie@gmail.com"
	this.EnableRender = true
	this.TplName = "home/index.html"
}

type SettingData struct {
	IsAuth   bool               `json:"is_auth"`
	User     *auth.User	    `json:"user"`
	Settings settings.Setting  `json:"settings"`
}

func (this  *HomeController)Settings() {
	settingData := SettingData{IsAuth:false, Settings:settings.GetSettings()}
	if this.HasAuth() {
		u := this.GetUserData()
		settingData.User = &u
		settingData.IsAuth = true
	}
	this.Data["json"] = &settingData
	this.ServeJSON()
}