package dayone

// yikes
// use combinations :(
func fixReportPartOne(expenses []int) int {
	var possibleCombinations [][]int
	for _, firstExpense := range expenses {
		var remainingExpenses = expenses[1:]
		for _, secondExpense := range remainingExpenses {
			possibleCombinations = append(possibleCombinations, []int{firstExpense, secondExpense})
		}
	}

	var solutionTuple []int
	for _, combination := range possibleCombinations {
		if combination[0]+combination[1] == 2020 {
			solutionTuple = combination
		}
	}

	return solutionTuple[0] * solutionTuple[1]
}

func fixReportPartTwo(expenses []int) int {
	var possibleCombinations [][]int
	for index, firstExpense := range expenses {
		for secondIndex, secondExpense := range expenses[index:] {
			for _, thirdExpense := range expenses[secondIndex:] {
				possibleCombinations = append(possibleCombinations, []int{firstExpense, secondExpense, thirdExpense})
			}
		}
	}

	var solutionTuple []int
	for _, combination := range possibleCombinations {
		if combination[0]+combination[1]+combination[2] == 2020 {
			solutionTuple = combination
		}
	}

	return solutionTuple[0] * solutionTuple[1] * solutionTuple[2]
}
