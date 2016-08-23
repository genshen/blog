package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	PostStatusDraft = iota
	PostStatusActive
	PostStatusDeleted
)

const (
	CollectionName_Posts = "posts"
)

type Posts struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Title        string
	Content      string
	Summary      string
	ViewCount    int `bson:"view_count"`
	CommentCount int `bson:"comment_count"`
	ReplyCount   int `bson:"reply_count"`
	//Classify         string
	//Tags         string
	Status       int
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}