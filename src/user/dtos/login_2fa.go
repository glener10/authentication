package user_dtos

type Login2FARequest struct {
	Code string `example:"random2FACode"`
}
