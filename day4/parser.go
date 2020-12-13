package dayfour

import (
	"io/ioutil"
	"strings"
)

func ParsePassports(path string) []string {
	bytes, _ := ioutil.ReadFile(path)
	contents := string(bytes)

	return strings.Split(contents, "\n\n")
}
