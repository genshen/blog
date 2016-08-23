package controllers

import (
	"gensh.me/blog/components/context/posts"
)

type PostsController struct {
	BaseController
}

func (this *PostsController) Category() {
	posts := posts.LoadPostLists()
	this.Data["json"] = posts
	this.ServeJSON()
}