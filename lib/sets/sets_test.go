package sets_test

import (
	"testing"

	"github.com/Crystalix007/sudoku.git/lib/sets"
	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	assert.Equal(t, []uint{}, sets.Intersect([]bool{}, []bool{}))
	assert.Equal(t, []uint{}, sets.Intersect([]bool{true}, []bool{}))
	assert.Equal(t, []uint{}, sets.Intersect([]bool{false}, []bool{false}))
	assert.Equal(t, []uint{0}, sets.Intersect([]bool{true}, []bool{true}))
	assert.Equal(t, []uint{0}, sets.Intersect([]bool{true}, []bool{true, true}))
}
