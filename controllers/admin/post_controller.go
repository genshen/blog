package admin

import (
	"github.com/genshen/blog/components/utils"
	"github.com/genshen/blog/components/context/posts"
)

type PostsController struct {
	BaseAuthController
}

func (this *PostsController) List() {
	this.ServeJSON() // todo
}

func (this *PostsController) Add() {
	form := posts.PostAddForm{CategoryId: this.GetString("category_id"), SubCategoryId: this.GetString("sub_category_id"),
		Title: this.GetString("title"), Content: this.GetString("content"), Summary: this.GetString("summary")}
	if errors := form.ValidAndSave(); errors == nil {
		this.Data["json"] = &utils.SimpleJsonResponse{Status: 1, Addition: ""}
	} else {
		this.Data["json"] = &utils.SimpleJsonResponse{Status: 0,
			Error: utils.NewInstant(errors, map[string]string{"category_id": "", "title": "", "content": "", "summary": ""})}
	}
	this.ServeJSON()
}
