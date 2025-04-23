package service

import (
	auth "AUTHEN-AUTHOSERVER/internal/auth"
	model "AUTHEN-AUTHOSERVER/internal/model"
	repository "AUTHEN-AUTHOSERVER/internal/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)
func RegisterUser(username, password, email string) error {
	user := &model.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	return repository.CreateUser(user)
}
func Login(username, password string) (string, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password")
	}
	return auth.GenerateJWT(user.ID)
}