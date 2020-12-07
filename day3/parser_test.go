package daythree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTopology(t *testing.T) {
	expected := [][]string{
		{".", ".", "#", "."},
		{"#", ".", "#", "#"},
		{"#", "#", "#", "."},
	}
	assert.Equal(t, expected, ParseTopology("./fixtures/topology_input.txt"))
}
