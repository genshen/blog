package admin

const (
	UserId = "AdminUserID"
	Username = "AdminUsername"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) SignIn() {
	this.TplName = "admin/auth/signin.html"
}

func (this *AuthController) SignIn_POST() {
	this.TplName = "admin/auth/signin.html"
}