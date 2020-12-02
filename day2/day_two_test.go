package daytwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPasswordsFromInput(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_input.txt")
	assert.Equal(t, 2, CountValidPasswords(passwords, SledRule{}))
}

func TestValidPasswords(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_personal.txt")
	assert.Equal(t, 396, CountValidPasswords(passwords, SledRule{}))
}

func TestValidPasswordsFromInputPartTwo(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_input.txt")
	assert.Equal(t, 1, CountValidPasswords(passwords, TobogganRule{}))
}

func TestValidPasswordsPartTwo(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_personal.txt")
	assert.Equal(t, 428, CountValidPasswords(passwords, TobogganRule{}))
}
