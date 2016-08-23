package controllers

import (
	"gensh.me/blog/components/context/posts"
)

type PostsController struct {
	BaseController
}

func (this *PostsController) Category() {
	this.Data["json"] = posts.LoadPostLists()
	this.ServeJSON()
}

func (this *PostsController) Detail() {
	id := this.Ctx.Input.Param(":id")
	this.Data["json"] =  posts.LoadPostDetail(id)
	this.ServeJSON()
}