package dayone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayOne(t *testing.T) {
	assert.Equal(t, 514579, calc("dayone_1.txt"))
}