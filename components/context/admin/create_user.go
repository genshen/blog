package admin

import (
	"github.com/astaxie/beego/logs"
	"gensh.me/blog/models"
	"gensh.me/blog/models/database"
	"gensh.me/blog/components/auth"
)

func CreateUser(username, email, password string)bool {
	admin := models.Admin{Username: username, Email: email, PasswordToken: auth.SHA1(password), Status: models.AdminStatusActive}
	if err := database.DB.C(models.CollectionAdmin).Insert(&admin); err != nil {
		logs.Error(err)
		return false;
	}
	return true;
}
