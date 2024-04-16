package user_entities

type User struct {
	Id            int
	Email         string
	Password      string
	IsAdmin       *bool
	Inactive      *bool
	VerifiedEmail *bool
}
