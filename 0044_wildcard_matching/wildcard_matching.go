package wildcard_matching

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	var dp func(int, int) bool
	cache := make(map[string]bool)

	dp = func(i int, j int) bool {
		if i == len(s) && j == len(p) {
			return true
		}

		if j == len(p) {
			return false
		}

		key := fmt.Sprintf("%d%d", i, j)

		if value, ok := cache[key]; ok {
			return value
		}

		match := false

		if i < len(s) && (s[i] == p[j] || p[j] == '?') {
			match = dp(i+1, j+1)
		} else if p[j] == '*' {
			match = dp(i, j+1)

			if !match && i < len(s) {
				match = dp(i+1, j+1) || dp(i+1, j)
			}
		}

		cache[key] = match

		return cache[key]
	}

	return dp(0, 0)
}
