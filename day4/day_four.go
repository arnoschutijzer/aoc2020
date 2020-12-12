package dayfour

import (
	"strings"
)

type Passport struct {
	stringified string
	fields      map[string]string

	pid string
}

var validFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func (passport *Passport) checkValidity() bool {
	for _, field := range validFields {
		if !strings.Contains(passport.stringified, field) {
			return false
		}
	}

	return true
}

func OnlyValidPassports(passports []Passport) []Passport {
	var validPassports []Passport
	for _, passport := range passports {
		if passport.checkValidity() {
			validPassports = append(validPassports, passport)
		}
	}
	return validPassports
}
