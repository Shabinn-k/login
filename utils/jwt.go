package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var ACCESS_TOKEN = []byte(os.Getenv("ACCESS_SECRET"))
var REFRESH_TOKEN = []byte(os.Getenv("REFRESH_SECRET"))

func GenerateAccessToken(userId uint, role string) (string,error) {
	claims:=jwt.MapClaims{
		"user_id":userId,
		"role":role,
		"exp":time.Now().Add(time.Minute*15).Unix(),		
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(ACCESS_TOKEN)
}

func GenerateRefreshToken(userId uint) (string,error){
	claims:=jwt.MapClaims{
		"user_id":userId,
		"exp":time.Now().Add(time.Hour*24*7).Unix(),
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(REFRESH_TOKEN)
}

func VerifyAccessToken(tokenStr string)(*jwt.Token,error){
	token,err:=jwt.Parse(tokenStr,func(t *jwt.Token)(interface{},error){
		return ACCESS_TOKEN,nil
	})
	if err!=nil ||!token.Valid{
		return nil,errors.New("Invalid token")
	}
	return token,nil
}