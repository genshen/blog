package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/genshen/blog/components/utils"
)

const LenAccessTokenName = 13 //access_token=
const OAuth2JwtTokenLifetime = 1200
const OAuth2JwtIssuer = "oauth2_issuer"

const (
	UserStatusUnAuth = iota
	UserStatusHasAuth
)

type User struct {
	Status    int    `bson:"-" json:"status"`
	Name      string `json:"name"`
	HtmlUrl   string `json:"html_url"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	jwt.StandardClaims
}


// jwt interface implementation

func (u *User) JwtCreator(issuer string, expireAt int64) {
	u.ExpiresAt = expireAt
	u.Issuer = issuer
}

func (u *User) ToJwtClaims(token *jwt.Token) (utils.JwtClaims, error) {
	if claims, ok := token.Claims.(*User); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("the utils.JwtClaims interface is not implemented")
}

//type User struct {
//	Status int    `bson:"-" json:"status"`
//	Name   string `json:"name"`
//	Url    string `json:"url"`
//	Email  string `json:"-"`
//	Avatar string `json:"avatar"`
//}

type AuthInterface interface {
	ConvertToUser() User // convert interface to User (maybe it is useless).
	GetAccessToken(code string) string
	GetAuthUserInfo(accessToken string) error
}

func StartAuth(a AuthInterface, code string) error {
	if accessToken := a.GetAccessToken(code); accessToken == "" {
		return errors.New("can't fetch access_token")
	} else {
		return a.GetAuthUserInfo(accessToken)
	}
}
