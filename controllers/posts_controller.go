package controllers

import (
	"github.com/genshen/blog/components/context/posts"
)

const DefPostListCount = 20

type PostsController struct {
	BaseController
}

// load all posts
func (post *PostsController) List() {
	showSummary, _ := post.GetBool("show_summary", false)
	skip, _ := post.GetInt("skip", 0)
	count, _ := post.GetInt("count", DefPostListCount)
	if count > DefPostListCount {
		count = DefPostListCount
	}
	if showSummary {
		post.Data["json"] = posts.LoadPostSummaryLists(skip, count)
	} else {
		post.Data["json"] = posts.LoadPostLists(skip, count)
	}
	post.ServeJSON()
}

func (post *PostsController) Detail() {
	id := post.Ctx.Input.Param(":id")
	post.Data["json"] = posts.LoadPostDetail(id)
	post.ServeJSON()
}
