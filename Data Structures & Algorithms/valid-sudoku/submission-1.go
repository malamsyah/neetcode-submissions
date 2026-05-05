func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !validRow(board, i){
			return false
		}
		
		if !validCol(board, i) {
			return false
		}
	}

	m := make(map[int]map[byte]struct{})
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}

			idx := (row / 3) * 3 + (col / 3)
			if m[idx] == nil {
				m[idx] = make(map[byte]struct{})
			}

			_, found := m[idx][board[row][col]]
			if found {
				return false
			}

			m[idx][board[row][col]] = struct{}{}
		}
	}

	return true
}

func validRow(board [][]byte, i int) bool {
	m := make(map[byte]struct{})
	for j := 0; j < 9; j++ {
		if board[i][j] == '.' {
			continue
		}

		_, found := m[board[i][j]]
		if found {
			return false
		}

		m[board[i][j]] = struct{}{}
	}

	return true
}

func validCol(board [][]byte, i int) bool {
	m := make(map[byte]struct{})
	for j := 0; j < 9; j++ {
		if board[j][i] == '.' {
			continue
		}

		_, found := m[board[j][i]]
		if found {
			return false
		}

		m[board[j][i]] = struct{}{}
	}

	return true
}
