package board

import (
	"errors"
	"fmt"
	"io"

	"github.com/Crystalix007/sudoku.git/lib/sets"
)

var ErrInvalidValue = errors.New("lib/board: invalid value")

type Board interface {
	Clone() Board
	Column(x uint8) [9]Cell
	Empty() []Coordinate
	Equal(other Board) bool
	Filled() []Coordinate
	Get(x, y uint8) Cell
	Options(x, y uint8) []Cell
	Output(w io.StringWriter)
	Row(y uint8) [9]Cell
	Set(x, y uint8, val Cell) error
	Supercell(x, y uint8) [3][3]Cell
}

type board struct {
	// Row-first representation of the board grid.
	grid [9][9]Cell
}

var _ Board = &board{} // Ensure *board conforms to Board interface.

// New creates a new empty board.
func New() Board {
	return &board{}
}

// FromGrid creates a board from a given grid.
func FromGrid(grid [9][9]Cell) Board {
	return &board{
		grid: grid,
	}
}

// Clone retrieves a copy of the Board.
func (b *board) Clone() Board {
	clone := &board{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			clone.grid[i][j] = b.grid[i][j]
		}
	}

	return clone
}

// Column returns a column of cells in a board.
func (b *board) Column(x uint8) [9]Cell {
	column := [9]Cell{}

	for i := 0; i < 9; i++ {
		column[i] = b.grid[i][x]
	}

	return column
}

// Get retrives a cell's value.
func (b *board) Get(x uint8, y uint8) uint8 {
	return b.grid[x][y]
}

// Set sets a cell's value, or errors if the new value is invalid.
func (b *board) Set(x uint8, y uint8, val uint8) error {
	prevValue := b.grid[y][x]
	b.grid[y][x] = val

	if b.validateColumn(x) && b.validateRow(y) && b.validateSupercell(x/3, y/3) {
		return nil
	}

	b.grid[y][x] = prevValue
	return ErrInvalidValue
}

// Empty returns all the empty cells in a board.
func (b *board) Empty() []Coordinate {
	empty := []Coordinate{}

	for i := uint8(0); i < 9; i++ {
		for j := uint8(0); j < 9; j++ {
			if isEmpty(b.grid[i][j]) {
				empty = append(empty, Coordinate{
					X: j,
					Y: i,
				})
			}
		}
	}

	return empty
}

// Equal computes whether this board equals another Board.
func (b *board) Equal(other Board) bool {
	// We're confined to the Board interface, so check elementwise equality.
	for i := uint8(0); i < 9; i++ {
		for j := uint8(0); j < 9; j++ {
			if b.Get(j, i) != other.Get(j, i) {
				return false
			}
		}
	}

	return true
}

// Filled returns all the non-empty cells in a board.
func (b *board) Filled() []Coordinate {
	filled := []Coordinate{}

	for i := uint8(0); i < 9; i++ {
		for j := uint8(0); j < 9; j++ {
			if !isEmpty(b.grid[i][j]) {
				filled = append(filled, Coordinate{
					X: j,
					Y: i,
				})
			}
		}
	}

	return filled
}

// Row returns a row of cells in a board.
func (b *board) Row(y uint8) [9]Cell {
	row := [9]Cell{}

	for i := 0; i < 9; i++ {
		row[i] = b.grid[y][i]
	}

	return row
}

// Supercell returns a 3x3 "supercell" of cells.
func (b *board) Supercell(x, y uint8) [3][3]Cell {
	supercell := [3][3]Cell{}

	cellBaseX := x * 3
	cellBaseY := y * 3

	for i := uint8(0); i < 3; i++ {
		for j := uint8(0); j < 3; j++ {
			supercell[j][i] = b.grid[cellBaseY+i][cellBaseX+j]
		}
	}

	return supercell
}

// Options returns the available options at a given cell coordinate.
func (b *board) Options(x, y uint8) []Cell {
	rowOptions := b.optionsRow(y)
	colOptions := b.optionsColumn(x)
	supercellOptions := b.optionsSupercell(x/3, y/3)

	intersectedOptions := sets.Intersect(rowOptions[:], sets.Intersect(colOptions[:], supercellOptions[:]))
	optionIndices := sets.TrueIndices[Cell](intersectedOptions)

	for i := range optionIndices {
		optionIndices[i]++
	}

	return optionIndices
}

// Output writes out a representation of the board to an io.Writer.
func (b *board) Output(w io.StringWriter) {
	b.outputTopRow(w)

	for i := uint8(0); i < 9; i++ {
		if i != 0 {
			b.outputRowSeperator(i, w)
		}

		b.outputRow(i, w)
	}

	b.outputBottomRow(w)
}

func (b *board) validateColumn(col uint8) bool {
	counts := [9]uint8{}

	for j := uint8(0); j < 9; j++ {
		val := b.grid[j][col]

		if isEmpty(val) {
			continue
		}

		counts[val-1]++
	}

	for _, count := range counts {
		if count > 1 {
			return false
		}
	}

	return true
}

func (b *board) validateRow(row uint8) bool {
	counts := [9]uint8{}

	for i := uint8(0); i < 9; i++ {
		val := b.grid[row][i]

		if isEmpty(val) {
			continue
		}

		counts[val-1]++
	}

	for _, count := range counts {
		if count > 1 {
			return false
		}
	}

	return true
}

// validateSupercell validates a 3x3 group of cells. This is indexed in the same format
func (b *board) validateSupercell(superX, superY uint8) bool {
	counts := [9]uint8{}

	superCell := b.Supercell(superX, superY)

	for i := uint8(0); i < 3; i++ {
		for j := uint8(0); j < 3; j++ {
			val := superCell[j][i]

			if isEmpty(val) {
				continue
			}

			counts[val-1]++
		}
	}

	for _, count := range counts {
		if count > 1 {
			return false
		}
	}

	return true
}

func (b *board) optionsColumn(x uint8) [9]bool {
	return UnusedValues(b.Column(x))
}

func (b *board) optionsRow(y uint8) [9]bool {
	return UnusedValues(b.Row(y))
}

func (b *board) optionsSupercell(superX, superY uint8) [9]bool {
	return UnusedValues(FlattenSupercell(b.Supercell(superX, superY)))
}

func (b *board) outputTopRow(w io.StringWriter) {
	topRow := "-------------------\n"
	w.WriteString(topRow)
}

func (b *board) outputRowSeperator(i uint8, w io.StringWriter) {
	rowSeperator := "-------------------\n"
	if i%3 != 0 {
		rowSeperator = "|     |     |     |\n"
	}
	w.WriteString(rowSeperator)
}

func (b *board) outputRow(i uint8, w io.StringWriter) {
	for j := uint8(0); j < 9; j++ {
		if j%3 == 0 {
			w.WriteString("|")
		} else {
			w.WriteString(" ")
		}

		if cell := b.Get(i, j); !isEmpty(cell) {
			w.WriteString(fmt.Sprintf("%d", cell))
		} else {
			w.WriteString(" ")
		}
	}
	w.WriteString("|\n")
}

func (b *board) outputBottomRow(w io.StringWriter) {
	bottomRow := "-------------------\n"
	w.WriteString(bottomRow)
}
