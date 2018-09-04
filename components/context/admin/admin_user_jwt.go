package admin

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/genshen/blog/components/utils"
)

const JwtAdminConfigQueryTokenKey = "_jwt"
const AdminConfigJwtTokenLifetime = 1200

type UserInfo struct {
	Username string
	Email    string
	jwt.StandardClaims
}

func (user *UserInfo) JwtCreator(issuer string, expireAt int64) {
	user.ExpiresAt = expireAt
	user.Issuer = issuer
}

func (user *UserInfo) ToJwtClaims(token *jwt.Token) (utils.JwtClaims, error) {
	if claims, ok := token.Claims.(*UserInfo); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("the utils.JwtClaims interface is not implemented")
}
