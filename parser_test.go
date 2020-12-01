package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	assert.Equal(t, []string{"AAA", "BBB", "CCC"}, ReadFile("./fixtures/parser_test_1.txt"))
}
