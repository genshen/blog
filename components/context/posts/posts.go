package posts

import (
	"github.com/genshen/blog/models"
	"github.com/genshen/blog/models/database"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type PostList struct {
	Id            string    `json:"id"`
	Title         string    `json:"title"`
	CategoryId    string    `json:"category_id"`
	SubCategoryId string    `json:"sub_category_id"`
	ViewCount     int       `json:"view_count"`
	CommentCount  int       `json:"comment_count"`
	ReplyCount    int       `json:"reply_count"`
	CreatedAt     time.Time `json:"create_at"`
}

type PostSummaryList struct {
	PostList
	Summary string `json:"summary"`
	Cover   string `json:"cover"`
}

// get posts list without summary information.
func LoadPostLists(skip, count int) *[]PostList {
	//todo memory copy
	var post models.Posts
	// do query
	iter := database.DB.C(models.CollectionName_Posts).Find(bson.M{}).
		Select(bson.M{"_id": 1, "title": 1, "category_id": 1, "sub_category_id": 1,
			"view_count": 1, "comment_count": 1, "reply_count": 1, "created_at": 1}).
		Limit(count).
		Skip(skip).
		Iter()

	list := make([]PostList, 0, count)
	for iter.Next(&post) {
		list = append(list,
			PostList{
				Id:            post.Id.Hex(),
				CategoryId:    post.CategoryId.Hex(),
				SubCategoryId: post.SubCategoryId.Hex(),
				Title:         post.Title,
				ViewCount:     post.ViewCount,
				CommentCount:  post.CommentCount,
				ReplyCount:    post.ReplyCount,
				CreatedAt:     post.CreatedAt,
			});
	}
	return &list
}

// get post list with summary information
func LoadPostSummaryLists(skip, count int) *[]PostSummaryList {
	var post models.Posts
	// do query
	iter := database.DB.C(models.CollectionName_Posts).Find(bson.M{}).
		Select(bson.M{"_id": 1, "title": 1, "category_id": 1, "sub_category_id": 1,
			"view_count": 1, "comment_count": 1, "reply_count": 1, "created_at": 1,
			"summary": 1, "cover": 1}).
		Limit(count).
		Skip(skip).
		Iter()

	list := make([]PostSummaryList, 0, count)
	for iter.Next(&post) {
		list = append(list,
			PostSummaryList{
				PostList: PostList{
					Id:            post.Id.Hex(),
					CategoryId:    post.CategoryId.Hex(),
					SubCategoryId: post.SubCategoryId.Hex(),
					Title:         post.Title,
					ViewCount:     post.ViewCount,
					CommentCount:  post.CommentCount,
					ReplyCount:    post.ReplyCount,
					CreatedAt:     post.CreatedAt,
				},
				Summary: post.Summary,
				Cover:   post.Cover,
			});
	}
	return &list
}

func LoadPostDetail(id string) *models.Posts {
	post := models.Posts{}
	objId := bson.ObjectIdHex(id)
	database.DB.C(models.CollectionName_Posts).FindId(objId).One(&post)
	return &post
}
