package longest_palindromic_substring

func longestPalindrome(s string) string {
	check := func(i, j int) bool {
		left := i
		right := j - 1

		for left < right {
			if s[left] != s[right] {
				return false
			}

			left++
			right--
		}

		return true
	}

	for length := len(s); length > 0; length-- {
		for start := 0; start <= len(s)-length; start++ {
			if check(start, start+length) {
				return s[start : start+length]
			}
		}
	}

	return ""
}
