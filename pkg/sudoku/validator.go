package sudoku

func (s *SudokuGrid) IsValid() bool {
	N := 9

	for r := 0; r < N; r++ {
		seen := [10]bool{}
		for c := 0; c < N; c++ {
			val := s.Grid[r][c].Value
			if val < 1 || val > 9 || seen[val] {
				return false
			}
			seen[val] = true
		}
	}

	for c := 0; c < N; c++ {
		seen := [10]bool{}
		for r := 0; r < N; r++ {
			val := s.Grid[r][c].Value
			if val < 1 || val > 9 || seen[val] {
				return false
			}
			seen[val] = true
		}
	}

	for br := 0; br < 3; br++ {
		for bc := 0; bc < 3; bc++ {
			seen := [10]bool{}
			for r := 0; r < 3; r++ {
				for c := 0; c < 3; c++ {
					val := s.Grid[br*3+r][bc*3+c].Value
					if val < 1 || val > 9 || seen[val] {
						return false
					}
					seen[val] = true
				}
			}
		}
	}

	return true
}
