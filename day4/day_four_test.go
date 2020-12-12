package dayfour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatesSinglePassport(t *testing.T) {
	passports := ParsePassports("./fixtures/one_valid_passport.txt")
	validPassports := OnlyValidPassports(passports)
	assert.Equal(t, 1, len(validPassports))
}

func TestValidatesMultiplePassports(t *testing.T) {
	passports := ParsePassports("./fixtures/two_valid_passports.txt")
	validPassports := OnlyValidPassports(passports)
	assert.Equal(t, 2, len(validPassports))
}

func TestValidatesPartOneInput(t *testing.T) {
	passports := ParsePassports("./fixtures/part_one_input.txt")
	validPassports := OnlyValidPassports(passports)
	assert.Equal(t, 254, len(validPassports))
}
