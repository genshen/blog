package admin

import (
	"gensh.me/blog/models/database"
	"gensh.me/blog/models"
	security "gensh.me/blog/components/auth"
	"gopkg.in/mgo.v2/bson"
)

func AddAdmin(username string, email string, password string) {
	database.DB.C(models.CollectionAdmin).Insert(&models.Admin{Username:username,
		Email:email, PasswordToken:security.SHA1(password), Status:models.AdminStatusActive})
}

func ValidPassword(email string, password string) (string, *models.Admin) {
	admin := models.Admin{}
	database.DB.C(models.CollectionAdmin).Find(
		bson.M{"email": email, "password_token":security.SHA1(password)}).One(&admin)
	if id := admin.Id.Hex(); id != "" {
		return id, &admin
	}
	return "", nil
}
