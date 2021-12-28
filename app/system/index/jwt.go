package index

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const jwtIssuer = "issuer"

const validDuration = 3600 * 24 * 30

var jwtKey = "key"

var (
	invalidTokenErr     = errors.New("token is invalid")
	failToParseTokenErr = errors.New("fail to parse token")
)

type jwtClaims struct {
	AccountId int64
	jwt.StandardClaims
}

func generateToken(accountId int64) (string, error) {
	nowUnix := time.Now().Unix()
	claims := jwtClaims{
		AccountId: accountId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: nowUnix,
			IssuedAt:  nowUnix,
			ExpiresAt: nowUnix + validDuration,
			Issuer:    jwtIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}

func parseToken(tokenStr string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, invalidTokenErr
	}
	if claims, ok := token.Claims.(*jwtClaims); ok {
		return claims, nil
	}
	return nil, failToParseTokenErr
}
