package utils_validators

import (
	"testing"

	"gotest.tools/v3/assert"
)

type DataProviderIsValidEmail struct {
	InvalidEmail string
}

func TestIsValidEmailWithInvalidsEmails(t *testing.T) {
	dataProviderWeakPassword := []*DataProviderIsValidEmail{
		{
			InvalidEmail: "a",
		},
		{
			InvalidEmail: "a@",
		},
		{
			InvalidEmail: "a.com",
		},
		{
			InvalidEmail: "abc@.com",
		},
	}

	for _, data := range dataProviderWeakPassword {
		assert.Equal(t, IsValidEmail(data.InvalidEmail), false, "should return false to invalid email '"+data.InvalidEmail+"'")
	}
}

func TestIsValidEmailWithValidEmail(t *testing.T) {
	assert.Equal(t, IsValidEmail("fulano@fulano.com"), true, "should return true to valid email")
}
