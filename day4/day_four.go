package dayfour

import (
	"strings"
)

type Passport struct {
	fields string

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
		if !strings.Contains(passport.fields, field) {
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
