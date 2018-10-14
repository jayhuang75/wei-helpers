package wei

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	errInfo := Logging("info", "test")
	assert.Equal(t, errInfo, nil)

	errError := Logging("error", "error")
	assert.Equal(t, errError, nil)
}
