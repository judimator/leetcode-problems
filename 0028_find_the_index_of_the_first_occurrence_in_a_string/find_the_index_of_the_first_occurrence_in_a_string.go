package find_the_index_of_the_first_occurrence_in_a_string

func strStr(haystack string, needle string) int {
	i := 0

	for i <= len(haystack)-1 {
		for j := 0; j <= len(needle)-1; j++ {
			if i < len(haystack) && haystack[i] == needle[j] {
				if j == len(needle)-1 {
					return i - j
				}

				i++
			} else {
				i = i - j + 1

				break
			}
		}
	}

	return -1
}
