package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	assert.Equal(t, []int{123, 456, 789}, ReadFile("./fixtures/parser_test_1.txt"))
}
