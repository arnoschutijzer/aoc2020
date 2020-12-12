package dayfour

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func ParsePassports(path string) []Passport {
	bytes, _ := ioutil.ReadFile(path)
	contents := string(bytes)

	passportStrings := strings.Split(contents, "\n\n")

	var passports []Passport
	for _, passportString := range passportStrings {
		passports = append(passports, Passport{
			stringified: passportString,
			fields:      GetFieldsFrom(passportString),
		})
	}

	return passports
}

func GetFieldsFrom(stringifiedPassport string) map[string]string {
	fields := make(map[string]string)
	// this can be improved to include all characters excluding whitespace...
	re := regexp.MustCompile(`(?m)(?P<name>\w+):(?P<value>#?\w+)($| )`)
	matches := re.FindAllStringSubmatch(stringifiedPassport, -1)

	for _, match := range matches {
		key := match[re.SubexpIndex("name")]
		value := match[re.SubexpIndex("value")]
		fields[key] = value
	}

	return fields
}
