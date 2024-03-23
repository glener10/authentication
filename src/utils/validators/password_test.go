package utils_validators

import (
	"testing"

	"gotest.tools/v3/assert"
)

type DataProviderWeakPasswordType struct {
	WeakPassword   string
	ExpectedReturn string
}

func TestIsStrongPasswordWithWeaksPasswords(t *testing.T) {
	dataProviderWeakPassword := []*DataProviderWeakPasswordType{
		{
			WeakPassword:   "a",
			ExpectedReturn: "the password must be at least 8 characters long",
		},
		{
			WeakPassword:   "AAAAAAAA",
			ExpectedReturn: "the password must be at least 1 lowercase character",
		},
		{
			WeakPassword:   "aaaaaaaa",
			ExpectedReturn: "the password must be at least 1 uppercase character",
		},
		{
			WeakPassword:   "aaaaaaaA",
			ExpectedReturn: `the password must be at least 1 special character: [!@#$%^&*()\-_=+{}[\]:;'"<>,.?/\\|]`,
		},
		{
			WeakPassword:   "aaaaaaaA:",
			ExpectedReturn: "the password must be at least 1 number",
		},
	}

	for _, data := range dataProviderWeakPassword {
		assert.Equal(t, IsStrongPassword(data.WeakPassword).Error(), data.ExpectedReturn, "should return the error message '"+data.ExpectedReturn+"' to weak password '"+data.WeakPassword+"'")
	}
}

func TestIsStrongPasswordWithStrongPassword(t *testing.T) {
	assert.Equal(t, IsStrongPassword("aaaaaA#7"), nil, "should return nil to a strong password")
}
