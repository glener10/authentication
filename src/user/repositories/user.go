package user_repository

import (
	"errors"

	"github.com/glener10/rotating-pairs-back/src/db"
	user_dtos "github.com/glener10/rotating-pairs-back/src/user/dtos"
)

func CreateUser(user user_dtos.CreateUserRequest) (*user_dtos.CreateUserResponse, error) {
	db := db.GetDB()

	query := "INSERT INTO app.users (email, password) VALUES ($1, $2) RETURNING id"
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
