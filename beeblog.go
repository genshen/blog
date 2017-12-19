package main

import (
	_ "github.com/genshen/blog/routers"
	_"github.com/genshen/blog/models/database"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}