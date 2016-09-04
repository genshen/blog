package settings

import (
	"log"
	"gopkg.in/mgo.v2/bson"
	"gensh.me/blog/models"
	"gensh.me/blog/models/database"
)

func LoadCategories() {
	mu.Lock()
	defer mu.Unlock()
	err := database.DB.C(models.CollectionName_Category).Find(bson.M{}).All(&s.Categories)
	if err != nil {
		log.Println("error loading categories")
	}
}