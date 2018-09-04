package keys

import (
	"github.com/genshen/blog/components/utils"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
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
	qiniuConfig := utils.CustomConfig.Storage.QiNiuConfig
	QiniuConfig.accessKey = qiniuConfig.AccessKey
	QiniuConfig.secretKey = qiniuConfig.SecretKey
	QiniuConfig.bucketName = qiniuConfig.BucketName
	QiniuConfig.expires = qiniuConfig.TokenExpires
	QiniuConfig.Domain = qiniuConfig.ConfigDomain
	QiniuConfig.UploadPath = qiniuConfig.UploadPath
}

func NewUploadToken() string {
	putPolicy := storage.PutPolicy{
		Scope: QiniuConfig.bucketName,
	}
	putPolicy.Expires = QiniuConfig.expires
	mac := qbox.NewMac(QiniuConfig.accessKey, QiniuConfig.secretKey)
	return putPolicy.UploadToken(mac)
}
