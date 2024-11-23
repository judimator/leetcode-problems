package length_of_last_word

func lengthOfLastWord(s string) int {
	c := 0
	start := false

	for i := len(s) - 1; i >= 0; i-- {
		if start == false && s[i] == ' ' {
			continue
		}

		start = true

		if s[i] == ' ' {
			return c
		}

		c++
	}

	return c
}
