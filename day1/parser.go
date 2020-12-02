package dayone

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(path string) []int {
	data, _ := ioutil.ReadFile(path)
	textinput := string(data)
	list := strings.Split(textinput, "\n")

	var expenses []int
	for _, expense := range list {
		parsedExpense, _ := strconv.Atoi(expense)
		expenses = append(expenses, parsedExpense)
	}

	return expenses
}
