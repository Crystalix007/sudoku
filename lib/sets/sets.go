package sets

import (
	"github.com/Crystalix007/sudoku.git/lib/maths"
	"golang.org/x/exp/constraints"
)

func Intersect(a, b []bool) []bool {
	intersection := []bool{}

	if len(a) != len(b) {
		return intersection
	}

	for i := 0; i < maths.Max(len(a), len(b)); i++ {
		intersection = append(intersection, a[i] && b[i])
	}

	return intersection
}

func TrueIndices[T constraints.Unsigned](bs []bool) []T {
	trueIndices := []T{}

	for i, b := range bs {
		if b {
			trueIndices = append(trueIndices, T(i))
		}
	}

	return trueIndices
}
