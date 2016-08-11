package main

import (
	_ "gensh.me/blog/routers"
	"github.com/astaxie/beego"
	"gensh.me/blog/models/database"
)

func init() {
	database.InitDB()
}

func main() {
	beego.Run()
}