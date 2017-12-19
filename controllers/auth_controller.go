package controllers

import (
	"encoding/json"
	"github.com/genshen/blog/components/auth"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) Callback() {
	user := &auth.User{}
	if this.HasAuth() {
		u := this.GetUserData()
		user = &u
	} else {
		code := this.GetString("code")
		github := auth.GithubAuthUser{}
		if len(code) > 0 {
			u, err := auth.StartAuth(&github, code)
			if err == nil {
				user = u
				user.Status = auth.UserStatusHasAuth
				this.LoginUser(user)
			}
		}
	}
	b, _ := json.Marshal(user)
	this.Data["json"] = string(b)
	this.TplName = "home/auth_callback.html"
}

func (this *AuthController)LoginUser(u *auth.User) {
	this.SetSession(UserData, *u)
	this.SetSession(IsAuth, true)
}