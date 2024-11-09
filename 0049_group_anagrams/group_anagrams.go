package group_anagrams

import (
	"slices"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	hashmap := make(map[string][]string)

	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		sorted := strings.Join(chars, "")

		hashmap[sorted] = append(hashmap[sorted], str)
	}

	var result [][]string

	for _, value := range hashmap {
		sort.Strings(value)
		result = append(result, value)
	}

	slices.SortFunc(result, func(a, b []string) int {
		lenA := len(a)
		lenB := len(b)

		if lenA == lenB {
			return 0
		}

		if lenA < lenB {
			return -1
		}

		return 1
	})

	return result
}
