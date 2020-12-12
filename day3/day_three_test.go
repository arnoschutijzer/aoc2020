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
	trip := NewTrip([][]string{}, &defaultSlope)

	assert.Equal(t, 0, trip.encounteredTrees)
	assert.Equal(t, 0, trip.currentXPosition)
	assert.Equal(t, 0, trip.currentYPosition)
}

func TestTripCorrectlyMovesOn(t *testing.T) {
	trip := NewTrip([][]string{}, &defaultSlope)

	assert.Equal(t, 0, trip.currentXPosition)
	assert.Equal(t, 0, trip.currentYPosition)

	trip.moveOn()

	assert.Equal(t, 3, trip.currentXPosition)
	assert.Equal(t, 1, trip.currentYPosition)
}

func TestCounts2TreesFromFixture(t *testing.T) {
	trip := NewTrip(ParseTopology("./fixtures/topology_input.txt"), &defaultSlope)
	assert.Equal(t, 2, trip.walkTheWalkAndCountTheTrees())
}

func TestCountsTreesFromFixture(t *testing.T) {
	trip := NewTrip(ParseTopology("./fixtures/topology_personal.txt"), &defaultSlope)
	assert.Equal(t, 228, trip.walkTheWalkAndCountTheTrees())
}

func TestCountsTreesFromFixturePartTwo(t *testing.T) {
	slopes := []*Slope{
		{
			right: 1,
			down:  1,
		},
		&defaultSlope,
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

	topology := ParseTopology("./fixtures/topology_personal.txt")

	assert.Equal(t, 6818112000, CountTreesForMultipleSlopes(topology, slopes))
}
