package permutations_ii

import "strconv"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	// Try to sort to avoid duplicates
	permutation := make([]int, len(nums))
	visit := make([]bool, len(nums))
	permute := make(map[string]bool)

	var dfs func(int)
	dfs = func(index int) {
		if index == len(nums) {
			var key string

			for _, num := range permutation {
				key = key + strconv.Itoa(num)
			}

			if _, ok := permute[key]; ok {
				return
			}

			copiedPermutation := make([]int, len(nums))
			copy(copiedPermutation, permutation)

			res = append(res, copiedPermutation)

			permute[key] = true

			return
		}

		for i := 0; i < len(nums); i++ {
			if visit[i] == false {
				visit[i] = true
				permutation[index] = nums[i]
				dfs(index + 1)
				visit[i] = false
			}
		}
	}

	dfs(0)

	return res
}
