package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	l := searchBinary(nums, target, true)
	r := searchBinary(nums, target, false)

	return []int{l, r}
}

func searchBinary(nums []int, target int, left bool) int {
	l, r := 0, len(nums)-1
	i := -1

	for l <= r {
		m := (r + l) / 2

		if nums[m] > target {
			r--
		} else if nums[m] < target {
			l++
		} else {
			i = m

			if left {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return i
}
