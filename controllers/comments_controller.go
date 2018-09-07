package controllers

import (
	"github.com/genshen/blog/components/auth"
	"github.com/genshen/blog/components/context/admin"
	"github.com/genshen/blog/components/context/comments"
	"github.com/genshen/blog/components/utils"
	"github.com/genshen/blog/controllers/base"
)

type CommentController struct {
	user auth.User // oauth2 user data
	base.JwtAuthController
}

func (com *CommentController) JwtQueryTokenKey() string {
	return admin.JwtAdminConfigQueryTokenKey
}

func (com *CommentController) TokenVerify(token string) {
	if claims, err := utils.JwtVerify(&auth.User{}, token); err != nil { // todo set claims
		com.SetUnAuth() // or call UnAuth interface directly
	} else {
		// verify passed.
		com.user = *(claims.(*auth.User)) // do interface conversion and pointer conversion.
	}
}

func (com *CommentController) Add() {
	var result *utils.SimpleJsonResponse
	postId := com.GetString("post_id")
	content := com.GetString("content")
	form := comments.CommentAddForm{PostId: postId, Content: content}
	if errors, comment_id := form.ValidAndSave(&com.user); errors == nil {
		result = &utils.SimpleJsonResponse{Status: 1, Addition: comment_id}
	} else {
		result = &utils.SimpleJsonResponse{Status: 0,
			Error: utils.NewInstant(errors, map[string]string{"post_id": "", "content": ""})}
	}

	com.Data["json"] = result
	com.ServeJSON()
}

func (com *CommentController) ReplyAdd() {
	var result *utils.SimpleJsonResponse

	comment_id := com.GetString("comment_id")
	content := com.GetString("content")
	form := comments.ReplyAddForm{CommentId: comment_id, Content: content}
	if errors := form.ValidAndSave(&com.user); errors == nil {
		result = &utils.SimpleJsonResponse{Status: 1, Addition: "success"}
	} else {
		result = &utils.SimpleJsonResponse{Status: 0,
			Error: utils.NewInstant(errors, map[string]string{"comment_id": "", "content": ""})}
	}

	com.Data["json"] = result
	com.ServeJSON()
}
