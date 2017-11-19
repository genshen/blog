package controllers

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/genshen/blog/components/auth"
)

const (
	UserData = "user_data"
	IsAuth = "is_auth"
)

type BaseController struct {
	beego.Controller
}

func baseControllerInit(){
	gob.Register(auth.User{})
}

func (this *BaseController)HasAuth() bool {
	is_auth := this.GetSession(IsAuth)
	if is_auth == nil {
		return false
	}
	return is_auth.(bool)
}

func (this *BaseController)GetUserData() auth.User {
	user, ok := this.GetSession(UserData).(auth.User)
	if !ok {
		return auth.User{}
	}
	return user
}