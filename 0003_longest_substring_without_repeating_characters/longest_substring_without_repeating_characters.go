package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	hashmap := make(map[string]int)
	maxLength := 0

	for r < len(s) {
		k := string(s[r])

		if _, ok := hashmap[k]; ok {
			maxLength = max(maxLength, r-l)
			l = max(l, hashmap[k]+1)
		}

		hashmap[k] = r
		r += 1
	}

	return max(maxLength, r-l)
}
