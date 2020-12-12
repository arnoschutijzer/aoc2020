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
			fields: passportString,
			pid:    GetIDFromPassword(passportString),
		})
	}

	return passports
}

func GetIDFromPassword(stringifiedPassport string) string {
	re := regexp.MustCompile(`(?m)pid:(?P<value>[0-9]{9})($| )`)
	matches := re.FindStringSubmatch(stringifiedPassport)

	indexOfSub := re.SubexpIndex("value")
	if indexOfSub >= len(matches) {
		return ""
	}

	return matches[re.SubexpIndex("value")]
}
