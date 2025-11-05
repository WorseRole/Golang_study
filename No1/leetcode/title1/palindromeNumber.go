package title1

func IsPalindrome(x int) bool {
	// 判断回文数 123 反转罢了
	// 123 / 10 = 12 	123 % 10 = 3 	:3
	// 12 / 10 = 1		12 % 10 = 2		:32
	// 1 / 10 = 0		1 % 10 = 1		:321

	var i int = x
	j := 0
	for i > 0 {
		j = (j * 10) + (i % 10)
		i = i / 10
	}

	return j == x

	// if j == x {
	// 	return true
	// } else {
	// 	return false
	// }
}
