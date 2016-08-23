package database

import (
	"log"
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
)

var DB *mgo.Database

type DBLog struct{}

func (l DBLog)Output(calldepth int, s string) error {
	log.Println(s)
	return nil
}

func init() {
	if db_debug := beego.AppConfig.DefaultBool("db_debug", false); db_debug {
		mgo.SetDebug(true)
		mgo.SetLogger(DBLog{})
	}

	session, err := mgo.Dial(beego.AppConfig.String("db_config"))
	if err != nil {
		log.Fatalln(err)
	}
	DB = session.DB(beego.AppConfig.String("db_name"))
}