package sudoku

import (
	"math/rand"
	"time"
)

func (s *SudokuGrid) MakePuzzle(emptyCells int) {
	positions := make([][2]int, 0, 81)
	for row := 0; row < rowLimit; row++ {
		for col := 0; col < colLimit; col++ {
			positions = append(positions, [2]int{row, col})
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(positions), func(i, j int) { positions[i], positions[j] = positions[j], positions[i] })

	for _, pos := range positions {
		if emptyCells <= 0 {
			break
		}

		row, col := pos[0], pos[1]
		backup := s.Grid[row][col].Value
		s.Grid[row][col].Value = 0

		count := 0
		copyGrid := s.copy()
		copyGrid.solve(&count)
		if count != 1 {
			s.Grid[row][col].Value = backup
		} else {
			emptyCells--
		}
	}
}

func (s *SudokuGrid) copy() *SudokuGrid {
	newGrid := &SudokuGrid{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			newGrid.Grid[r][c].Value = s.Grid[r][c].Value
		}
	}
	return newGrid
}

func (s *SudokuGrid) InvertPuzzle(solution *SudokuGrid) *SudokuGrid {
	inverted := &SudokuGrid{}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if s.Grid[r][c].Value == 0 {
				inverted.Grid[r][c].Value = solution.Grid[r][c].Value
			} else {
				inverted.Grid[r][c].Value = 0
			}
		}
	}

	return inverted
}
