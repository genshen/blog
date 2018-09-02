package routers

import (
	"github.com/astaxie/beego"
	"github.com/genshen/blog/controllers/admin"
)

func initAdminRoute() {
	//admin router
	//beego.Router(admin.AdminSignOutUri, &admin.AuthController{}, "get:SignOut") // todo remove signout variable
	beego.Router(admin.AdminSignInUri, &admin.AuthController{}, "get,post:SignIn")
	if beego.BConfig.RunMode == beego.DEV {
		//can register a admin user in Dev Mode
		beego.Router(admin.AdminPagesPrefix+"/dev/sign_up", &admin.AuthController{}, "get,post:SignUp")
	}

	beego.Router(admin.AdminPagesPrefix+"/", &admin.PanelController{}, "get:Get")

	if beego.AppConfig.DefaultBool("storage::EnableQiNiuCloud", false) {
		beego.Router(admin.AdminApiPrefix+"/upload_token", &admin.StorageController{}, "get:QiNiuCloudStorageUploadToken")
	} else {
		beego.Router(admin.AdminApiPrefix+"/upload_token", &admin.StorageController{}, "get:LocalStorageUploadToken")
		// Note: no adminApi prefix !!
		beego.Router(beego.AppConfig.DefaultString("storage::LocalStorageDomain", "/images/:hash"),
			&admin.LocalStorageHashController{}, "get:LocalStorageResource")
		// Note:urlFor is ues in function storage_controller.go#initStorage
		beego.Router(admin.AdminApiPrefix+"/upload", &admin.StorageController{}, "post:LocalUpload")
		admin.InitStorage()
	}
	beego.Router(admin.AdminApiPrefix+"/article", &admin.PostsController{}, "get:List")
	beego.Router(admin.AdminApiPrefix+"/article/publish", &admin.PostsController{}, "post:Add")
	//beego.Router(admin.AdminApiPrefix+"/post/delete", &admin.PostsController{}, "post:Del")
	beego.Router(admin.AdminApiPrefix+"/category/add", &admin.CategoryController{}, "post:CategoryAdd")
	beego.Router(admin.AdminApiPrefix+"/sub_category/add", &admin.CategoryController{}, "post:SubCategoryAdd")
}
