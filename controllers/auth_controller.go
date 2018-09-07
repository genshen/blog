package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/genshen/blog/components/auth"
	"github.com/genshen/blog/components/utils"
	"log"
	"os/exec"
	"runtime"
)

type AuthController struct {
	beego.Controller
}

func (a *AuthController) Callback() {
	//user := &auth.User{}
	data := utils.SimpleJsonResponse{Status: 0, Error: "bad request."}
	if code := a.GetString("code"); len(code) > 0 {
		runCmdBeforeAuth()
		github := auth.GithubAuthUser{}

		if err := auth.StartAuth(&github, code); err == nil {
			// generate token
			if token, _, err := utils.NewJwtToken(&(github.User), auth.OAuth2JwtIssuer, auth.OAuth2JwtTokenLifetime); err == nil {
				data = utils.SimpleJsonResponse{Status: 1, Addition: token}
			}
		} else {
			data = utils.SimpleJsonResponse{Status: 0, Error: "generating jwt error."}
			logs.Warning(err)
		}
	}
	b, _ := json.Marshal(data)
	a.Data["json"] = string(b)
	a.TplName = "home/auth_callback.html"
}

//
//func (a *AuthController) LoginUser(u *auth.User) {
//	a.SetSession(UserData, *u)
//	a.SetSession(IsAuth, true)
//}

func runCmdBeforeAuth() {
	// for now, cmd call is only used on linux.
	if cmdstr := utils.CustomConfig.Auth.BeforeAuth; cmdstr != "" && runtime.GOOS == "linux" {
		// run this command
		cmd := exec.Command("sh", "-c", cmdstr) // todo only for linux OS.
		// cmd.Dir = cacheDir
		cmd.Stdout = nil
		cmd.Stderr = nil
		if err := cmd.Run(); err != nil {
			log.Println("run command error", err)
		}
	}
}
