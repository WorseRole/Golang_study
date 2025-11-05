package title1

// 用map实现
func SingleNumber(nums []int) int {
	var tempMap map[int]int = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if num, ok := tempMap[nums[i]]; ok {
			tempMap[nums[i]] = num + 1
		} else {
			tempMap[nums[i]] = 1
		}
	}
	for key, value := range tempMap {
		if value == 1 {
			return key
		}
	}
	return 0
}

// func SingleNumber(nums []int) int {
// 	n := len(nums)

// Outer:
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			if i == j {
// 				continue
// 			}
// 			if nums[i] == nums[j] {
// 				continue Outer
// 			}
// 		}
// 		return nums[i]
// 	}
// 	return 0
// }
