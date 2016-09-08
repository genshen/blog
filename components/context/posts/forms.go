package posts

import (
	"github.com/astaxie/beego/validation"
	"gensh.me/blog/models/database"
	"gensh.me/blog/models"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type PostAddForm struct {
	CategoryId string
	Title      string
	Content    string
	Summary    string
}

func (this *PostAddForm)ValidAndSave() []*validation.Error {
	valid := validation.Validation{}
	valid.Required(this.Title, "title").Message("标题不能为空")
	valid.MaxSize(this.Title, 64, "title").Message("标题长度不能超过64个字符")
	valid.Required(this.Content, "content").Message("内容不能为空")
	valid.Required(this.Summary, "summary").Message("内容摘要不能为空")
	if !bson.IsObjectIdHex(this.CategoryId) {
		valid.SetError("category_id", "未找到对应分类")
	}
	if valid.HasErrors() {
		return valid.Errors
	}
	return this.save(&valid)
}

func (this *PostAddForm) save(v *validation.Validation) []*validation.Error {
	now := time.Now()
	//todo return _id
	m := models.Posts{Title:this.Title, CategoryId:bson.ObjectIdHex(this.CategoryId), Content:this.Content, Summary:this.Summary,
		Status:models.PostStatusActive, CreatedAt:now, UpdatedAt:now}
	err := database.DB.C(models.CollectionName_Posts).Insert(&m)
	if err != nil {
		v.SetError("title", "保存失败,请重试")
		v.SetError("content", "保存失败,请重试")
		return v.Errors
	}
	return nil
}