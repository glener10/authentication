package user_repositories

import (
	"database/sql"
	"errors"
	"time"

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
		err = r.Db.QueryRow("SELECT id, email, is_admin, password, inactive, verified_email, code_verify_email, code_verify_email_expiry, code_change_email, code_change_email_expiry, password_recovery_code, password_recovery_code_expiry, twofa, twofa_secret FROM users WHERE email = $1", find).Scan(&user.Id, &user.Email, &user.IsAdmin, &user.Password, &user.Inactive, &user.VerifiedEmail, &user.CodeVerifyEmail, &user.CodeVerifyEmailExpiry, &user.CodeChangeEmail, &user.CodeChangeEmailExpiry, &user.PasswordRecoveryCode, &user.PasswordRecoveryCodeExpiry, &user.Twofa, &user.TwofaSecret)
	} else {
		err = r.Db.QueryRow("SELECT id, email, is_admin, password, inactive, verified_email, code_verify_email, code_change_email, password_recovery_code, twofa, twofa_secret FROM users WHERE id = $1", find).Scan(&user.Id, &user.Email, &user.IsAdmin, &user.Password, &user.Inactive, &user.VerifiedEmail, &user.CodeVerifyEmail, &user.CodeChangeEmail, &user.PasswordRecoveryCode, &user.Twofa, &user.TwofaSecret)
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

func (r *SQLRepository) UpdateEmailVerificationCode(find string, code string, expiration time.Time) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET code_verify_email = $1, code_verify_email_expiry = $2 WHERE email = $3 RETURNING id, email", code, expiration, find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET code_verify_email = $1, code_verify_email_expiry = $2 WHERE id = $3 RETURNING id, email", code, expiration, find).Scan(&user.Id, &user.Email)
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

func (r *SQLRepository) CheckCodeVerifyEmail(find string, code string) (*bool, error) {
	var codeInDb string
	var expire time.Time
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("SELECT code_verify_email, code_verify_email_expiry FROM users WHERE email = $1", find).Scan(&codeInDb, &expire)
	} else {
		err = r.Db.QueryRow("SELECT code_verify_email, code_verify_email_expiry FROM users WHERE id = $1", find).Scan(&codeInDb, &expire)
	}
	if err != nil {
		return nil, errors.New("no element with the parameter (id/email) '" + find + "'")
	}

	if code != codeInDb {
		return nil, errors.New("your code is invalid")
	}

	if expire.Before(time.Now()) {
		return nil, errors.New("your code has expired")
	}

	successBool := true
	return &successBool, nil
}

func (r *SQLRepository) ResetEmailVerificationCode(find string) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET code_verify_email = NULL, code_verify_email_expiry = NULL WHERE email = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET code_verify_email = NULL, code_verify_email_expiry = NULL WHERE id = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
	}
	if err != nil {
		return nil, errors.New("error to reset verify email code in repository with the parameter (id/email) '" + find + "'")
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	return &userWithoutSensitiveData, nil
}

func (r *SQLRepository) UpdatePasswordRecoveryCode(find string, code string, expiration time.Time) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET password_recovery_code = $1, password_recovery_code_expiry = $2 WHERE email = $3 RETURNING id, email", code, expiration, find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET password_recovery_code = $1, password_recovery_code_expiry = $2 WHERE id = $3 RETURNING id, email", code, expiration, find).Scan(&user.Id, &user.Email)
	}
	if err != nil {
		return nil, errors.New("error to update password recovery code in repository with the parameter (id/email) '" + find + "'")
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	return &userWithoutSensitiveData, nil
}

func (r *SQLRepository) CheckPasswordRecoveryCode(find string, code string) (*bool, error) {
	var codeInDb string
	var expire time.Time
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("SELECT password_recovery_code, password_recovery_code_expiry FROM users WHERE email = $1", find).Scan(&codeInDb, &expire)
	} else {
		err = r.Db.QueryRow("SELECT password_recovery_code, password_recovery_code_expiry FROM users WHERE id = $1", find).Scan(&codeInDb, &expire)
	}
	if err != nil {
		return nil, errors.New("no element with the parameter (id/email) '" + find + "'")
	}

	if code != codeInDb {
		return nil, errors.New("your code is invalid")
	}

	if expire.Before(time.Now()) {
		return nil, errors.New("your code has expired")
	}

	successBool := true
	return &successBool, nil
}

func (r *SQLRepository) ResetPasswordRecoveryCode(find string) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET password_recovery_code = NULL, password_recovery_code_expiry = NULL WHERE email = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET password_recovery_code = NULL, password_recovery_code_expiry = NULL WHERE id = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
	}
	if err != nil {
		return nil, errors.New("error to reset password recovery code in repository with the parameter (id/email) '" + find + "'")
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	return &userWithoutSensitiveData, nil
}

func (r *SQLRepository) UpdateChangeEmailCode(find string, code string, expiration time.Time) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET code_change_email = $1, code_change_email_expiry = $2 WHERE email = $3 RETURNING id, email", code, expiration, find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET code_change_email = $1, code_change_email_expiry = $2 WHERE id = $3 RETURNING id, email", code, expiration, find).Scan(&user.Id, &user.Email)
	}
	if err != nil {
		return nil, errors.New("error to update change email code in repository with the parameter (id/email) '" + find + "'")
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	return &userWithoutSensitiveData, nil
}

func (r *SQLRepository) CheckChangeEmailCode(find string, code string) (*bool, error) {
	var codeInDb string
	var expire time.Time
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("SELECT code_change_email, code_change_email_expiry FROM users WHERE email = $1", find).Scan(&codeInDb, &expire)
	} else {
		err = r.Db.QueryRow("SELECT code_change_email, code_change_email_expiry FROM users WHERE id = $1", find).Scan(&codeInDb, &expire)
	}
	if err != nil {
		return nil, errors.New("no element with the parameter (id/email) '" + find + "'")
	}

	if code != codeInDb {
		return nil, errors.New("your code is invalid")
	}

	if expire.Before(time.Now()) {
		return nil, errors.New("your code has expired")
	}

	successBool := true
	return &successBool, nil
}

func (r *SQLRepository) ResetChangeEmailCode(find string) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET code_change_email = NULL, code_change_email_expiry = NULL WHERE email = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET code_change_email = NULL, code_change_email_expiry = NULL WHERE id = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
	}
	if err != nil {
		return nil, errors.New("error to reset password recovery code in repository with the parameter (id/email) '" + find + "'")
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	return &userWithoutSensitiveData, nil
}

func (r *SQLRepository) Active2FA(find string, secret string) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET twofa = true, twofa_secret = $1 WHERE email = $2 RETURNING id, email", secret, find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET twofa = true, twofa_secret = $1 WHERE id = $2 RETURNING id, email", secret, find).Scan(&user.Id, &user.Email)
	}
	if err != nil {
		return nil, errors.New("error to active 2FA in repository with the parameter (id/email) '" + find + "'")
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	return &userWithoutSensitiveData, nil
}

func (r *SQLRepository) Desactive2FA(find string) (*user_dtos.UserWithoutSensitiveData, error) {
	var user user_entity.User
	var err error
	if utils_validators.IsValidEmail(find) {
		err = r.Db.QueryRow("UPDATE users SET twofa = false, twofa_secret = null WHERE email = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
	} else {
		err = r.Db.QueryRow("UPDATE users SET twofa = false, twofa_secret = null WHERE id = $1 RETURNING id, email", find).Scan(&user.Id, &user.Email)
	}
	if err != nil {
		return nil, errors.New("error to active 2FA in repository with the parameter (id/email) '" + find + "'")
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	return &userWithoutSensitiveData, nil
}
