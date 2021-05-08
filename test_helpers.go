package xilution

import (
	b64 "encoding/base64"
	"strings"
	"time"

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

func buildTestOrganization() Organization {
	return Organization{
		Type:           "organization",
		ID:             buildTestId(),
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		Name:           gofakeit.Company(),
		Logo:           b64.StdEncoding.EncodeToString([]byte(string(gofakeit.ImagePng(50, 50)))),
		Domain:         gofakeit.DomainName(),
		IamClientId:    buildTestId(),
		OrganizationId: buildTestId(),
		Active:         gofakeit.Bool(),
	}
}

func buildTestClient() Client {
	return Client{
		Type:           "client",
		ID:             buildTestId(),
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		Name:           gofakeit.Word(),
		Grants:         []string{"password", "client_credentials", "authorization_code", "refresh_token"},
		RedirectUris:   []string{gofakeit.URL(), gofakeit.URL(), gofakeit.URL()},
		OrganizationId: buildTestId(),
		Active:         gofakeit.Bool(),
	}
}

func buildTestUser() User {
	return User{
		Type:           "user",
		ID:             buildTestId(),
		OwningUserId:   buildTestId(),
		CreatedAt:      gofakeit.Date().Format(time.RFC3339),
		ModifiedAt:     gofakeit.Date().Format(time.RFC3339),
		FirstName:      gofakeit.FirstName(),
		LastName:       gofakeit.LastName(),
		Email:          gofakeit.Email(),
		Username:       gofakeit.Username(),
		OrganizationId: buildTestId(),
		Active:         gofakeit.Bool(),
	}
}
