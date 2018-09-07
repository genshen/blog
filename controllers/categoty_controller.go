package controllers

import (
	"github.com/astaxie/beego"
	"github.com/genshen/blog/components/context/category"
)

type CategoryController struct {
	beego.Controller
}

//get all categories
func (c *CategoryController)Get() {
	categories := category.GetCategories()
	c.Data["json"] = &categories
	c.ServeJSON()
}
