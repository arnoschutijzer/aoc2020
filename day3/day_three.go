package daythree

type Trip struct {
	topology         [][]string
	encounteredTrees int
	currentXPosition int
	currentYPosition int
}

func NewTrip(topology [][]string) *Trip {
	return &Trip{
		topology:         topology,
		currentXPosition: 0,
		currentYPosition: 0,
		encounteredTrees: 0,
	}
}

func (trip *Trip) moveOn() {
	trip.currentXPosition += 3
	trip.currentYPosition += 1
}

func (trip *Trip) checkCurrentSpotForTrees() {
	rowLength := trip.topology[trip.currentYPosition]

	if trip.topology[trip.currentYPosition][trip.currentXPosition%len(rowLength)] == "#" {
		trip.encounteredTrees += 1
	}
}

func CountEncounteredTrees(topology [][]string) int {
	trip := NewTrip(topology)

	row := 0
	for row < len(trip.topology)-1 {
		trip.moveOn()
		trip.checkCurrentSpotForTrees()
		row++
	}

	return trip.encounteredTrees
}
