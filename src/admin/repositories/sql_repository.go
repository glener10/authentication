package admin_repositories

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

func (r *SQLRepository) PromoteUserAdmin(find string) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET is_admin = true WHERE email = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET is_admin = true WHERE id = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
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

func (r *SQLRepository) FindAllUsers() ([]*user_dtos.UserWithoutSensitiveData, error) {
	rows, err := r.Db.Query("SELECT id, email, is_admin FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*user_dtos.UserWithoutSensitiveData
	for rows.Next() {
		var user user_dtos.UserWithoutSensitiveData
		if err := rows.Scan(&user.Id, &user.Email, &user.IsAdmin); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
