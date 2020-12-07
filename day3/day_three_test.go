package daythree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var defaultSlope = Slope{
	right: 3,
	down:  1,
}

func TestTripConstructorInitializesCorrectly(t *testing.T) {
	trip := NewTrip([][]string{}, defaultSlope)

	assert.Equal(t, 0, trip.encounteredTrees)
	assert.Equal(t, 0, trip.currentXPosition)
	assert.Equal(t, 0, trip.currentYPosition)
}

func TestTripCorrectlyMovesOn(t *testing.T) {
	trip := NewTrip([][]string{}, defaultSlope)

	assert.Equal(t, 0, trip.currentXPosition)
	assert.Equal(t, 0, trip.currentYPosition)

	trip.moveOn()

	assert.Equal(t, 3, trip.currentXPosition)
	assert.Equal(t, 1, trip.currentYPosition)
}

func TestCounts2TreesFromFixture(t *testing.T) {
	trip := NewTrip(ParseTopology("./fixtures/topology_input.txt"), defaultSlope)
	assert.Equal(t, 2, trip.WalkTheWalkAndCountTheTrees())
}

func TestCountsTreesFromFixture(t *testing.T) {
	trip := NewTrip(ParseTopology("./fixtures/topology_personal.txt"), defaultSlope)
	assert.Equal(t, 228, trip.WalkTheWalkAndCountTheTrees())
}

func TestCountsTreesFromFixturePartTwo(t *testing.T) {
	slopes := []Slope{
		{
			right: 1,
			down:  1,
		},
		defaultSlope,
		{
			right: 5,
			down:  1,
		},
		{
			right: 7,
			down:  1,
		},
		{
			right: 1,
			down:  2,
		},
	}

	var treesForSlopes []int
	topology := ParseTopology("./fixtures/topology_personal.txt")

	for _, slope := range slopes {
		trip := NewTrip(topology, slope)
		treesForSlopes = append(treesForSlopes, trip.WalkTheWalkAndCountTheTrees())
	}

	multipliedTrees := 1
	for _, trees := range treesForSlopes {
		multipliedTrees = multipliedTrees * trees
	}

	assert.Equal(t, 6818112000, multipliedTrees)
}
