package palindrome_number

func isPalindrome(x int) bool {
	outputSlice := make([]int, 0)
	tmpVar := x

	if x < 0 {
		return false
	}

	for {
		rightMost := tmpVar % 10
		outputSlice = append(outputSlice, rightMost)
		tmpVar = tmpVar / 10

		if tmpVar <= 0 {
			break
		}
	}

	j := 0

	for i := len(outputSlice); i > 0; i-- {
		if outputSlice[i-1] != outputSlice[j] {
			return false
		}

		j++
	}

	return true
}
