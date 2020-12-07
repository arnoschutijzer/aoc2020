package daythree

func CountEncounteredTrees(topology [][]string) int {
	trees := 0

	if topology[1][3] == "#" {
		trees += 1
	}

	return trees
}
