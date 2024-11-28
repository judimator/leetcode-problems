package spiral_matrix_ii

func generateMatrix(n int) [][]int {
	res := make([][]int, n)

	for i := range res {
		res[i] = make([]int, n)
	}

	startRow, startCol, endRow, endCol := 0, 0, n-1, n-1
	curr := 1
	finalVal := n * n
	for startRow <= endRow {
		for i := startCol; i <= endCol && curr <= finalVal; i++ {
			res[startRow][i] = curr
			curr++
		}

		for i := startRow + 1; i <= endRow && curr <= finalVal; i++ {
			res[i][endCol] = curr
			curr++
		}

		for i := endCol - 1; i >= startCol && curr <= finalVal; i-- {
			res[endRow][i] = curr
			curr++
		}

		for i := endRow - 1; i > startRow && curr <= finalVal; i-- {
			res[i][startCol] = curr
			curr++
		}

		startRow, startCol, endRow, endCol = startRow+1, startCol+1, endRow-1, endCol-1
	}

	return res
}
