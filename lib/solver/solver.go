package solver

import "github.com/Crystalix007/sudoku.git/lib/board"

func Solve(b board.Board) []board.Board {
	empty := b.Empty()
	return solveRecursive(b, empty)
}

func solveRecursive(b board.Board, empty []board.Coordinate) []board.Board {
	solutions := []board.Board{}

	if len(empty) == 0 {
		return append(solutions, b.Clone())
	}

	toFill, empty := empty[0], empty[1:]

	options := b.Options(toFill.X, toFill.Y)

	for _, option := range options {
		if err := b.Set(toFill.X, toFill.Y, option); err == nil {
			newSolutions := solveRecursive(b, empty)
			solutions = append(solutions, newSolutions...)
		}
	}

	b.Set(toFill.X, toFill.Y, 0)

	return solutions
}
