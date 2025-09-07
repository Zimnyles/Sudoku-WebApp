package sudoku

type SudokuPair struct {
	Puzzle   *SudokuGrid
	Inverted *SudokuGrid
	Solution *SudokuGrid
}

type SudokuGrid struct {
	Grid [9][9]Cell
}

type Cell struct {
	Value uint8
}

const blockSize = 3
const rowLimit = 9
const colLimit = 9
const N = 9
