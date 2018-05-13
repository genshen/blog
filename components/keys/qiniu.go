package keys

import (
	"github.com/astaxie/beego"
	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/api.v7/auth/qbox"
)

var QiniuConfig struct {
	bucketName string
	expires    uint32
	Domain     string
	UploadPath string
	accessKey  string
	secretKey  string
}

func loadQiniuKeys() {
	QiniuConfig.accessKey = beego.AppConfig.String("qiniu_access_key")
	QiniuConfig.secretKey = beego.AppConfig.String("qiniu_secret_key")
	QiniuConfig.bucketName = beego.AppConfig.String("qiniu_bucket_name")
	QiniuConfig.expires = uint32(beego.AppConfig.DefaultInt("qiniu_token_expires", 60))
	QiniuConfig.Domain = beego.AppConfig.String("qiniu_config_domain")
	QiniuConfig.UploadPath = beego.AppConfig.String("qiniu_config_upload_path")
}

func NewUploadToken() string {
	putPolicy := storage.PutPolicy{
		Scope: QiniuConfig.bucketName,
	}
	putPolicy.Expires = QiniuConfig.expires
	mac := qbox.NewMac(QiniuConfig.accessKey, QiniuConfig.secretKey)
	return putPolicy.UploadToken(mac)
}
