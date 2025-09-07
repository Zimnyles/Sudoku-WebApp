package sudoku

import (
	"math/rand"
)

func NewSudokuPair(emptyCells int) *SudokuPair {
	solution := NewEmptyGrid()
	solution.Randomize()  
	solution.MakePuzzle(0) 

	puzzle := solution.copy() 
	puzzle.MakePuzzle(emptyCells)

	inverted := puzzle.InvertPuzzle(solution)

	return &SudokuPair{
		Puzzle:   puzzle,
		Inverted: inverted,
		Solution: solution,
	}
}

func NewEmptyGrid() *SudokuGrid {
	s := &SudokuGrid{}
	s.baseGrid()
	return s
}

func (s *SudokuGrid) Randomize() {

	s.baseGrid()
	s.shuffleDigits()
	s.shuffleRows()
	s.shuffleRowBlocks()
	s.swapCols()
	s.swapRows()

	for i := 0; i < 10; i++ {
		s.swapRows()
		s.swapCols()
	}
	if rand.Intn(2) == 0 {
		s.mirrorHorizontal()
	}
	if rand.Intn(2) == 0 {
		s.mirrorVertical()
	}
}

func (s *SudokuGrid) baseGrid() {
	for row := 0; row < rowLimit; row++ {
		for col := 0; col < colLimit; col++ {
			s.Grid[row][col].Value = uint8((row*blockSize+row/blockSize+col)%N + 1)
		}
	}
}

func (s *SudokuGrid) shuffleDigits() {
	perm := rand.Perm(9)
	for row := 0; row < rowLimit; row++ {
		for col := 0; col < colLimit; col++ {
			s.Grid[row][col].Value = uint8(perm[s.Grid[row][col].Value-1] + 1)
		}
	}
}

func (s *SudokuGrid) shuffleRows() {
	for block := 0; block < blockSize; block++ {
		perm := rand.Perm(3)
		start := block * 3
		tmp := [3][9]Cell{}
		for i := 0; i < 3; i++ {
			tmp[i] = s.Grid[start+perm[i]]
		}
		for i := 0; i < 3; i++ {
			s.Grid[start+i] = tmp[i]
		}
	}
}

func (s *SudokuGrid) shuffleRowBlocks() {
	perm := rand.Perm(3)
	tmp := [9][9]Cell{}
	for i := 0; i < 3; i++ {
		copy(tmp[i*3 : (i+1)*3][:], s.Grid[perm[i]*3 : (perm[i]+1)*3][:])
	}
	s.Grid = tmp
}

func (s *SudokuGrid) swapRows() {
	block := rand.Intn(3)
	r1 := rand.Intn(3)
	r2 := rand.Intn(3)
	if r1 != r2 {
		s.Grid[block*3+r1], s.Grid[block*3+r2] = s.Grid[block*3+r2], s.Grid[block*3+r1]
	}
}

func (s *SudokuGrid) swapCols() {
	block := rand.Intn(3)
	c1 := rand.Intn(3)
	c2 := rand.Intn(3)
	if c1 != c2 {
		for r := 0; r < 9; r++ {
			s.Grid[r][block*3+c1], s.Grid[r][block*3+c2] = s.Grid[r][block*3+c2], s.Grid[r][block*3+c1]
		}
	}
}

func (s *SudokuGrid) mirrorHorizontal() {
	for r := 0; r < 4; r++ {
		for c := 0; c < 9; c++ {
			s.Grid[r][c], s.Grid[8-r][c] = s.Grid[8-r][c], s.Grid[r][c]
		}
	}
}

func (s *SudokuGrid) mirrorVertical() {
	for row := 0; row < rowLimit; row++ {
		for col := 0; col < 4; col++ {
			s.Grid[row][col], s.Grid[row][8-col] = s.Grid[row][8-col], s.Grid[row][col]
		}
	}
}
