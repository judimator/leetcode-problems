package combination_sum_ii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var result [][]int
	var tmpResult []int
	var dfs func(i int, target int)

	dfs = func(i int, target int) {
		if target == 0 {
			result = append(result, append([]int{}, tmpResult...))

			return
		}

		for j := i; j < len(candidates); j++ {
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}

			if candidates[j] > target {
				return
			}

			tmpResult = append(tmpResult, candidates[j])
			dfs(j+1, target-candidates[j])
			tmpResult = tmpResult[0 : len(tmpResult)-1]
		}
	}

	dfs(0, target)

	return result
}
