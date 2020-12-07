package daythree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTripConstructorInitializesCorrectly(t *testing.T) {
	trip := NewTrip([][]string{})

	assert.Equal(t, 0, trip.encounteredTrees)
	assert.Equal(t, 0, trip.currentXPosition)
	assert.Equal(t, 0, trip.currentYPosition)
}

func TestTripCorrectlyMovesOn(t *testing.T) {
	trip := NewTrip([][]string{})

	assert.Equal(t, 0, trip.currentXPosition)
	assert.Equal(t, 0, trip.currentYPosition)

	trip.moveOn()

	assert.Equal(t, 3, trip.currentXPosition)
	assert.Equal(t, 1, trip.currentYPosition)
}

func TestCounts2TreesFromFixture(t *testing.T) {
	topology := ParseTopology("./fixtures/topology_input.txt")
	assert.Equal(t, 2, CountEncounteredTrees(topology))
}

func TestCountsTreesFromFixture(t *testing.T) {
	topology := ParseTopology("./fixtures/topology_personal.txt")
	assert.Equal(t, 228, CountEncounteredTrees(topology))
}
