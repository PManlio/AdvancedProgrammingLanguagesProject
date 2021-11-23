package utils

import (
	"fmt"
	"log"
	"time"

	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var secret = func() string {
	var env map[string]string
	env, err := godotenv.Read("./.env")
	if err != nil {
		log.Fatal("Errore nel caricare il file .env:", err)
	}
	return env["SECRET"]
}

type Jwt struct {
	username string
	date     time.Time
}

func (jwt *Jwt) Generate() (string, error) {
	token := jwtLib.New(jwtLib.SigningMethodHS256)
	claims := token.Claims.(jwtLib.MapClaims)

	claims["authorized"] = true
	claims["user"] = jwt.username
	claims["exp"] = jwt.date.Add(time.Minute * 600000).Unix() // 10k ore: circa 416 giorni

	tokenString, err := token.SignedString(secret)
	if err != nil {
		fmt.Errorf("Something went wrong during JWT creation: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func IsJWTTokenValid(tokenFromHeader string) (bool, error) {
	token, err := jwtLib.Parse(tokenFromHeader, func(t *jwtLib.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwtLib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error during JWT Decode")
		}
		return secret, nil
	})
	return token.Valid, err
}
