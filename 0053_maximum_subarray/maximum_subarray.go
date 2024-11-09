package maximum_subarray

func maxSubArray(nums []int) int {
	var maxSub = nums[0]
	var curSum = 0

	for _, num := range nums {
		if curSum < 0 {
			curSum = 0
		}

		curSum += num
		maxSub = max(maxSub, curSum)
	}

	return maxSub
}

//func maxSubArray(nums []int) int {
//	sum := nums[0]
//l := len(nums)
//']
//	for i := 0; i < l; i++ {
//		tmpSum := 0
//		left := i
//
//		for lef t < l {
//			tmpSum = tmpSum + nums[left]
//
//			if tmpSum > sum {
//				sum = tmpSum
//			}
//
//			left++
//		}
//	}
//
//	return sum
//}
