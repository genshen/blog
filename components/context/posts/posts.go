package posts

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"github.com/genshen/blog/models"
	"github.com/genshen/blog/models/database"
)

type PostList struct {
	Id           string    `json:"id"`
	CategoryId   string    `json:"category_id"`
	Title        string    `json:"title"`
	Summary      string    `json:"summary"`
	Cover        string    `json:"cover"`
	ViewCount    int       `json:"view_count"`
	CommentCount int       `json:"comment_count"`
	CreatedAt    time.Time `json:"create_at"`
}

func LoadPostLists() *[]PostList {
	//todo cache and memory copy
	posts := []models.Posts{}
	database.DB.C(models.CollectionName_Posts).Find(bson.M{}).All(&posts)
	list := make([]PostList, 0, len(posts))
	for _, post := range posts {
		list = append(list, PostList{Id: post.Id.Hex(), CategoryId: post.CategoryId.Hex(),
			Title: post.Title, Summary: post.Summary, Cover: post.Cover,
			ViewCount: post.ViewCount, CommentCount: post.CommentCount, CreatedAt: post.CreatedAt});
	}
	return &list
}

func LoadPostDetail(id string) *models.Posts {
	post := models.Posts{}
	objId := bson.ObjectIdHex(id)
	database.DB.C(models.CollectionName_Posts).FindId(objId).One(&post)
	return &post
}
