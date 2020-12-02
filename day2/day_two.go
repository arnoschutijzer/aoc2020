package daytwo

import "strings"

type Rule interface {
	isValid(password Password) bool
}

type TobogganRule struct{}

func (rule TobogganRule) isValid(password Password) bool {
	firstCharacter := string(password.password[password.minOccurence-1])
	doesFirstCharacterMatch := firstCharacter == password.letter

	if password.maxOccurence > len(password.password) {
		return doesFirstCharacterMatch
	}

	secondCharacter := string(password.password[password.maxOccurence-1])
	doesSecondCharacterMatch := secondCharacter == password.letter

	return (doesFirstCharacterMatch || doesSecondCharacterMatch) && (firstCharacter != secondCharacter)
}

type SledRule struct{}

func (rule SledRule) isValid(password Password) bool {
	letterCount := strings.Count(password.password, password.letter)
	return letterCount >= password.minOccurence && letterCount <= password.maxOccurence
}

type Password struct {
	letter       string
	minOccurence int
	maxOccurence int

	password string
}

func CountValidPasswords(passwords []Password, rule Rule) int {
	validPasswords := 0

	for _, password := range passwords {
		if rule.isValid(password) {
			validPasswords += 1
		}
	}

	return validPasswords
}
