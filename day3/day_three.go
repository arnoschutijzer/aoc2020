package daythree

type Trip struct {
	slope            Slope
	topology         [][]string
	encounteredTrees int
	currentXPosition int
	currentYPosition int
}

type Slope struct {
	right int
	down  int
}

func NewTrip(topology [][]string, slope Slope) *Trip {
	return &Trip{
		slope:            slope,
		topology:         topology,
		currentXPosition: 0,
		currentYPosition: 0,
		encounteredTrees: 0,
	}
}

func (trip *Trip) moveOn() {
	trip.currentXPosition += trip.slope.right
	trip.currentYPosition += trip.slope.down
}

func (trip *Trip) checkCurrentSpotForTrees() {
	rowLength := trip.topology[trip.currentYPosition]

	if trip.topology[trip.currentYPosition][trip.currentXPosition%len(rowLength)] == "#" {
		trip.encounteredTrees += 1
	}
}

func (trip *Trip) WalkTheWalkAndCountTheTrees() int {
	row := 0
	for row < len(trip.topology)-1 {
		trip.moveOn()
		trip.checkCurrentSpotForTrees()
		row += trip.slope.down
	}

	return trip.encounteredTrees
}
