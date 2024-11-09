package rotate_image

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix[0])
	visited := make(map[string]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			swap(n, matrix, visited, matrix[i][j], i, j)
		}
	}
}

func swap(n int, matrix [][]int, visited map[string]bool, value int, i int, j int) {
	newI := j
	newJ := n - i - 1

	key := fmt.Sprintf("%d_%d", newI, newJ)

	if _, ok := visited[key]; ok {
		return
	}

	nextValue := matrix[newI][newJ]

	matrix[newI][newJ] = value
	visited[key] = true

	swap(n, matrix, visited, nextValue, newI, newJ)
}
