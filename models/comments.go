package models

import (
	"gopkg.in/mgo.v2/bson"
	"gensh.me/blog/components/auth"
	"time"
)

const (
	CommentStatusDeleted = iota
	CommentStatusActive
)

const (
	CollectionName_Comments = "comments"
)

type Comment struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	PostId   bson.ObjectId  //postid
	User     auth.User
	Content  string
	Replies  []Reply
	Status   int
	CreatedAt time.Time
}

type Reply struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	User     auth.User
	Content  string
	Status   int
	CreateAt time.Time
}