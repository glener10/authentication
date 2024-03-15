package user_repository

import (
	"errors"

	"github.com/glener10/rotating-pairs-back/src/db"
	user_dtos "github.com/glener10/rotating-pairs-back/src/user/dtos"
	user_entity "github.com/glener10/rotating-pairs-back/src/user/entities"
)

func CreateUser(user user_dtos.CreateUserRequest) (*user_dtos.CreateUserResponse, error) {
	db := db.GetDB()

	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"
	var pk int
	err := db.QueryRow(query, user.Email, user.Password).Scan(&pk)
	if err != nil {
		return nil, errors.New("Error creating user: " + err.Error())
	}
	object := user_dtos.CreateUserResponse{
		Id:    pk,
		Email: user.Email,
	}
	return &object, nil
}

func FindByEmail(email string) (*user_entity.User, error) {
	db := db.GetDB()

	var user user_entity.User
	err := db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.New("error to find by email: " + email)
	}
	return &user, nil
}
