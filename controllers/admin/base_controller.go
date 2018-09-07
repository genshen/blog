package admin

import (
	"github.com/genshen/blog/components/context/admin"
	"github.com/genshen/blog/components/utils"
	"github.com/genshen/blog/controllers/base"
)

// app controller needing auth will inherit this struct, and implement OnUnAuth interface.
type BaseAuthController struct {
	user admin.UserInfo
	base.JwtAuthController
}

func (b *BaseAuthController) JwtQueryTokenKey() string {
	return admin.JwtAdminConfigQueryTokenKey
}

func (b *BaseAuthController) TokenVerify(token string) {
	if claims, err := utils.JwtVerify(&admin.UserInfo{}, token); err != nil { // todo set claims
		b.SetUnAuth() // or call UnAuth interface directly
	} else {
		// verify passed.
		b.user = *(claims.(*admin.UserInfo)) // do interface conversion and pointer conversion.
	}
}
