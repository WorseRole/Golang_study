package title1

import (
	"slices"
	// "sort"
)

func Merge(intervals [][]int) [][]int {
	// 参数校验
	if len(intervals) <= 0 {
		return [][]int{}
	}

	slices.SortFunc(intervals, func(p, q []int) int {
		return p[0] - q[0]
	})
	// sort.Slice(intervals, func(i int, j int) bool {
	// 	return intervals[i][0] < intervals[j][0]
	// })

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		current := intervals[i]

		// 如果当前区间与最后一个结果区间重叠
		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			result = append(result, current)
		}
	}
	return result

}
