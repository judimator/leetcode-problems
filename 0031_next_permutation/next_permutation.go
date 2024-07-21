package next_permutation

func nextPermutation(nums []int) {
	l := len(nums)
	i := l - 2

	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := l - 1

		for j > i && nums[i] >= nums[j] {
			j--
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	left := i + 1
	right := l - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]

		left++
		right--
	}
}
