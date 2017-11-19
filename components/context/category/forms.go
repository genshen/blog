package category

import (
	"github.com/astaxie/beego/validation"
	"github.com/genshen/blog/models"
	"github.com/genshen/blog/models/database"
	"qiniupkg.com/x/log.v7"
	"gopkg.in/mgo.v2/bson"
)

type CategoryAddForm struct {
	Name string
	Slug string
}

type SubCategoryAddForm struct {
	CategoryId string
	Name       string
	Slug       string
	PostCount  int
}

func (this *CategoryAddForm)ValidAndSave() ([]*validation.Error, string) {
	valid := validation.Validation{}
	valid.Required(this.Name, "name").Message("名称不能为空")
	valid.Required(this.Slug, "slug").Message("内链不能为空")
	valid.MaxSize(this.Slug, 16, "slug").Message("内链长度不能超过16个字符")
	if valid.HasErrors() {
		return valid.Errors, ""
	}
	return this.save(&valid)
}

func (this *CategoryAddForm)save(v *validation.Validation) ([]*validation.Error, string) {
	_id := bson.NewObjectId()
	m := models.Category{Id:_id, Name:this.Name, Slug:this.Slug}
	err := database.DB.C(models.CollectionName_Category).Insert(&m)
	if err != nil {
		log.Println(err)
		v.SetError("name", "保存失败,请重试")
		v.SetError("slug", "保存失败,请重试")
		return v.Errors, ""
	}
	LoadCategories() //refresh
	return nil, _id.Hex()
}

func (this *SubCategoryAddForm)ValidAndSave() ([]*validation.Error, string) {
	valid := validation.Validation{}
	valid.Required(this.Name, "name").Message("名称不能为空")
	valid.Required(this.Slug, "slug").Message("内链不能为空")
	valid.MaxSize(this.Slug, 16, "slug").Message("内链长度不能超过16个字符")
	if !bson.IsObjectIdHex(this.CategoryId) {
		valid.SetError("id", "未找到对应的父分类")
	}
	if valid.HasErrors() {
		return valid.Errors, ""
	}
	return this.save(&valid)
}

func (this *SubCategoryAddForm)save(v *validation.Validation) ([]*validation.Error, string) {
	_id := bson.NewObjectId()
	m := models.SubCategory{Id:_id, Name:this.Name, Slug:this.Slug} //PostCount  = 0
	//err := database.DB.C(models.CollectionName_Category).Insert(&m)
	err := database.DB.C(models.CollectionName_Category).Update(bson.M{"_id":bson.ObjectIdHex(this.CategoryId)},
		bson.M{"$push":bson.M{"subcategory":&m}})
	if err != nil {
		log.Println(err)
		v.SetError("name", "保存失败,请重试")
		v.SetError("slug", "保存失败,请重试")
		return v.Errors, ""
	}
	LoadCategories()
	return nil, _id.Hex()
}