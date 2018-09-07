package models

import (
	"github.com/genshen/blog/components/auth"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CommentStatusDeleted = iota
	CommentStatusActive
)

const (
	ReplyStatusDeleted = iota
	ReplyStatusActive
)

const (
	CollectionName_Comments = "comments"
)

type Comment struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	PostId    bson.ObjectId `json:"-"` //post id in db
	User      auth.User     `json:"user"`
	Content   string        `json:"content"`
	Replies   []Reply       `json:"replies"`
	Status    int           `json:"status"`
	CreatedAt time.Time     `json:"create_at"`
}

type Reply struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	User     auth.User     `json:"user"`
	Content  string        `json:"content"`
	Status   int           `json:"status"`
	CreateAt time.Time     `json:"create_at"`
}
