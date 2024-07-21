package substring_with_concatenation_of_all_words

func findSubstring(s string, words []string) []int {
	lenWord := len(words[0])
	lenWords := len(words)
	lenString := len(s)
	result := make([]int, 0)

	if lenString < lenWord*lenWords {
		return result
	}

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	for i := 0; i < lenWord; i++ {
		left, right := i, i
		wordSeen := make(map[string]int)

		for right <= lenString-lenWord {
			word := s[right : right+lenWord]
			right += lenWord

			if _, exists := wordCount[word]; exists {
				wordSeen[word]++

				for wordSeen[word] > wordCount[word] {
					leftWord := s[left : left+lenWord]
					left += lenWord
					wordSeen[leftWord]--
				}

				if right-left == lenWord*lenWords {
					result = append(result, left)
				}
			} else {
				wordSeen = make(map[string]int)
				left = right
			}
		}
	}

	return result
}
