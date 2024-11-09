package two_sum

func twoSum(nums []int, target int) []int {
	callback := func(nums []int, target int) []int {
		left := 0
		right := len(nums) - 1

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				return []int{left, right}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}

		return nil
	}

	return callback(nums, target)
}
