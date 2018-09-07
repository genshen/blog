package controllers

import (
	"github.com/astaxie/beego"
	"github.com/genshen/blog/components/context/category"
	"github.com/genshen/blog/components/context/settings"
	"github.com/genshen/blog/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "home/index.html"
}

type SettingData struct {
	//IsAuth     bool                `json:"is_auth"`
	//User       *auth.User          `json:"user"`
	Categories []models.Category `json:"categories"`
	Settings   *settings.Setting `json:"settings"`
}

func (this *HomeController) Settings() {
	settingData := SettingData{Settings: &settings.S}
	settingData.Categories = category.GetCategories()
	this.Data["json"] = &settingData
	this.ServeJSON()
}
