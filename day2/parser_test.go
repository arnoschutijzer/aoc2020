package daytwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	passwords := ReadFile("./fixtures/parser_test_1.txt")

	var expected = []Password{
		{
			letter:       "a",
			maxOccurence: 3,
			minOccurence: 1,
			password:     "abcde",
		},
		{
			letter:       "b",
			maxOccurence: 3,
			minOccurence: 1,
			password:     "cdefg",
		},
		{
			letter:       "c",
			maxOccurence: 9,
			minOccurence: 2,
			password:     "ccccccccc",
		},
	}

	assert.Equal(t, passwords, expected)
}
