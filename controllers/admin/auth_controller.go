package admin

import (
	"github.com/astaxie/beego"
	"github.com/genshen/blog/components/context/admin"
	"github.com/genshen/blog/components/utils"
	"html/template"
)

const (
	UserId   = "AdminUserID"
	Username = "AdminUsername"
)

type AuthController struct {
	beego.Controller
}

func (auth *AuthController) SignIn() {
	if auth.Ctx.Request.Method == "POST" {
		signInForm := admin.SignInForm{Email: auth.GetString("email"), Password: auth.GetString("password")}
		if errs := signInForm.Valid(); errs != nil {
			s := utils.NewSingleErrorInstant(errs)
			auth.Data["json"] = &utils.SimpleJsonResponse{Status: 0, Error: &s}
		} else {
			next := auth.GetString("next")
			if len(next) > 0 && next[0] != '/' {
				next = "/" + next
			} else if next == "" {
				next = utils.CustomConfig.Api.AdminHomePath
			}

			var con admin.UserInfo // todo get user information
			if token, _, err := utils.NewJwtToken(&con, admin.JwtAdminIssuer, admin.AdminConfigJwtTokenLifetime); err != nil {
				auth.Data["json"] = &utils.SimpleJsonResponse{Status: 0,
					Error: map[string]string{"jwt_error": "generating jwt error."}}
			} else {
				a := struct {
					Next     string `json:"next"`
					JwtToken string `json:"jwt_token"`
				}{next, token}
				auth.Data["json"] = &utils.SimpleJsonResponse{Status: 1, Addition: a}
			}
		}
		auth.ServeJSON()
	} else {
		auth.Data["dev"] = false;
		auth.TplName = "admin/auth/sign_in.html"
	}
}

/*only in Dev Mode*/
type SignUpForm struct {
	Email    string
	Username string
	Password string
}

func (auth *AuthController) SignUp() {
	if (auth.Ctx.Request.Method == "POST") { //todo
		email := auth.GetString("Email")
		username := auth.GetString("Username")
		password := auth.GetString("Password")
		admin.CreateUser(username, email, password)
		auth.Data["json"] = &utils.SimpleJsonResponse{Status: 1, Addition: ""};
		auth.ServeJSON()
	} else {
		auth.Data["xsrfdata"] = template.HTML(auth.XSRFFormHTML())
		auth.Data["Form"] = &SignUpForm{}
		auth.TplName = "admin/auth/sign_up.html"
	}
}
