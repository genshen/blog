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

	if db_auth := beego.AppConfig.DefaultBool("db_auth", false); db_auth {
		if err := DB.Login(beego.AppConfig.String("db_auth_user"),beego.AppConfig.String("db_auth_pwd")); err != nil { //if auth failed
			log.Fatalln(err)
		}
	}

}