package admin

import (
	"gensh.me/blog/components/context/posts"
	"gensh.me/blog/components/utils"
	"qiniupkg.com/api.v7/kodo"
	"gensh.me/blog/components/keys"
)

type PostsController struct {
	BaseController
}

func (this *PostsController) List() {
	this.ServeJSON()
}

func (this *PostsController) Add() {
	form := posts.PostAddForm{CategoryId:this.GetString("category_id"),SubCategoryId:this.GetString("sub_category_id"),
		Title:this.GetString("title"), Content:this.GetString("content"), Summary:this.GetString("summary")}
	if errors := form.ValidAndSave(); errors == nil {
		this.Data["json"] = &utils.SimpleJsonResponse{Status:1, Addition:""}
	} else {
		this.Data["json"] = &utils.SimpleJsonResponse{Status:0,
			Error:utils.NewInstant(errors,
				map[string]string{"category_id":"", "title":"", "content":"", "summary":""})}
	}
	this.ServeJSON()
}

type QiNiuToken struct {
	Token      string `json:"token"`
	Domain     string `json:"domain"`
	UploadPath string `json:"upload_path"`
}

func (this *PostsController)UploadToken() {
	zone := 0
	c := kodo.New(zone, nil) // 创建一个 Client 对象
	policy := &kodo.PutPolicy{
		Scope:   keys.QiniuConfig.BucketName,
		Expires: keys.QiniuConfig.Expires,
	}
	up_token := c.MakeUptoken(policy)
	this.Data["json"] = &QiNiuToken{Token:up_token, Domain:keys.QiniuConfig.Domain, UploadPath:keys.QiniuConfig.UploadPath}
	this.ServeJSON()
}