package board

// Cell represents the state in a single square in the Sudoku board.
type Cell = uint8

// Coordinate represents an x-y coordinate on the board.
type Coordinate struct {
	X, Y uint8
}

// Determines whether the given cell value represents the "empty" value.
func isEmpty(val Cell) bool {
	return val == 0
}

// FlattenSupercell takes a 2D supercell grid, and flattens it into a linear array.
// (0,0), (1,0), (2,0), (0, 1), ...
func FlattenSupercell(cs [3][3]Cell) [9]Cell {
	flattened := [9]Cell{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			flattened[i*3+j] = cs[i][j]
		}
	}

	return flattened
}

// UsedValues returns a set of values which have been used.
func UsedValues(cs [9]Cell) [9]bool {
	usedValues := [9]bool{}

	for _, cs := range cs {
		if !isEmpty(cs) {
			usedValues[cs-1] = true
		}
	}

	return usedValues
}

// UnusedValues returns a set of values which haven't been used.
func UnusedValues(cs [9]Cell) [9]bool {
	usedValues := UsedValues(cs)

	// Invert all usedValues
	for i := range usedValues {
		usedValues[i] = !usedValues[i]
	}

	return usedValues
}
