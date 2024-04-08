package user_dtos

type UserWithoutSensitiveData struct {
	Id       int    `example:"1"`
	Email    string `example:"fulano@fulano.com"`
	IsAdmin  *bool  `example:"true"`
	Inactive *bool  `example:"false"`
}
