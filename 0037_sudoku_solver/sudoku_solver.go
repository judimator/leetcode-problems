package sudoku_solver

func solveSudoku(board [][]byte) {
	var empty [][]int
	var block [3][3][9]bool
	var col [9][9]bool
	var row [9][9]bool
	var solved bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				empty = append(empty, []int{i, j})
			} else {
				value := board[i][j] - 49

				block[i/3][j/3][value] = true
				row[i][value] = true
				col[j][value] = true
			}
		}
	}

	var dfs func(k int)
	dfs = func(k int) {
		if k == len(empty) {
			solved = true

			return
		}

		i := empty[k][0]
		j := empty[k][1]

		for z := 0; z < 9; z++ {
			if row[i][z] == false && col[j][z] == false && block[i/3][j/3][z] == false {
				row[i][z], col[j][z], block[i/3][j/3][z] = true, true, true
				board[i][j] = byte(z + 49)

				dfs(k + 1)

				if solved {
					return
				}

				row[i][z], col[j][z], block[i/3][j/3][z] = false, false, false
			}
		}
	}

	dfs(0)
}
