package sudoku

import "fmt"

func (s *SudokuGrid) PrintPuzzle() {
	for row := 0; row < rowLimit; row++ {
		for col := 0; col < colLimit; col++ {
			val := s.Grid[row][col].Value
			if val == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", val)
			}
		}
		fmt.Println()
	}
}

func (p *SudokuPair) PrintPair() {
	fmt.Println("Puzzle:")
	p.Puzzle.PrintPuzzle()
	fmt.Println("Inverted:")
	p.Inverted.PrintPuzzle()
	fmt.Println("Solution:")
	p.Solution.PrintPuzzle()
	fmt.Println("Empty:")
	p.Empty.PrintPuzzle()
}
