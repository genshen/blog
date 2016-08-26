package admin

import (
	"gensh.me/blog/components/context/posts"
	"gensh.me/blog/components/utils"
)

type PostsController struct {
	BaseController
}

func (this *PostsController) List() {
	this.ServeJSON()
}

func (this *PostsController) Add() {
	form := posts.PostAddForm{Title:this.GetString("title"), Content:this.GetString("content"), Summary:this.GetString("summary")}
	if errors := form.ValidAndSave(); errors == nil {
		this.Data["json"] = &utils.SimpleJsonResponse{Status:1, Addition:""}
	}else{
		this.Data["json"] = &utils.SimpleJsonResponse{Status:0,
			Error:utils.NewInstant(errors,map[string]string{"title":"","content":"","summary":""})}
	}
	this.ServeJSON()
}