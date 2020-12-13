package dayfour

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ContainsAllFields(passport string) bool {
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
		if !strings.Contains(passport, field) {
			return false
		}
	}

	return true
}

func CompliesWithRules(passport string) bool {
	rules := []bool{
		hasValidID(passport),
		hasValidBirthYear(passport),
		hasValidExpirationYear(passport),
		hasValidIssueYear(passport),
		hasValidEyeColor(passport),
		hasValidHeight(passport),
		hasValidHairColor(passport),
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
	parsedValue, _ := strconv.Atoi(value)

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
	height, _ := strconv.Atoi(value)

	if unit == "cm" {
		return height >= 150 && height <= 193
	}

	return height >= 59 && height <= 76
}

func OnlyPassportsThatContainAllFields(passports []string) []string {
	var validPassports []string
	for _, passport := range passports {
		if ContainsAllFields(passport) {
			validPassports = append(validPassports, passport)
		}
	}
	return validPassports
}

func OnlyPassportsThatComplyWithRules(passports []string) []string {
	var validPassports []string
	for _, passport := range passports {
		if CompliesWithRules(passport) {
			validPassports = append(validPassports, passport)
		}
	}
	return validPassports
}
