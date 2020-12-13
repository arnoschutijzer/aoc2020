package dayfour

import (
	"io/ioutil"
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
		})
	}

	return passports
}
