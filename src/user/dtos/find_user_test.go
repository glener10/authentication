package user_dtos

import (
	"testing"

	"github.com/glener10/authentication/tests"
	"github.com/stretchr/testify/assert"
)

func TestSendErrorIfParamIsEmptyOrNull(t *testing.T) {
	err := ValidateFindUser("")
	assert.Equal(t, err.Error(), "find parameter is required", "find parameter should be required")
}

func TestSendErrorIfParamIsNotIntegerAndEmail(t *testing.T) {
	err := ValidateFindUser("invalidParameter")
	assert.Equal(t, err.Error(), "wrong format, parameter need to be a id or a e-mail")
}

func TestSuccessWhenParamIsInteger(t *testing.T) {
	err := ValidateFindUser("1")
	assert.NoError(t, err)
}

func TestSuccessWhenParamIsEmail(t *testing.T) {
	err := ValidateFindUser(tests.ValidEmail)
	assert.NoError(t, err)
}
