package main

import (
	"fmt"
	"os"

	"github.com/Crystalix007/sudoku.git/lib/board"
	"github.com/Crystalix007/sudoku.git/lib/solver"
)

func main() {
	board := board.FromGrid([9][9]board.Cell{
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

	solutions := solver.Solve(board)

	if len(solutions) == 0 {
		fmt.Printf("Failed to find solution\n")
		os.Exit(1)
	}

	fmt.Printf("Found %d solution(s):\n\n", len(solutions))

	solutions[0].Output(os.Stdout)
}
