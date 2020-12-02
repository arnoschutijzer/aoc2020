package daytwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPasswordsFromInput(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_input.txt")
	assert.Equal(t, 2, CalculateValidPasswords(passwords))
}

func TestValidPasswords(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_personal.txt")
	assert.Equal(t, 396, CalculateValidPasswords(passwords))
}
