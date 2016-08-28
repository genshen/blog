package controllers

import (
	"gensh.me/blog/components/utils"
	"gensh.me/blog/components/context/comments"
	"strconv"
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
		if errors, comment_id := form.ValidAndSave(&user); errors == nil {
			result = &utils.SimpleJsonResponse{Status:1, Addition:comment_id}
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

func (this *CommentController)ReplyAdd() {
	var result *utils.SimpleJsonResponse
	if this.HasAuth() {
		user := this.GetUserData()
		comment_id := this.GetString("comment_id")
		content := this.GetString("content")
		form := comments.ReplyAddForm{CommentId:comment_id, Content:content}
		if errors := form.ValidAndSave(&user); errors == nil {
			result = &utils.SimpleJsonResponse{Status:1, Addition:"success"}
		} else {
			result = &utils.SimpleJsonResponse{Status:0,
				Error:utils.NewInstant(errors, map[string]string{"comment_id":"", "content":""})}
		}
	} else {
		result = &utils.SimpleJsonResponse{Status:2, Error:"需要登录后才能操作"}
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *CommentController)Load() {
	post_id := this.Ctx.Input.Param(":post_id")
	start, _ := strconv.ParseInt(this.Ctx.Input.Param(":start"), 10, 8)
	this.Data["json"] = comments.FindCommentsById(post_id, int(start))
	this.ServeJSON()
}