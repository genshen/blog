package admin

import (
	"github.com/genshen/blog/components/utils"
	"github.com/genshen/blog/components/context/category"
)

type CategoryController struct {
	BaseAuthController
}

//get all categories
func (this *CategoryController)Get() {
	c := category.GetCategories()
	this.Data["json"] = &c
	this.ServeJSON()
}

func (this *CategoryController)CategoryAdd() {
	form := category.CategoryAddForm{Name:this.GetString("name"), Slug:this.GetString("slug")}
	if errors, id := form.ValidAndSave(); errors == nil {
		this.Data["json"] = &utils.SimpleJsonResponse{Status:1, Addition:id}
	} else {
		this.Data["json"] = &utils.SimpleJsonResponse{Status:0,
			Error:utils.NewInstant(errors, map[string]string{"name":"", "slug":""})}
	}
	this.ServeJSON()
}

func (this *CategoryController)SubCategoryAdd() {
	form := category.SubCategoryAddForm{CategoryId:this.GetString("id"), Name:this.GetString("name"), Slug:this.GetString("slug")}
	if errors, id := form.ValidAndSave(); errors == nil {
		this.Data["json"] = &utils.SimpleJsonResponse{Status:1, Addition:id}
	} else {
		this.Data["json"] = &utils.SimpleJsonResponse{Status:0,
			Error:utils.NewInstant(errors, map[string]string{"id":"", "name":"", "slug":""})}
	}
	this.ServeJSON()
}