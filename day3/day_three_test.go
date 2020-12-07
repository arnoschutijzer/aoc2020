package daythree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounts1TreeWith1Iteration(t *testing.T) {
	topology := ParseTopology("./fixtures/topology_input.txt")
	assert.Equal(t, 1, CountEncounteredTrees(topology))
}
