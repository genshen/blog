package admin

import (
	"github.com/astaxie/beego"
	"github.com/genshen/blog/components/utils"
	"strings"
)

var (
	AdminSignUpUri  string
	AdminSignOutUri string
	AdminSignInUri  string
	AdminHomePage   string

	AdminPagesPrefix string
	AdminApiPrefix   string
)

type UnAuth interface {
	OnUnAuth()
}

// app controller needing auth will inherit this struct, and implement OnUnAuth interface.
type BaseAuthController struct {
	user utils.UserInfo
	beego.Controller
}

func init() {
	AdminPagesPrefix = utils.CustomConfig.Api.AdminPagesPrefix
	AdminApiPrefix = utils.CustomConfig.Api.AdminApiPrefix

	AdminSignUpUri = AdminPagesPrefix + utils.CustomConfig.Api.AdminSignUpPath
	AdminSignInUri = AdminPagesPrefix + utils.CustomConfig.Api.AdminSignInPath
	AdminSignOutUri = AdminPagesPrefix + utils.CustomConfig.Api.AdminSignOutPath

	AdminHomePage = AdminPagesPrefix + utils.CustomConfig.Api.AdminHomePath
}

func (b *BaseAuthController) Prepare() {
	var authHead = b.Ctx.Input.Header("Authorization")
	var token string
	if authHead != "" {
		// Authorization: Bearer <token>
		lIndex := strings.LastIndex(authHead, " ")
		if lIndex < 0 || lIndex+1 >= len(authHead) {
			b.SetUnAuth()
			return
		} else {
			token = authHead[lIndex+1:]
		}
	} else {
		if token = b.GetString(utils.JwtAdminConfigQueryTokenKey); token == "" {
			b.SetUnAuth()
			return
		} // else token != "", then passed and go on running
	}

	if claims, err := utils.JwtVerify(token); err != nil { // todo set claims
		b.SetUnAuth()
	} else {
		// check passed.
		b.user = claims.UserInfo
	}
}

func (b *BaseAuthController) SetUnAuth() {
	if app, ok := b.AppController.(UnAuth); ok {
		app.OnUnAuth()
	} else {
		//log.Panic("the UnAuth interface must be implemented.")
		b.Ctx.Output.Status = 401
		b.Ctx.Output.Body([]byte("UnAuthenticated"))
		b.StopRun()
	}
}
