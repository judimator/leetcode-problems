package string_to_integer_atoi

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	sign := 1
	res := 0
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return res
	}

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			break
		}
		dig := int(ch - '0')
		res = res*10 + dig

		if sign*res < math.MinInt32 {
			return math.MinInt32
		} else if sign*res > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return sign * res
}
