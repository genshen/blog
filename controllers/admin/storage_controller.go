package admin

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/genshen/blog/components/keys"
	"github.com/genshen/blog/components/utils"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

type StorageController struct {
	BaseAuthController
}

type StorageToken struct {
	Token      string `json:"token"`
	Domain     string `json:"domain"`
	UploadPath string `json:"upload_path"`
}

const (
	LocalUploadURLFor          = "StorageController.LocalUpload"
	LocalStorageResourceURLFor = "StorageController.LocalStorageResource"
)

var localStorageConfig struct {
	UploadUrl  string
	StorageDir string
	Domain     string
}

// init storage configure in router/router.go
func InitStorage() {
	if !utils.CustomConfig.Storage.EnableQiNiuCloud {
		localStorageConfig.UploadUrl = beego.URLFor(LocalUploadURLFor)
		localStorageConfig.StorageDir = utils.CustomConfig.Storage.LocalStorageDir
		localStorageConfig.Domain = beego.URLFor(LocalStorageResourceURLFor, ":hash", "")
		for pre := 0; pre <= 0xF; pre++ {
			for next := 0; next <= 0xF; next++ {
				suffix := fmt.Sprintf("/%x/%x", pre, next)
				newPath := filepath.Join(".", localStorageConfig.StorageDir, suffix)
				if _, err := os.Stat(newPath); os.IsNotExist(err) {
					if err = os.MkdirAll(newPath, 0744); err != nil {
						log.Fatalln(err)
					}
				}
			}
		}
	}
}

func (c *StorageController) QiNiuCloudStorageUploadToken() {
	up_token := keys.NewUploadToken()
	c.Data["json"] = &StorageToken{Token: up_token, Domain: keys.QiniuConfig.Domain, UploadPath: keys.QiniuConfig.UploadPath}
	c.ServeJSON()
}

func (c *StorageController) LocalStorageUploadToken() {
	c.Data["json"] = &StorageToken{Token: "token", Domain: localStorageConfig.Domain, UploadPath: localStorageConfig.UploadUrl}
	c.ServeJSON()
}

// Note:urlFor is ues in function storage_controller.go#initStorage
func (c *StorageController) LocalUpload() {
	// todo add lock
	f, h, err := c.GetFile("file")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	var path = h.Filename
	var body = "{\"error\":true}" //default: has error.
	if hash := hashFile(f); len(hash) >= 3 {
		path = filepath.Join(localStorageConfig.StorageDir, hash[0:1]+"/"+hash[1:2]+"/"+hash[2:])
		c.SaveToFile("file", path) //todo check exist.
		body = "{\"key\":\"" + hash + "\",\"hash\":\"" + hash + "\"}"
	}
	c.Ctx.Output.Header("Content-type", "application/json")
	c.Ctx.Output.Body([]byte(body))
}

// calculate a file hash,returns a Hexadecimal string of hash sum.
func hashFile(file multipart.File) string {
	//file, err := os.Open("log.log")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	// Create new hasher, which is a writer interface
	hasher := sha1.New()
	_, err := io.Copy(hasher, file)
	if err != nil {
		logs.Error("error Hash File:", err)
		return ""
	}

	// Hash and print. Pass nil since
	// the data is not coming in as a slice argument
	// but is coming through the writer interface
	sum := hasher.Sum(nil)
	return fmt.Sprintf("%x", sum)
}
