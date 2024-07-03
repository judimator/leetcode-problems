package generate_parentheses

func generateParenthesis(n int) []string {
	var result []string
	var fn func(string string, open int, close int)

	fn = func(string string, open int, close int) {
		if len(string) == 2*n {
			result = append(result, string)

			return
		}

		if n > open {
			fn(string+"(", open+1, close)
		}

		if close < open {
			fn(string+")", open, close+1)
		}
	}

	fn("(", 1, 0)

	return result
}
