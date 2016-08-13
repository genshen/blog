package database

import (
	"log"
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
)

var DB *mgo.Database
type Person struct {
	Name string
	Phone string
}

func init() {
	session, err := mgo.Dial(beego.AppConfig.String("db_config"))
	if err != nil {
		log.Fatalln(err)
	}
	DB = session.DB(beego.AppConfig.String("db_name"))
}