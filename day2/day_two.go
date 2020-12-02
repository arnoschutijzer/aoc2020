package daytwo

import "strings"

type Password struct {
	letter       string
	minOccurence int
	maxOccurence int

	password string
}

func CalculateValidPasswords(passwords []Password) int {
	validPasswords := 0

	for _, password := range passwords {
		letterCount := strings.Count(password.password, password.letter)

		if letterCount >= password.minOccurence && letterCount <= password.maxOccurence {
			validPasswords += 1
		}
	}

	return validPasswords
}
