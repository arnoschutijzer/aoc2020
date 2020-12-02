package daytwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPasswordsFromInput(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_input.txt")
	assert.Equal(t, 2, CalculateValidPasswords(passwords, PartOne))
}

func TestValidPasswords(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_personal.txt")
	assert.Equal(t, 396, CalculateValidPasswords(passwords, PartOne))
}

func TestValidPasswordsFromInputPartTwo(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_input.txt")
	assert.Equal(t, 1, CalculateValidPasswords(passwords, PartTwo))
}

func TestValidPasswordsPartTwo(t *testing.T) {
	passwords := ReadFile("./fixtures/daytwo_personal.txt")
	assert.Equal(t, 428, CalculateValidPasswords(passwords, PartTwo))
}
