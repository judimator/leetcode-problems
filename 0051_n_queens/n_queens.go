package n_queens

import "strings"

func solveNQueens(n int) [][]string {
	var result [][]string
	var board = make([][]string, n)
	var diagonal = make([]int, 2*n)
	var antiDiagonal = make([]int, 2*n)
	var column = make([]int, n)
	var dfs func(i int)

	dfs = func(i int) {
		if i == n {
			tmp := make([]string, n)

			for k := 0; k < n; k++ {
				tmp[k] = strings.Join(board[k], "")
			}

			result = append(result, tmp)

			return
		}

		for j := 0; j < n; j++ {
			if column[j]+diagonal[i+j]+antiDiagonal[n-j+i] == 0 {
				board[i][j] = "Q"
				column[j], diagonal[i+j], antiDiagonal[n-j+i] = 1, 1, 1
				dfs(i + 1)
				column[j], diagonal[i+j], antiDiagonal[n-j+i] = 0, 0, 0
				board[i][j] = "."
			}
		}
	}

	for i := 0; i < n; i++ {
		board[i] = make([]string, n)

		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	for i := 0; i < 2*n; i++ {
		diagonal[i] = 0
		antiDiagonal[i] = 0
	}

	for i := 0; i < n; i++ {
		column[i] = 0
	}

	dfs(0)

	return result
}
