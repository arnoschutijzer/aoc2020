package dayone

import (
	stat "gonum.org/v1/gonum/stat/combin"
)

// there's probably some helpers somewhere...
func fixReport(expenses []int, size int) int {
	combinations := stat.Combinations(len(expenses), size)
	var foundExpense []int
	for _, combination := range combinations {
		var tuple []int
		var sumTotal int
		for index := range combination {
			tuple = append(tuple, expenses[combination[index]])
			sumTotal += expenses[combination[index]]
		}
		if sumTotal == 2020 {
			foundExpense = tuple
			break
		}
	}

	result := foundExpense[0]
	for _, value := range foundExpense[1:] {
		result = result * value
	}

	return result
}
