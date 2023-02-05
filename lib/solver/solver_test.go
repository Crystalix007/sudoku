package solver_test

import (
	"math/rand"
	"os"
	"testing"

	"github.com/Crystalix007/sudoku.git/lib/arrays"
	"github.com/Crystalix007/sudoku.git/lib/board"
	"github.com/Crystalix007/sudoku.git/lib/solver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolverStaticExample(t *testing.T) {
	question := board.FromGrid([9][9]board.Cell{
		{0, 7, 0, 0, 2, 0, 0, 4, 6},
		{0, 6, 0, 0, 0, 0, 8, 9, 0},
		{2, 0, 0, 8, 0, 0, 0, 1, 5},
		{0, 8, 4, 0, 9, 7, 0, 0, 0},
		{7, 1, 0, 0, 0, 0, 0, 5, 9},
		{0, 0, 0, 1, 3, 0, 4, 8, 0},
		{6, 9, 7, 0, 0, 2, 0, 0, 8},
		{0, 5, 8, 0, 0, 0, 0, 6, 0},
		{4, 3, 0, 0, 8, 0, 0, 7, 0},
	})

	answer := board.FromGrid([9][9]board.Cell{
		{8, 7, 5, 9, 2, 1, 3, 4, 6},
		{3, 6, 1, 7, 5, 4, 8, 9, 2},
		{2, 4, 9, 8, 6, 3, 7, 1, 5},
		{5, 8, 4, 6, 9, 7, 1, 2, 3},
		{7, 1, 3, 2, 4, 8, 6, 5, 9},
		{9, 2, 6, 1, 3, 5, 4, 8, 7},
		{6, 9, 7, 4, 1, 2, 5, 3, 8},
		{1, 5, 8, 3, 7, 9, 2, 6, 4},
		{4, 3, 2, 5, 8, 6, 9, 7, 1},
	})

	solutions := solver.Solve(question)
	assert.Equal(t, 1, len(solutions), "solver finds solutions")
	assert.True(t, answer.Equal(solutions[0]))
}

func TestSolverDynamicExamples(t *testing.T) {
	// To test the solver with dynamic samples, we will test determinism.
	// Essentially, if we can get back to the same result from the same input,
	// then the algorithm is deterministic.
	//
	// To test this, we solve one time. Then erase some numbers to create
	// blanks. Then check that the result reached is the same as the original
	// solve.

	initBoard := board.New()
	empty := initBoard.Empty()

	const initialValueCount = 36

	for i := 0; i < initialValueCount; i++ {
		nextIndex := arrays.RandomIndex(empty)

		var nextCell board.Coordinate
		nextCell, empty = empty[nextIndex], arrays.RemoveIndex(uint(nextIndex), empty)

		options := initBoard.Options(nextCell.X, nextCell.Y)
		nextOptionIndex := arrays.RandomIndex(options)
		nextOption := options[nextOptionIndex]

		require.NoError(t, initBoard.Set(nextCell.X, nextCell.Y, nextOption), "no error setting option on cell")
	}

	// We might not have generated a valid board. I.e. the board may not have a
	// solution. Until we have, go and tweak values used.
	var completion1 board.Board
	for {
		completions := solver.Solve(initBoard)
		if len(completions) != 0 {
			completion1 = completions[0]
			break
		}

		filled := initBoard.Filled()
		randFilledIndex := arrays.RandomIndex(filled)
		randFilled := filled[randFilledIndex]

		initBoard.Set(randFilled.X, randFilled.Y, 0)
	}

	initBoard.Output(os.Stderr)

	question := completion1.Clone()

	const removedValueCount = 1

	for i := 0; i < removedValueCount; i++ {
		nextCell := board.Coordinate{
			X: uint8(rand.Int31n(9)),
			Y: uint8(rand.Int31n(9)),
		}

		require.NoError(t, initBoard.Set(nextCell.X, nextCell.Y, 0), "no error blanking cell")
	}

	completions := solver.Solve(question)
	require.NotEmpty(t, completions, "has completions")
	assert.Contains(t, completions, completion1, "finds original source solution")
}
