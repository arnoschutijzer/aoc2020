package daytwo

import "strings"

type RuleSet string

const (
	PartOne RuleSet = "partOne"
	PartTwo RuleSet = "partTwo"
)

type Password struct {
	letter       string
	minOccurence int
	maxOccurence int

	password string
}

func (password Password) isValidPartOne() bool {
	letterCount := strings.Count(password.password, password.letter)
	return letterCount >= password.minOccurence && letterCount <= password.maxOccurence
}

func (password Password) isValidPartTwo() bool {
	firstCharacter := string(password.password[password.minOccurence-1])
	doesFirstCharacterMatch := firstCharacter == password.letter

	if password.maxOccurence > len(password.password) {
		return doesFirstCharacterMatch
	}

	secondCharacter := string(password.password[password.maxOccurence-1])
	doesSecondCharacterMatch := secondCharacter == password.letter

	if firstCharacter == secondCharacter {
		return false
	}

	return doesFirstCharacterMatch || doesSecondCharacterMatch
}

func CalculateValidPasswords(passwords []Password, ruleset RuleSet) int {
	validPasswords := 0

	for _, password := range passwords {
		switch ruleset {
		case PartOne:
			if password.isValidPartOne() {
				validPasswords += 1
			}
			break

		case PartTwo:
			if password.isValidPartTwo() {
				validPasswords += 1
			}
			break
		}
	}

	return validPasswords
}
