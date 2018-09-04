package admin

import (
	"github.com/genshen/blog/components/utils"
	"path/filepath"
	"os"
	"net/http"
	"time"
	"github.com/astaxie/beego"
)

// in this file, it serve local file uploaded to backend to frontend.

type LocalStorageHashController struct {
	beego.Controller
}

func (c *LocalStorageHashController) LocalStorageResource() {
	hash := c.Ctx.Input.Param(":hash")
	if len(hash) >= 3 {
		var path string = filepath.Join(utils.CustomConfig.Storage.LocalStorageDir, hash[0:1]+"/"+hash[1:2]+"/"+hash[2:])
		// todo check exist
		file, _ := os.Open(path)
		defer file.Close()
		http.ServeContent(c.Ctx.ResponseWriter, c.Ctx.Request, "", time.Now(), file)
	}
}
