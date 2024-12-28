package two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok && j != i {
			return []int{i, j}
		}
	}
	// Return an empty slice if no solution is found
	return []int{}
}
