package controllers

import (
	"gensh.me/blog/components/utils"
	"gensh.me/blog/components/context/comments"
)

type CommentController struct {
	BaseController
}

func (this *CommentController) Add() {
	var result *utils.SimpleJsonResponse
	if this.HasAuth() {
		user := this.GetUserData()
		post_id := this.GetString("post_id")
		content := this.GetString("content")
		form := comments.CommentAddForm{PostId:post_id, Content:content}
		if errors := form.ValidAndSave(&user); errors == nil {
			result = &utils.SimpleJsonResponse{Status:1, Addition:"评论成功"}
		} else {
			result = &utils.SimpleJsonResponse{Status:0,
				Error:utils.NewInstant(errors, map[string]string{"post_id":"", "content":""})}
		}
	} else {
		result = &utils.SimpleJsonResponse{Status:2, Error:"需要登录后才能操作"}
	}
	this.Data["json"] = result
	this.ServeJSON()
}