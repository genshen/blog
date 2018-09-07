package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	ConfigJwtSecret = "ds"
)

type JwtClaims interface {
	jwt.Claims
	// Initialize this JwtClaims with issuer and lifetime
	JwtCreator(issuer string, lifetime int64)
	// convert token to JwtClaims
	// It returns parsed pointer of JwtClaims and error if it has.
	ToJwtClaims(token *jwt.Token) (JwtClaims, error)
}

// create a jwt token using given claims data and issuer/ expire duration.
// and return this token as string type.
func NewJwtToken(jwtClaims JwtClaims, issuer string, lifetime int64) (tokenString string, expire int64, err error) {
	expireAt := time.Now().Add(time.Second * time.Duration(lifetime)).Unix()
	jwtClaims.JwtCreator(issuer, expireAt) // just set expire and issuer

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims.(jwt.Claims))

	// Signs the token with a secret.
	if signedToken, err := token.SignedString([]byte(ConfigJwtSecret)); err != nil {
		return "", 0, err
	} else {
		return signedToken, expireAt, nil
	}
}

// Verify a jwt token
// the parameter of jwtClaims can be empty (pointer must be passed)
// It returns parsed pointer of JwtClaims and error if it has.
func JwtVerify(jwtClaims JwtClaims, tokenString string) (JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwtClaims, func(token *jwt.Token) (interface{}, error) {
		// Make sure token's signature wasn't changed
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected siging method")
		}
		return []byte(ConfigJwtSecret), nil
	})
	if err == nil {
		return jwtClaims.ToJwtClaims(token)
	}
	return nil, errors.New("unauthenticated")
}
