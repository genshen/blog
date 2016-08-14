package admin

import "github.com/astaxie/beego/validation"

type SignInForm struct {
	ID       string
	Username string
	Email    string
	Password string
}

func (this *SignInForm)Valid() ([]*validation.Error) {
	valid := validation.Validation{}
	valid.Required(this.Email, "email").Message("邮箱不能为空")
	valid.Email(this.Email, "email").Message("邮箱格式不正确")
	valid.Required(this.Password, "password").Message("密码不能为空")
	valid.AlphaNumeric(this.Password, "password").Message("密码只能为字母和数字")
	if valid.HasErrors() {
		return valid.Errors
	}
	//验证密码
	if id, admin := ValidPassword(this.Email, this.Password); id == "" {
		valid.SetError("email", "用户名或密码错误")
		valid.SetError("password", "用户名或密码错误")
		return valid.Errors
	} else {
		this.ID = id
		this.Username = admin.Username
		return nil
	}
}