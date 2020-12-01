package dayone

import (
	"testing"

	parser "arnoschutijzer.io/aoc2020/v2"
	"github.com/stretchr/testify/assert"
)

func TestCalculatesMultiplicationOfPairWhichSumEquals2020(t *testing.T) {
	expenses := parser.ReadFile("./fixtures/dayone_1.txt")
	assert.Equal(t, 514579, fixReportPartOne(expenses))
}

func TestCalculatesMultiplicationOfPairWhichSumEquals2020_2(t *testing.T) {
	expenses := parser.ReadFile("./fixtures/dayone_personal.txt")
	assert.Equal(t, 805731, fixReportPartOne(expenses))
}

func TestCalculatesMultiplicationOfTupleWithSize3WhichSumEquals2020(t *testing.T) {
	expenses := parser.ReadFile("./fixtures/dayone_1.txt")
	assert.Equal(t, 241861950, fixReportPartTwo(expenses))
}

func TestCalculatesMultiplicationOfTupleWithSize3WhichSumEquals2020_2(t *testing.T) {
	expenses := parser.ReadFile("./fixtures/dayone_personal.txt")
	assert.Equal(t, 192684960, fixReportPartTwo(expenses))
}
}