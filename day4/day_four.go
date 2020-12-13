package dayfour

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	stringified string
}

func (passport *Passport) containsAllFields() bool {
	var validFields = []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, field := range validFields {
		if !strings.Contains(passport.stringified, field) {
			return false
		}
	}

	return true
}

func (passport *Passport) compliesWithRules() bool {
	rules := []bool{
		hasValidID(passport.stringified),
		hasValidBirthYear(passport.stringified),
		hasValidExpirationYear(passport.stringified),
		hasValidIssueYear(passport.stringified),
		hasValidEyeColor(passport.stringified),
		hasValidHeight(passport.stringified),
		hasValidHairColor(passport.stringified),
	}

	for _, rule := range rules {
		if !rule {
			return false
		}
	}

	return true
}

func matchesRegex(aString string, expression string) bool {
	re := regexp.MustCompile(expression)
	match := re.FindStringSubmatch(aString)

	return len(match) > 0
}

func hasValidID(stringifiedPassport string) bool {
	return matchesRegex(stringifiedPassport, `(?m)pid:(?P<value>[0-9]{9})($| )`)
}

func hasValidEyeColor(stringifiedPassport string) bool {
	return matchesRegex(stringifiedPassport, `ecl:(amb|blu|brn|gry|grn|hzl|oth)`)
}

func hasValidHairColor(stringifiedPassport string) bool {
	return matchesRegex(stringifiedPassport, `hcl:#[a-f0-9]{6}`)
}

func isSubmatchBetween(aString string, keyword string, expression string, max int, min int) bool {
	re := regexp.MustCompile(expression)
	match := re.FindStringSubmatch(aString)

	if len(match) == 0 {
		return false
	}

	value := match[re.SubexpIndex(keyword)]
	parsedValue, error := strconv.Atoi(value)

	if error != nil {
		return false
	}

	return parsedValue <= max && parsedValue >= min
}

func hasValidIssueYear(stringifiedPassport string) bool {
	keyword := "value"
	expression := fmt.Sprintf("(?m)iyr:(?P<%s>[0-9]{4})($| )", keyword)
	return isSubmatchBetween(stringifiedPassport, keyword, expression, 2020, 2010)
}

func hasValidExpirationYear(stringifiedPassport string) bool {
	keyword := "value"
	expression := fmt.Sprintf("(?m)eyr:(?P<%s>[0-9]{4})($| )", keyword)
	return isSubmatchBetween(stringifiedPassport, keyword, expression, 2030, 2020)
}

func hasValidBirthYear(stringifiedPassport string) bool {
	keyword := "value"
	expression := fmt.Sprintf("(?m)byr:(?P<%s>[0-9]{4})($| )", keyword)
	return isSubmatchBetween(stringifiedPassport, keyword, expression, 2002, 1920)
}

func hasValidHeight(stringifiedPassport string) bool {
	re := regexp.MustCompile(`(?m)hgt:(?P<value>[0-9]+)(?P<unit>(cm|in))($| )`)
	match := re.FindStringSubmatch(stringifiedPassport)

	if len(match) == 0 {
		return false
	}

	value := match[re.SubexpIndex("value")]
	unit := match[re.SubexpIndex("unit")]
	height, error := strconv.Atoi(value)

	if error != nil {
		return false
	}

	if unit == "cm" {
		return height >= 150 && height <= 193
	}

	return height >= 59 && height <= 76
}

func OnlyPassportsThatContainAllFields(passports []Passport) []Passport {
	var validPassports []Passport
	for _, passport := range passports {
		if passport.containsAllFields() {
			validPassports = append(validPassports, passport)
		}
	}
	return validPassports
}

func OnlyPassportsThatComplyWithRules(passports []Passport) []Passport {
	var validPassports []Passport
	for _, passport := range passports {
		if passport.compliesWithRules() {
			validPassports = append(validPassports, passport)
		}
	}
	return validPassports
}
