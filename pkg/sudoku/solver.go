package sudoku

func (s *SudokuGrid) solve(count *int) bool {
	if *count > 1 {
		return false
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if s.Grid[r][c].Value == 0 {
				for num := 1; num <= 9; num++ {
					if s.canPlace(r, c, uint8(num)) {
						s.Grid[r][c].Value = uint8(num)
						s.solve(count)
						s.Grid[r][c].Value = 0
					}
				}
				return false
			}
		}
	}
	*count++
	return true
}

func (s *SudokuGrid) canPlace(r, c int, val uint8) bool {
	for i := 0; i < 9; i++ {
		if s.Grid[r][i].Value == val || s.Grid[i][c].Value == val {
			return false
		}
	}
	br, bc := r/3*3, c/3*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s.Grid[br+i][bc+j].Value == val {
				return false
			}
		}
	}
	return true
}
