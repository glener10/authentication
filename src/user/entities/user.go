package user_entities

type User struct {
	Id       int
	Email    string
	Password string

	VerifiedEmail *bool

	IsAdmin  *bool
	Inactive *bool

	CodeVerifyEmail    *bool
	CodeChangeEmail    *bool
	CodeChangePassword *bool
}
