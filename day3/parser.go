package daythree

import (
	"io/ioutil"
	"strings"
)

func ParseTopology(path string) [][]string {
	bytes, _ := ioutil.ReadFile(path)
	input := string(bytes)

	rows := strings.Split(input, "\n")

	var topology [][]string

	for _, value := range rows {
		topology = append(topology, strings.Split(value, ""))
	}

	return topology
}
