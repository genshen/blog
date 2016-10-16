package keys

import "github.com/astaxie/beego"

var GitHubKey  struct {
	AuthUrl      string
	ClientId     string
	ClientSecret string
}

func loadGithubKeys() {
	GitHubKey.AuthUrl = beego.AppConfig.String("github_auth_url")
	GitHubKey.ClientId = beego.AppConfig.String("github_client_id")
	GitHubKey.ClientSecret = beego.AppConfig.String("github_client_secret")
}