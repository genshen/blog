package category

import (
	"log"
	"sync"
	"gopkg.in/mgo.v2/bson"
	"github.com/genshen/blog/models"
	"github.com/genshen/blog/models/database"
)

var mu sync.RWMutex
var categories []models.Category

func init() {
	LoadCategories()
}

func LoadCategories() {
	mu.Lock()
	defer mu.Unlock()
	err := database.DB.C(models.CollectionName_Category).Find(bson.M{}).All(&categories)
	if err != nil {
		log.Println("error loading categories")
	}
}

func GetCategories() []models.Category {
	mu.RLock()
	defer mu.RUnlock()
	return categories
}