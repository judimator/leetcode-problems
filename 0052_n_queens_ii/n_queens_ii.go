package n_queens__ii

func totalNQueens(n int) int {
	var diagonal = make([]int, 2*n)
	var antiDiagonal = make([]int, 2*n)
	var column = make([]int, n)
	var dfs func(i int)
	var result int

	for i := 0; i < 2*n; i++ {
		diagonal[i] = 0
		antiDiagonal[i] = 0
	}

	for i := 0; i < n; i++ {
		column[i] = 0
	}

	dfs = func(i int) {
		if i == n {
			result = result + 1

			return
		}

		for j := 0; j < n; j++ {
			if column[j]+diagonal[i+j]+antiDiagonal[n-j+i] == 0 {
				column[j], diagonal[i+j], antiDiagonal[n-j+i] = 1, 1, 1
				dfs(i + 1)
				column[j], diagonal[i+j], antiDiagonal[n-j+i] = 0, 0, 0
			}
		}
	}

	dfs(0)

	return result
}
