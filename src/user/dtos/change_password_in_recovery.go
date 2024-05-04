package user_dtos

type ChangePasswordInRecoveryRequest struct {
	NewPassword string `validate:"required" example:"aaaaaaaA#1"`
	Code        string `validate:"required" example:"123456"`
}
