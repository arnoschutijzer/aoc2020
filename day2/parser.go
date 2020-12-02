package daytwo

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func ReadFile(path string) []Password {
	data, _ := ioutil.ReadFile(path)
	textinput := string(data)
	list := strings.Split(textinput, "\n")

	var passwords []Password

	for _, passwordEntry := range list {
		var password Password
		re := regexp.MustCompile(`([0-9])*-(0-9)*\w+`)
		characterCount := re.FindString(passwordEntry)

		minOccurence, _ := strconv.Atoi((string(characterCount[0])))
		maxOccurence, _ := strconv.Atoi((string(characterCount[2])))
		password.minOccurence = minOccurence
		password.maxOccurence = maxOccurence

		re = regexp.MustCompile(`([A-z]):+`)
		password.letter = strings.ReplaceAll(re.FindString(passwordEntry), ":", "")

		seperatorIndex := strings.Index(passwordEntry, ": ")
		password.password = string(passwordEntry[seperatorIndex+2:])

		passwords = append(passwords, password)
	}

	return passwords
}
