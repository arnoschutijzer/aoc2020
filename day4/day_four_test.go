package dayfour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatesSinglePassport(t *testing.T) {
	passports := ParsePassports("./fixtures/one_valid_passport.txt")
	validPassports := OnlyPassportsThatContainAllFields(passports)
	assert.Equal(t, 1, len(validPassports))
}

func TestValidatesMultiplePassports(t *testing.T) {
	passports := ParsePassports("./fixtures/two_valid_passports.txt")
	validPassports := OnlyPassportsThatContainAllFields(passports)
	assert.Equal(t, 2, len(validPassports))
}

func TestValidatesPartOneInput(t *testing.T) {
	passports := ParsePassports("./fixtures/part_one_input.txt")
	validPassports := OnlyPassportsThatContainAllFields(passports)
	assert.Equal(t, 254, len(validPassports))
}

func TestHasValidID(t *testing.T) {
	assert.True(t, hasValidID("pid:123456789"))
	assert.True(t, hasValidID("pid:123456789 smth:else"))
	assert.True(t, hasValidID("eyr:2020\npid:123456789 smth:else"))
	assert.False(t, hasValidID("pid:12345w789"))
	assert.False(t, hasValidID("pid:1"))
}

func TestHasValidExpirationYear(t *testing.T) {
	assert.True(t, hasValidExpirationYear("eyr:2020"))
	assert.True(t, hasValidExpirationYear("eyr:2030 pid:1234"))
	assert.False(t, hasValidExpirationYear("\neyr:20"))
	assert.False(t, hasValidExpirationYear("eyr:2019"))
	assert.False(t, hasValidExpirationYear("eyr:2031"))
}

func TestHasValidIssueYear(t *testing.T) {
	assert.True(t, hasValidIssueYear("iyr:2010"))
	assert.True(t, hasValidIssueYear("iyr:2020 pid:1234"))
	assert.False(t, hasValidIssueYear("\niyr:20"))
	assert.False(t, hasValidIssueYear("iyr:2009"))
	assert.False(t, hasValidIssueYear("iyr:2021"))
}

func TestHasValidHeight(t *testing.T) {
	assert.True(t, hasValidHeight("hgt:150cm"))
	assert.True(t, hasValidHeight("hgt:193cm"))
	assert.True(t, hasValidHeight("hgt:59in"))
	assert.True(t, hasValidHeight("hgt:76in"))
	assert.False(t, hasValidHeight("hgt:149cm"))
	assert.False(t, hasValidHeight("hgt:194cm"))
	assert.False(t, hasValidHeight("hgt:58in"))
	assert.False(t, hasValidHeight("hgt:77in"))
}

func TestHasValidEyeColor(t *testing.T) {
	assert.True(t, hasValidEyeColor("ecl:brn"))
	assert.True(t, hasValidEyeColor("ecl:amb"))
	assert.True(t, hasValidEyeColor("ecl:gry"))
	assert.True(t, hasValidEyeColor("ecl:oth "))
	assert.False(t, hasValidEyeColor("hcl:dkooqd"))
	assert.False(t, hasValidEyeColor("something else"))
	assert.False(t, hasValidEyeColor("ecl:123"))
	assert.False(t, hasValidEyeColor("ecl:"))
}

func TestHasValidHairColor(t *testing.T) {
	assert.True(t, hasValidHairColor("hcl:#123abc"))
	assert.False(t, hasValidHairColor("hcl:#123abz"))
	assert.False(t, hasValidHairColor("hcl:123abc"))
}

func TestValidatesFourPassports(t *testing.T) {
	passports := ParsePassports("./fixtures/four_valid_passports.txt")
	validPassports := OnlyPassportsThatComplyWithRules(passports)
	assert.Equal(t, 4, len(validPassports))
}

func TestValidatesPartTwoInput(t *testing.T) {
	passports := ParsePassports("./fixtures/part_one_input.txt")
	validPassports := OnlyPassportsThatComplyWithRules(passports)
	assert.Equal(t, 184, len(validPassports))
}
