package base

import (
	"github.com/astaxie/beego"
	"log"
	"strings"
)

// the derived class can make an implementation of this interface to
// send un-authenticated error back to client.
type UnAuth interface {
	OnUnAuth()
}

type JwtAuthInterface interface {
	// it returns the key of jwt token in http query.
	JwtQueryTokenKey() string
	// if the jwt token exists, this function will be called to verify and save jwt data.
	TokenVerify(token string)
}

type JwtAuthController struct {
	beego.Controller
}

func (jwt *JwtAuthController) Prepare() {
	// check jwt interface implementation
	jwtInterface, ok := jwt.AppController.(JwtAuthInterface)
	if !ok {
		ctrName, actionName := jwt.GetControllerAndAction()
		log.Fatal("the struct of dose not implement interface JwtAuthController in ", ctrName, actionName)
		return
	}

	var authHead = jwt.Ctx.Input.Header("Authorization")
	var token string
	// try to get jwt token in http request header.
	if authHead != "" {
		// Authorization: Bearer <token>
		lIndex := strings.LastIndex(authHead, " ")
		if lIndex < 0 || lIndex+1 >= len(authHead) {
			jwt.SetUnAuth()
			return
		} else {
			token = authHead[lIndex+1:]
		}
	} else {
		// try to get jwt token in http query data
		if token = jwt.GetString(jwtInterface.JwtQueryTokenKey()); token == "" {
			jwt.SetUnAuth()
			return
		} // else token != "", then passed and go on running
	}
	jwtInterface.TokenVerify(token)
}

func (jwt *JwtAuthController) SetUnAuth() {
	if app, ok := jwt.AppController.(UnAuth); ok {
		app.OnUnAuth()
	} else {
		//log.Panic("the UnAuth interface must be implemented.")
		jwt.Ctx.Output.Status = 401
		jwt.Ctx.Output.Body([]byte("UnAuthenticated"))
		jwt.StopRun()
	}
}
