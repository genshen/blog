package controllers

import (
	"github.com/genshen/blog/components/context/settings"
	"github.com/genshen/blog/components/auth"
	"github.com/genshen/blog/models"
	"github.com/genshen/blog/components/context/category"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.TplName = "home/index.html"
}

type SettingData struct {
	IsAuth     bool                `json:"is_auth"`
	User       *auth.User          `json:"user"`
	Categories []models.Category   `json:"categories"`
	Settings   *settings.Setting   `json:"settings"`
}

func (this  *HomeController)Settings() {
	settingData := SettingData{IsAuth:false, Settings:&settings.S}
	settingData.Categories = category.GetCategories()
	if this.HasAuth() {
		u := this.GetUserData()
		settingData.User = &u
		settingData.IsAuth = true
	}
	this.Data["json"] = &settingData
	this.ServeJSON()
}