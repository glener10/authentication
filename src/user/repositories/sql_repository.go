package user_repositories

import (
	"database/sql"
	"errors"

	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_entity "github.com/glener10/authentication/src/user/entities"
)

type SQLRepository struct {
	Db *sql.DB
}

func (r *SQLRepository) CreateUser(user user_dtos.CreateUserRequest) (*user_dtos.CreateUserResponse, error) {
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"
	var pk int
	err := r.Db.QueryRow(query, user.Email, user.Password).Scan(&pk)
	if err != nil {
		return nil, errors.New("Error creating user: " + err.Error())
	}
	object := user_dtos.CreateUserResponse{
		Id:    pk,
		Email: user.Email,
	}
	return &object, nil
}

func (r *SQLRepository) FindByEmail(email string) (*user_entity.User, error) {
	var user user_entity.User
	err := r.Db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.New("error to find by email: " + email)
	}
	return &user, nil
}
