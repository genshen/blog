package comments

import (
	"gensh.me/blog/models"
	"gensh.me/blog/models/database"
	"gopkg.in/mgo.v2/bson"
)

func FindCommentsById(post_id string,start int) *[]models.Comment{
	comments := []models.Comment{}
	database.DB.C(models.CollectionName_Comments).Find(bson.M{"postid":bson.ObjectIdHex(post_id)}).
	Limit(20).Skip(start).All(&comments)
	return &comments
}