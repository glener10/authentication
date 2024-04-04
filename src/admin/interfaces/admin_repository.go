package admin_interfaces

import (
	user_dtos "github.com/glener10/authentication/src/user/dtos"
)

type IAdminRepository interface {
	PromoteUserAdmin(find string) (*user_dtos.UserWithoutSensitiveData, error)
}
