package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	lines := []string{"Donkey Kong IX", "Donkey Kong I", "David III", "August II", "David I"}
	Sort(lines)
	assert.Equal(t, lines[0], "August II")
	assert.Equal(t, lines[1], "David I")
	assert.Equal(t, lines[2], "David III")
	assert.Equal(t, lines[3], "Donkey Kong I")
	assert.Equal(t, lines[4], "Donkey Kong IX")
}
