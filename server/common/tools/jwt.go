package tools

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	Uid int64
	jwt.RegisteredClaims
}

type CustomClaims struct {
	ID interface{}
	jwt.RegisteredClaims
}

func GenCustomToken(id interface{}, expiresTime int64) (string, error) {
	claim := CustomClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Truncate(time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiresTime) * time.Second)),
		},
	}
	return genCustomTokenWithClaim(claim)
}

func genCustomTokenWithClaim(claim jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte("goligoli"))
}

func GenToken(uid int64) (string, error) {
	claim := MyClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Truncate(time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(114514) * time.Second)),
		},
	}
	return genCustomTokenWithClaim(claim)
}

func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("goligoli"), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*MyClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid Token")
}

func ParseCustomToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("goligoli"), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid Token")
}

func GetUID(ctx *app.RequestContext) (int64, error) {
	uid, err := ctx.Get("UID")
	if !err {
		return 0, errors.New("no uid")
	}
	return uid.(int64), nil
}
