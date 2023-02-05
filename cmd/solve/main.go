package main

import (
	"fmt"
	"os"

	"github.com/Crystalix007/sudoku.git/lib/board"
	"github.com/Crystalix007/sudoku.git/lib/solver"
)

func main() {
	board1 := board.FromGrid([9][9]board.Cell{
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

	// board2 := board.FromGrid([9][9]board.Cell{
	// 	{0, 7, 0, 0, 2, 0, 0, 4, 6},
	// 	{0, 6, 1, 7, 5, 4, 8, 9, 2},
	// 	{2, 4, 9, 8, 6, 3, 7, 1, 5},
	// 	{5, 8, 4, 6, 9, 7, 1, 2, 3},
	// 	{7, 1, 3, 2, 4, 8, 6, 5, 9},
	// 	{9, 2, 6, 1, 3, 5, 4, 8, 7},
	// 	{6, 9, 7, 4, 1, 2, 5, 3, 8},
	// 	{1, 5, 8, 3, 7, 9, 2, 6, 4},
	// 	{4, 3, 2, 5, 8, 6, 9, 7, 1},
	// })

	solutions := solver.Solve(board1)

	if len(solutions) == 0 {
		fmt.Printf("Failed to find solution\n")
		os.Exit(1)
	}

	fmt.Printf("Found %d solution(s):\n\n", len(solutions))

	solutions[0].Output(os.Stdout)
}
