package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"os"
)

//const SecretKey =  "secret"//SECRET_KEY
type Claims struct{
	jwt.StandardClaims
}

func GenerateJwt(issuer string) (string,error){
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.StandardClaims{
		Issuer: issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	return claims.SignedString([]byte( os.Getenv("SECRET_KEY")))
}

func ParseJwt (cookie string)(string,error){
	token, err := jwt.ParseWithClaims(cookie,&jwt.StandardClaims{},func(t *jwt.Token) (interface{}, error) {
		return []byte( os.Getenv("SECRET_KEY")),nil 
	})

	if err != nil || !token.Valid{
		return "",err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer,nil
}