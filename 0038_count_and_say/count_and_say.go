package count_and_say

import "strconv"

func countAndSay(n int) string {
	fn := func(value string) string {
		if value == "" {
			return "1"
		}

		var result []string
		prev := string(value[0])
		j := 1

		for i := 1; i < len(value); i++ {
			if string(value[i]) == prev {
				j++
			} else {
				result = append(result, strconv.Itoa(j), prev)

				prev = string(value[i])
				j = 1
			}
		}

		result = append(result, strconv.Itoa(j), prev)
		output := ""

		for _, val := range result {
			output = output + val
		}

		return output
	}

	result := ""

	for i := 1; i <= n; i++ {
		result = fn(result)
	}

	return result
}
