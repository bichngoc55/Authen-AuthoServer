package repository

import (
	"AUTHEN-AUTHOSERVER/internal/db"
	models "AUTHEN-AUTHOSERVER/internal/model"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error { 
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
 
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
	_, err = db.DB.Exec(query, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}

	return nil
}
func GetUserByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, password, email FROM users WHERE username = $1`
	user := &models.User{}
	err := db.DB.Get(user, query, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
