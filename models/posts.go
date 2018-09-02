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

// post outline with post contents.
type Posts struct {
	Id            bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title         string        `bson:"title" json:"title"`
	Content       string        `bson:"content" json:"content"`
	Summary       string        `bson:"summary" json:"summary"`
	Cover         string        `bson:"cover" json:"cover"`
	CategoryId    bson.ObjectId `bson:"category_id" json:"category_id"`
	SubCategoryId bson.ObjectId `bson:"sub_category_id" json:"sub_category_id"`
	ViewCount     int           `bson:"view_count" json:"view_count"`
	CommentCount  int           `bson:"comment_count" json:"comment_count"`
	ReplyCount    int           `bson:"reply_count" json:"reply_count"`
	CreatedAt     time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time     `bson:"updated_at" json:"updated_at"`
	Status        int
	//Tags
}
