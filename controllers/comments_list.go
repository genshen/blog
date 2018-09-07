package controllers

import (
	"github.com/astaxie/beego"
	"github.com/genshen/blog/components/context/comments"
	"strconv"
)

type CommentsListController struct {
	beego.Controller
}

func (com *CommentsListController) CommentsList() {
	postId := com.Ctx.Input.Param(":post_id")
	start, _ := strconv.ParseInt(com.Ctx.Input.Param(":start"), 10, 8)
	com.Data["json"] = comments.FindCommentsById(postId, int(start))
	com.ServeJSON()
}
