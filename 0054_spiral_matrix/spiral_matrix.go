package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	lenI := len(matrix)
	lenJ := len(matrix[0])
	size := lenI * lenJ

	stopI := lenI - 1
	stopJ := lenJ - 1

	k, i, j, padding, direction := 0, 0, 0, 0, 0

	var result []int

	if lenI == 2 && lenJ == 0 && size == 2 {
		return []int{matrix[0][0], matrix[0][1]}
	}

	if lenI == 2 && lenJ == 1 && size == 2 {
		return []int{matrix[0][0], matrix[1][0]}
	}

	for k < size {
		if direction == 0 {
			if j == stopJ-padding {
				direction = 1
			} else {
				result = append(result, matrix[i][j])

				k++
				j++
			}
		} else if direction == 1 {
			if i == stopI-padding {
				direction = 2
			} else {
				result = append(result, matrix[i][j])

				k++
				i++
			}
		} else if direction == 2 {
			if j == padding {
				direction = 3
			} else {
				result = append(result, matrix[i][j])

				k++
				j--
			}
		} else {
			if i == padding+1 { // new loop
				direction = 0
				padding++
			} else {
				result = append(result, matrix[i][j])

				k++
				i--
			}
		}
	}

	return result
}
