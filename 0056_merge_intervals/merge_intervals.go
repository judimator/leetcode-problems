package merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	var result [][]int

	l := len(intervals)

	if l == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i := 0
	tmpI := intervals[0][0]
	tmpJ := intervals[0][1]
	startInterval := true

	for i < l-1 {
		if startInterval == true {
			tmpI = intervals[i][0]
			tmpJ = intervals[i][1]
		}

		next := intervals[i+1]

		if tmpJ >= next[0] {
			tmpI = min(tmpI, next[0])
			tmpJ = max(tmpJ, next[1])

			startInterval = false
		} else {
			result = append(result, []int{tmpI, tmpJ})
			startInterval = true
		}

		i++
	}

	if startInterval == true {
		result = append(result, []int{intervals[l-1][0], intervals[l-1][1]})

		return result
	}

	result = append(result, []int{tmpI, tmpJ})

	return result
}
