package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SIGNING_KEY = []byte(os.Getenv("secret"))

func GenerateToken(username, password string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user"] = username
	claims["password"] = password
	claims["iat"] = time.Now().Unix()
	claims["eat"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SIGNING_KEY)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidToken(tokenString string) bool {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("There was an error parsing")
		}
		return SIGNING_KEY, nil
	})
	return err != nil
}
