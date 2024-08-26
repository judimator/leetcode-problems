package combination_sum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var result [][]int
	var tmpResult []int
	var dfs func(i int, target int)

	dfs = func(i int, target int) {
		if target == 0 {
			valid := make([]int, len(tmpResult))
			copy(valid, tmpResult)
			result = append(result, valid)

			return
		}

		if i >= len(candidates) || target < candidates[i] {
			return
		}

		for j := i; j < len(candidates); j++ {
			tmpResult = append(tmpResult, candidates[j])
			dfs(j, target-candidates[j])
			tmpResult = tmpResult[0 : len(tmpResult)-1]
		}
	}

	dfs(0, target)

	return result
}
