package main

import (
	_ "gensh.me/blog/routers"
	_"gensh.me/blog/models/database"
	"github.com/astaxie/beego"
)

func main() {
	//admin.CreateUser("genshen","genshenchu@gmail.com","genshen1234")
	beego.Run()
}