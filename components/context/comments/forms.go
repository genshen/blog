package comments

import (
	"time"
	"github.com/astaxie/beego/validation"
	"gopkg.in/mgo.v2/bson"
	"gensh.me/blog/models"
	"gensh.me/blog/models/database"
	"gensh.me/blog/components/auth"
)

type CommentAddForm struct {
	PostId  string
	Content string
}
type ReplyAddForm struct {
	CommentId string
	Content   string
}

func (this *CommentAddForm)ValidAndSave(user *auth.User) ([]*validation.Error, string) {
	valid := validation.Validation{}
	valid.Required(this.Content, "content").Message("评论内容不能为空")
	if !bson.IsObjectIdHex(this.PostId) {
		//valid required,length... here
		valid.SetError("post_id", "未找到对应的文章")
	}
	if valid.HasErrors() {
		return valid.Errors, ""
	}
	return this.save(user, &valid)
}

func (this *CommentAddForm) save(user *auth.User, v *validation.Validation) ([]*validation.Error, string) {
	now := time.Now()
	_id := bson.NewObjectId()
	m := models.Comment{Id:_id, PostId:bson.ObjectIdHex(this.PostId), User:*user, Content:this.Content,
		Status:models.CommentStatusActive, CreatedAt:now}
	err := database.DB.C(models.CollectionName_Comments).Insert(&m)
	if err != nil {
		v.SetError("post_id", "评论失败,请重试")
		v.SetError("content", "评论失败,请重试")
		return v.Errors, ""
	}
	return nil, _id.Hex()
}

func (this *ReplyAddForm)ValidAndSave(user *auth.User) []*validation.Error {
	valid := validation.Validation{}
	valid.Required(this.Content, "content").Message("回复内容不能为空")
	if !bson.IsObjectIdHex(this.CommentId) {
		//valid required,length... here
		valid.SetError("comment_id", "对应评论不存在")
	}
	if valid.HasErrors() {
		return valid.Errors
	}
	return this.save(user, &valid)
}

func (this *ReplyAddForm)save(user *auth.User, v *validation.Validation) []*validation.Error {
	now := time.Now()
	m := models.Reply{Id:bson.NewObjectId(), User:*user, Content:this.Content,
		Status:models.ReplyStatusActive, CreateAt:now}

	err := database.DB.C(models.CollectionName_Comments).Update(bson.M{"_id":bson.ObjectIdHex(this.CommentId)},
		bson.M{"$push":bson.M{"replies":&m}})
	if err != nil {
		v.SetError("post_id", "评论失败,请重试")
		v.SetError("content", "评论失败,请重试")
		return v.Errors
	}
	return nil
}