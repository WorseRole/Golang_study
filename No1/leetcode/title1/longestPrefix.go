package title1

func LongestCommonPrefix(strs []string) string {
	// 边界条件
	if len(strs) <= 0 {
		return ""
	}
	// 先定义一个 假设最长公共  rune 对string进行切片
	var result []rune = []rune{}
	result = []rune(strs[0])

	for i := 1; i < len(strs); i++ {
		l := len(result)
		temp := []rune(strs[i])
		if len(temp) < len(result) {
			l = len(temp)
		}
		j := 0
		for ; j < l; j++ {
			// 公共前缀
			if temp[j] != result[j] {
				break
			}
		}
		result = result[0:j]
	}
	return string(result)
}
