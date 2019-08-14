package service

import (
	getconfig "github.com/donaderoyan/go-graphql-api/config"
	"github.com/donaderoyan/go-graphql-api/app/model"
	"testing"
)

var (
	authService *AuthService
)

func init() {
	config := getconfig.LoadConfig("../")
	log := NewLogger(config)
	authService = NewAuthService(config, log)
}

func TestSignJWT(t *testing.T) {
	user := &model.User{
		ID:       "1",
		Email:    "test@1.com",
		Password: "123456",
	}
	tokenString, err := authService.SignJWT(user)
	if err != nil {
		t.Errorf("Error during signing JWT")
	}
	if *tokenString == "" || tokenString == nil {
		t.Errorf("Invalid JWT")
	}

}
