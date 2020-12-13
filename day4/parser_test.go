package dayfour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsesSinglePassport(t *testing.T) {
	expected := []Passport{
		{
			stringified: "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm",
		},
	}
	assert.Equal(t, expected, ParsePassports("./fixtures/single_passport.txt"))
}

func TestParsesMultiplePassports(t *testing.T) {
	expected := []Passport{
		{
			stringified: "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm",
		},
		{
			stringified: "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929",
		},
	}
	assert.Equal(t, expected, ParsePassports("./fixtures/two_passports.txt"))
}
