package xilution

import (
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dgrijalva/jwt-go"
)

func buildTestId() string {
	return strings.Join(strings.Split(gofakeit.UUID(), "-"), "")
}

func buildJwtToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})
	tokenString, _ := token.SignedString([]byte("AllYourBase"))

	return tokenString
}
