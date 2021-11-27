package utils

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"

	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func getSecret() []byte {
	var env map[string]string
	env, err := godotenv.Read("./.env")
	if err != nil {
		log.Fatal("Errore nel caricare il file .env:", err)
	}
	return []byte(env["SECRET"])
}

type Jwt struct {
	CodFisc string
	Date    time.Time
}

func (jwt *Jwt) GenerateToken() (string, error) {
	token := jwtLib.New(jwtLib.SigningMethodHS256)
	claims := token.Claims.(jwtLib.MapClaims)

	claims["authorized"] = true
	claims["user"] = jwt.CodFisc
	claims["exp"] = jwt.Date.Add(time.Minute * 600000).Unix() // 10k ore: circa 416 giorni

	tokenString, err := token.SignedString(getSecret())
	if err != nil {
		fmt.Println(err)
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
		return string(getSecret()), nil
	})
	return token.Valid, err
}

func DecryptBasic(b string) ([]string, error) {
	toDecode, err := base64.URLEncoding.DecodeString(b)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(toDecode), ":"), nil
}
