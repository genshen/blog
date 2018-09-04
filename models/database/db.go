package database

import (
	"github.com/genshen/blog/components/utils"
	"gopkg.in/mgo.v2"
	"log"
)

var DB *mgo.Database

type DBLog struct{}

func (l DBLog) Output(calldepth int, s string) error {
	log.Println(s)
	return nil
}

func init() {
	dbconfig := utils.CustomConfig.Database
	if dbconfig.DbDebug {
		mgo.SetDebug(true)
		mgo.SetLogger(DBLog{})
	}

	session, err := mgo.Dial(dbconfig.DbConfig)
	if err != nil {
		log.Fatalln(err)
	}
	DB = session.DB(dbconfig.DbName)

	if db_auth := dbconfig.DbAuth; db_auth {
		if err := DB.Login(dbconfig.DbAuthUser, dbconfig.DbAuthPwd); err != nil { //if auth failed
			log.Fatalln(err)
		}
	}
}
