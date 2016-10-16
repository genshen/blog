package keys

import (
	"qiniupkg.com/api.v7/kodo"
	"github.com/astaxie/beego"
)

var QiniuConfig  struct {
	BucketName string
	Expires    uint32
	Domain     string
	UploadPath string
}

func loadQiniuKeys() {
	kodo.SetMac(beego.AppConfig.String("qiniu_access_key"), beego.AppConfig.String("qiniu_secret_key"))
	QiniuConfig.BucketName = beego.AppConfig.String("qiniu_bucket_name")
	QiniuConfig.Expires = uint32(beego.AppConfig.DefaultInt("qiniu_token_expires", 60))
	QiniuConfig.Domain = beego.AppConfig.String("qiniu_config_domain")
	QiniuConfig.UploadPath = beego.AppConfig.String("qiniu_config_upload_path")
}
