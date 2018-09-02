package models

import "gopkg.in/mgo.v2/bson"

type Tags struct {
	Id            bson.ObjectId  `bson:"_id,omitempty" json:"id"`
}