package title1

func Twosum(nums []int, target int) []int {
	// 初始化
	m1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		key := target - nums[i]
		if value, exists := m1[key]; exists {
			return []int{value, i}
		}
		m1[nums[i]] = i
	}
	return []int{}
}
