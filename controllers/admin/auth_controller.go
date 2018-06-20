package admin

import (
	"github.com/genshen/blog/components/context/admin"
	"github.com/genshen/blog/components/utils"
	"html/template"
	"github.com/astaxie/beego"
)

const (
	UserId   = "AdminUserID"
	Username = "AdminUsername"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) SignIn() {
	if this.Ctx.Request.Method == "POST" {
		signInForm := admin.SignInForm{Email: this.GetString("email"), Password: this.GetString("password")}
		if errs := signInForm.Valid(); errs != nil {
			s := utils.NewSingleErrorInstant(errs)
			this.Data["json"] = &utils.SimpleJsonResponse{Status: 0, Error: &s}
		} else {
			next := this.GetString("next")
			if len(next) > 0 && next[0] != '/' {
				next = "/" + next
			} else if next == "" {
				next = AdminHomePage
			}

			var con utils.UserInfo // todo get user information
			if token, _, err := utils.JwtNewToken(con, ""); err != nil {
				this.Data["json"] = &utils.SimpleJsonResponse{Status: 0,
					Error: map[string]string{"jwt_error": "generating jwt error."}}
			} else {
				a := struct {
					Next     string `json:"next"`
					JwtToken string `json:"jwt_token"`
				}{next, token}
				this.Data["json"] = &utils.SimpleJsonResponse{Status: 1, Addition: a}
			}
		}
		this.ServeJSON()
	} else {
		this.Data["dev"] = false;
		this.TplName = "admin/auth/sign_in.html"
	}
}

/*only in Dev Mode*/
type SignUpForm struct {
	Email    string
	Username string
	Password string
}

func (this *AuthController) SignUp() {
	if (this.Ctx.Request.Method == "POST") { //todo
		email := this.GetString("Email")
		username := this.GetString("Username")
		password := this.GetString("Password")
		admin.CreateUser(username, email, password)
		this.Data["json"] = &utils.SimpleJsonResponse{Status: 1, Addition: ""};
		this.ServeJSON()
	} else {
		this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
		this.Data["Form"] = &SignUpForm{}
		this.TplName = "admin/auth/sign_up.html"
	}
}
