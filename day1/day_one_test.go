package dayone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatesMultiplicationOfPairWhichSumEquals2020(t *testing.T) {
	expenses := ReadFile("./fixtures/dayone_1.txt")
	assert.Equal(t, 514579, fixReport(expenses, 2))
}

func TestCalculatesMultiplicationOfPairWhichSumEquals2020_2(t *testing.T) {
	expenses := ReadFile("./fixtures/dayone_personal.txt")
	assert.Equal(t, 805731, fixReport(expenses, 2))
}

func TestCalculatesMultiplicationOfTupleWithSize3WhichSumEquals2020(t *testing.T) {
	expenses := ReadFile("./fixtures/dayone_1.txt")
	assert.Equal(t, 241861950, fixReport(expenses, 3))
}

func TestCalculatesMultiplicationOfTupleWithSize3WhichSumEquals2020_2(t *testing.T) {
	expenses := ReadFile("./fixtures/dayone_personal.txt")
	assert.Equal(t, 192684960, fixReport(expenses, 3))
}

func TestCalculatesMultiplicationOfTupleWithSize3WhichSumEquals2020WithCombinations(t *testing.T) {
	expenses := ReadFile("./fixtures/dayone_1.txt")
	assert.Equal(t, 241861950, fixReport(expenses, 3))
}

func TestCalculatesMultiplicationOfTupleWithSize3WhichSumEquals2020WithCombinations_2(t *testing.T) {
	expenses := ReadFile("./fixtures/dayone_personal.txt")
	assert.Equal(t, 192684960, fixReport(expenses, 3))
}
