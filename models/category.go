package models

import "gopkg.in/mgo.v2/bson"

const (
	CategoryStatusDeleted = iota
	CategoryStatusActive
)

const (
	CollectionName_Category = "category"
)

type Category struct {
	Id         bson.ObjectId  `bson:"_id,omitempty" json:"id"`
	FatherName string   `json:"f_name"`
	FatherSlug string   `json:"f_slug"`
	Name       string   `json:"name"`
	Slug       string   `json:"slug"`
	PostsCount int      `json:"posts_count"`
}