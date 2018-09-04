package keys

import (
	"github.com/astaxie/beego/logs"
	"github.com/genshen/blog/components/utils"
)

var GitHubKey struct {
	AuthUrl      string
	ClientId     string
	ClientSecret string
}

func loadGithubKeys() {
	if utils.CustomConfig.Auth.Keys != nil {
		githubConfig := utils.CustomConfig.Auth.Keys["github"]
		GitHubKey.AuthUrl = githubConfig.AuthUrl
		GitHubKey.ClientId = githubConfig.ClientId
		GitHubKey.ClientSecret = githubConfig.SecretId
	} else {
		logs.Warning("github key not found in config.")
	}
}
