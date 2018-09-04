package routers

import (
	"github.com/astaxie/beego"
	"github.com/genshen/blog/components/utils"
	"github.com/genshen/blog/controllers/admin"
)

func initAdminRoute() {
	api := utils.CustomConfig.Api
	//admin router
	//beego.Router(admin.AdminSignOutUri, &admin.AuthController{}, "get:SignOut") // todo remove signout variable
	beego.Router(api.AdminSignInPath, &admin.AuthController{}, "get,post:SignIn") // todo admin/sign_in POST can be in api.
	if beego.BConfig.RunMode == beego.DEV {
		//can register a admin user in Dev Mode
		beego.Router(api.AdminSignUpPath, &admin.AuthController{}, "get,post:SignUp")
	}

	beego.Router(api.AdminPagesPrefix+"/", &admin.PanelController{}, "get:Get")

	if utils.CustomConfig.Storage.EnableQiNiuCloud {
		beego.Router(api.AdminApiPrefix+"/upload_token", &admin.StorageController{}, "get:QiNiuCloudStorageUploadToken")
	} else {
		beego.Router(api.AdminApiPrefix+"/upload_token", &admin.StorageController{}, "get:LocalStorageUploadToken")
		// Note: no adminApi prefix !!
		beego.Router(utils.CustomConfig.Storage.LocalStorageDomain+":hash",
			&admin.LocalStorageHashController{}, "get:LocalStorageResource")
		// Note:urlFor is ues in function storage_controller.go#initStorage
		beego.Router(utils.CustomConfig.Storage.LocalStorageUploadUrl, &admin.StorageController{}, "post:LocalUpload")
		admin.InitStorage()
	}
	beego.Router(api.AdminApiPrefix+"/article", &admin.PostsController{}, "get:List")
	beego.Router(api.AdminApiPrefix+"/article/publish", &admin.PostsController{}, "post:Add")
	//beego.Router(admin.AdminApiPrefix+"/post/delete", &admin.PostsController{}, "post:Del")
	beego.Router(api.AdminApiPrefix+"/category/add", &admin.CategoryController{}, "post:CategoryAdd")
	beego.Router(api.AdminApiPrefix+"/sub_category/add", &admin.CategoryController{}, "post:SubCategoryAdd")
}
