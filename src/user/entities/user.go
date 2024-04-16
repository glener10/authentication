package user_entities

import "time"

type User struct {
	Id       int
	Email    string
	Password string

	VerifiedEmail *bool

	IsAdmin  *bool
	Inactive *bool

	CodeVerifyEmail       *string
	CodeVerifyEmailExpiry *time.Time

	CodeChangeEmail       *string
	CodeChangeEmailExpiry *time.Time

	CodeChangePassword       *string
	CodeChangePasswordExpiry *time.Time
}
