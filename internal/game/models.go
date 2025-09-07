package game

import "sudoku/pkg/sudoku"

const superEasy = 20
const easy = 30
const middle = 40
const hard = 50
const superHard = 55
const extreme = 60

type CellRequest struct {
	Row       uint8 `json:"row"`
	Col       uint8 `json:"col"`
	Value     uint8 `json:"value"`
	IsCorrect bool  `json:"isCorrect"`
}

type templateData struct {
	Grids        *sudoku.SudokuPair
	Fails        int
	PuzzleJSON   string
	SolutionJSON string
}
