package posts

import (
	"time"
	"gensh.me/blog/models"
	"gopkg.in/mgo.v2/bson"
	"gensh.me/blog/models/database"
)

type PostLists struct {
	Id           string    `json:"id"`
	Title        *string    `json:"title"`
	Summary      *string    `json:"summary"`
	Cover        *string    `json:"cover"`
	ViewCount    int        `json:"view_count"`
	CommentCount int        `json:"comment_count"`
	CreatedAt    *time.Time `json:"create_at"`
}

func LoadPostLists() *[]PostLists {
	//todo cache
	posts := []models.Posts{}
	database.DB.C(models.CollectionName_Posts).Find(bson.M{}).All(&posts)
	list := make([]PostLists, 0, len(posts))
	for _, post := range posts {
		list = append(list, PostLists{Id:post.Id.Hex(), Title:&post.Title, Summary:&post.Summary,
			Cover:&post.Cover, ViewCount:post.ViewCount, CommentCount:post.CommentCount,
			CreatedAt:&post.CreatedAt});
	}
	return &list
}