package utils

import (
	"fmt"
	"time"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

const (
	ConfigJwtTokenLifetime      = 600
	ConfigJwtSecret             = "ds"
	JwtAdminConfigQueryTokenKey = "_jwt"
)

type UserInfo struct {
	Username string
	Email    string
}

type Claims struct {
	UserInfo
	jwt.StandardClaims
}

// create a jwt token,and return this token as string type.
func JwtNewToken(connection UserInfo, issuer string) (tokenString string, expire int64, err error) {
	expireToken := time.Now().Add(time.Second * time.Duration(ConfigJwtTokenLifetime)).Unix()

	// We'll manually assign the claims but in production you'd insert values from a database
	claims := Claims{
		UserInfo: connection,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	if signedToken, err := token.SignedString([]byte(ConfigJwtSecret)); err != nil {
		return "", 0, err
	} else {
		return signedToken, expireToken, nil
	}
}

// Verify a jwt token
func JwtVerify(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure token's signature wasn't changed
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected siging method")
		}
		return []byte(ConfigJwtSecret), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("unauthenticated")
}
