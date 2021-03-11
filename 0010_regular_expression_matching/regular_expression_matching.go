package regular_expression_matching

import "fmt"

func isMatch(text string, pattern string) bool {
	memo := make(map[string]bool)

	var dp func(int, int) bool

	dp = func(i, j int) bool {
		key := fmt.Sprintf("%d,%d", i, j)

		if val, ok := memo[key]; ok {
			return val
		}

		if j == len(pattern) {
			ans := i == len(text)
			memo[key] = ans

			return ans
		}

		firstMatch := i < len(text) && (pattern[j] == '.' || pattern[j] == text[i])

		if j+1 < len(pattern) && pattern[j+1] == '*' {
			ans := dp(i, j+2) || (firstMatch && dp(i+1, j))
			memo[key] = ans

			return ans
		}

		ans := firstMatch && dp(i+1, j+1)
		memo[key] = ans

		return ans
	}

	return dp(0, 0)
}
