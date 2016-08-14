package models

import "gopkg.in/mgo.v2/bson"

const CollectionAdmin = "admin"
const (
	AdminStatusUnActive = iota
	AdminStatusActive
)

type Admin struct {
	Id            bson.ObjectId `bson:"_id,omitempty"`
	Username      string
	Email         string
	PasswordToken string        `bson:"password_token"`
	Status        int
}