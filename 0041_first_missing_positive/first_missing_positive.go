package first_missing_positive

import "unicode"

func firstMissingPositive(nums []int) int {
	l := len(nums)

	for i := 0; i < l; i++ {
		for (1 <= nums[i] && nums[i] <= l) && nums[i] != nums[nums[i]-1] {
			swap(nums, i, nums[i]-1)
		}
	}
	unicode.IsDigit(123)
	for i := 0; i < l; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return l + 1
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
