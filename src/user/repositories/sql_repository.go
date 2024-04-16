package user_repositories

import (
	"database/sql"
	"errors"

	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_entity "github.com/glener10/authentication/src/user/entities"
	utils_validators "github.com/glener10/authentication/src/utils/validators"
)

type SQLRepository struct {
	Db *sql.DB
}

func (r *SQLRepository) CreateUser(user user_dtos.CreateUserRequest) (*user_dtos.UserWithoutSensitiveData, error) {
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"
	var pk int
	err := r.Db.QueryRow(query, user.Email, user.Password).Scan(&pk)
	if err != nil {
		return nil, errors.New("Error creating user: " + err.Error())
	}
	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    pk,
		Email: user.Email,
	}
	return &userWithoutSensitiveData, nil
}

func (r *SQLRepository) FindUser(find string) (*user_entity.User, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("SELECT id, email, is_admin, password, inactive, verified_email, code_verify_email, code_change_email, code_change_password FROM users WHERE email = $1", find).Scan(&user.Id, &user.Email, &user.IsAdmin, &user.Password, &user.Inactive, &user.VerifiedEmail, &user.CodeVerifyEmail, &user.CodeChangeEmail, &user.CodeChangePassword)
	} else {
		err = r.Db.QueryRow("SELECT id, email, is_admin, password, inactive, verified_email, code_verify_email, code_change_email, code_change_password FROM users WHERE id = $1", find).Scan(&user.Id, &user.Email, &user.IsAdmin, &user.Password, &user.Inactive, &user.VerifiedEmail, &user.CodeVerifyEmail, &user.CodeChangeEmail, &user.CodeChangePassword)
	}
	if err != nil {
		return nil, errors.New("no element with the parameter (id/email) '" + find + "'")
	}
	return &user, nil
}

func (r *SQLRepository) ChangePassword(find string, newPassword string) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET password = $1 WHERE email = $2 RETURNING id, email", newPassword, find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET password = $1 WHERE id = $2 RETURNING id, email", newPassword, find).Scan(&user.Id, &user.Email)
	}
	if err != nil {
		return nil, errors.New("error to change password in repository with the parameter (id/email) '" + find + "'")
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	return &userWithoutSensitiveData, nil
}

func (r *SQLRepository) ChangeEmail(find string, newEmail string) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET email = $1 WHERE email = $2 RETURNING id, email", newEmail, find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET email = $1 WHERE id = $2 RETURNING id, email", newEmail, find).Scan(&user.Id, &user.Email)
	}
	if err != nil {
		return nil, errors.New("error to change email in repository with the parameter (id/email) '" + find + "'")
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	return &userWithoutSensitiveData, nil
}

func (r *SQLRepository) DeleteUser(find string) error {
	var err error
	var result sql.Result
	if utils_validators.IsValidEmail(find) {
		result, err = r.Db.Exec("DELETE FROM users WHERE email=$1", find)
	} else {
		result, err = r.Db.Exec("DELETE FROM users WHERE id=$1", find)
	}
	if err != nil {
		return errors.New("error to delete user in repository with the parameter (id/email) '" + find + "': " + err.Error())
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		return errors.New("error while checking affected rows: " + err.Error())
	}
	if numRows == 0 {
		return errors.New("doesnt exists user with '" + find + "' atribute")
	}
	return nil
}

func (r *SQLRepository) VerifyEmail(find string) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET verified_email = true WHERE email = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET verified_email = true WHERE id = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
	}
	if err != nil {
		return nil, errors.New("error to verify email in repository with the parameter (id/email) '" + find + "'")
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	return &userWithoutSensitiveData, nil
}
