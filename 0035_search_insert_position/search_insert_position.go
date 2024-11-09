package search_insert_position

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var m int

	for l <= r {
		m = (r + l) / 2

		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			r--
		} else {
			l++
		}
	}

	if nums[m] > target {
		return m
	} else {
		return m + 1
	}
}
