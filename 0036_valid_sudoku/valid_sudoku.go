package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	block := [3][3][9]bool{}
	col := [9][9]bool{}
	row := [9][9]bool{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}

			board[i][j] -= '1' // i.e. array length is [3][3]

			if block[i/3][j/3][board[i][j]] {
				return false
			}

			if col[i][board[i][j]] {
				return false
			}

			if row[j][board[i][j]] {
				return false
			}

			block[i/3][j/3][board[i][j]] = true
			col[i][board[i][j]] = true
			row[j][board[i][j]] = true
		}
	}

	return true
}
