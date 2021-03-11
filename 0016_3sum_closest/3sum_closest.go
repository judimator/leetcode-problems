package threesum_closest

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]

	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return sum
			}

			if abs(target-sum) < abs(target-closest) {
				closest = sum
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closest
}
