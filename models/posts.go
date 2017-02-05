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
	Id            bson.ObjectId  `bson:"_id,omitempty" json:"id"`
	Title         string         `json:"title"`
	Content       string         `json:"content"`
	Summary       string         `json:"summary"`
	Cover         string         `json:"cover"`
	ViewCount     int            `bson:"view_count" json:"view_count"`
	CommentCount  int            `bson:"comment_count" json:"comment_count"`
	ReplyCount    int            `bson:"reply_count" json:"reply_count"`
	CategoryId    bson.ObjectId  `bson:"category_id" json:"category_id"`
	SubCategoryId bson.ObjectId  `bson:"sub_category_id" json:"sub_category_id"`
	//Tags
	Status        int
	CreatedAt     time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time   `bson:"updated_at" json:"updated_at"`
}