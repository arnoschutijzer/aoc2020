package daytwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPasswordsFromInput(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_input.txt")
	assert.Equal(t, 2, CalculateValidPasswords(passwords, SledRule))
}

func TestValidPasswords(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_personal.txt")
	assert.Equal(t, 396, CalculateValidPasswords(passwords, SledRule))
}

func TestValidPasswordsFromInputPartTwo(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_input.txt")
	assert.Equal(t, 1, CalculateValidPasswords(passwords, TobogganRule))
}

func TestValidPasswordsPartTwo(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_personal.txt")
	assert.Equal(t, 428, CalculateValidPasswords(passwords, TobogganRule))
}
