package parser

import (
	"io/ioutil"
	"strings"
)

func ReadFile(path string) []string {
	data, _ := ioutil.ReadFile(path)
	textinput := string(data)
	list := strings.Split(textinput, "\n")

	return list
}
