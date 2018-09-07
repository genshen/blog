package auth

import (
	"github.com/astaxie/beego/httplib"
	"github.com/genshen/blog/components/keys"
	"strings"
)

type GithubAuthUser struct {
	User
}

func (g *GithubAuthUser) ConvertToUser() User {
	return g.User
}

func (g *GithubAuthUser) GetAccessToken(code string) string {
	req := httplib.Post(keys.GitHubKey.AuthUrl)
	req.Param("client_id", keys.GitHubKey.ClientId)
	req.Param("client_secret", keys.GitHubKey.ClientSecret)
	req.Param("code", code)
	req.Param("accept", "json")

	response, err := req.String()
	if err == nil {
		var index = strings.Index(response, "access_token=")
		if index >= 0 && len(response) > index+LenAccessTokenName {
			var after = response[index+LenAccessTokenName:]
			index = strings.Index(after, "&")
			if index < 0 {
				return after
			}
			return after[:index]
		}
	}
	return ""
}

func (g *GithubAuthUser) GetAuthUserInfo(accessToken string) (error) {
	req := httplib.Get("https://api.github.com/user?access_token=" + accessToken)
	err := req.ToJSON(&g)
	return err
}
