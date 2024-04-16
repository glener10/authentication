package user_interfaces

import (
	"time"

	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_entity "github.com/glener10/authentication/src/user/entities"
)

type IUserRepository interface {
	CreateUser(user user_dtos.CreateUserRequest) (*user_dtos.UserWithoutSensitiveData, error)
	FindUser(find string) (*user_entity.User, error)
	ChangePassword(find string, newPassword string) (*user_dtos.UserWithoutSensitiveData, error)
	ChangeEmail(find string, newEmail string) (*user_dtos.UserWithoutSensitiveData, error)
	DeleteUser(find string) error
	UpdateEmailVerificationCode(find string, code string, expiration time.Time) (*user_dtos.UserWithoutSensitiveData, error)
}
