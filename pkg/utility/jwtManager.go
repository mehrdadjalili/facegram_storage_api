package utility

import (
	"errors"
	"facegram_file_server/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type JwtClaims struct {
	UserData string `json:"user_data"`
	jwt.StandardClaims
}

type LoginTokensTypesOutput struct {
	Token string `json:"token"`
	Kind  string `json:"kind"`
}

type LoginTokensOutput struct {
	PageID string `json:"page_id"`
	Token  string `json:"token"`
}

func JwtTokenBuilder(data string) (string, error) {
	key := []byte(*config.GetExternalKeyConfig())
	expTime := time.Now().Add(30 * 24 * 60 * time.Minute)
	pageClaims := &JwtClaims{
		UserData: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	tryToken := jwt.NewWithClaims(jwt.SigningMethodHS256, pageClaims)
	token, err := tryToken.SignedString(key)
	if err != nil {
		return "", errors.New("error")
	}
	return token, nil
}
