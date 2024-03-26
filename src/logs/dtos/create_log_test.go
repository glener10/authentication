package logs_dtos

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2024-03-26T00:00:00Z")
	if err != nil {
		t.Error("error to convert timestamp:", err)
	}

	dto := &CreateLogRequest{
		UserID:        1,
		Success:       true,
		OperationCode: "LOGIN",
		Ip:            "192.168.0.1",
		Timestamp:     timestamp,
	}
	err = Validate(dto)
	assert.NoError(t, err)
}
