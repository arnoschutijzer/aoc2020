package daytwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPasswords(t *testing.T) {
	assert.Panics(t, CalculateValidPasswords)
}
