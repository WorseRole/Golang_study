package title1

// 给定一个表示 大整数 的整数数组 digits，
// 其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组。

func PlusOne(digits []int) []int {
	// 从 len大到小 进行往前运算 大于10则 往前进1 否则则跳出
	// 直接在 digits上搞就完事了
	for i := len(digits) - 1; i >= 0; i-- {
		// 只需要找到边界条件，就是最后一位是否为9
		if i == 0 && digits[i] == 9 {
			// 如果是的话，就需要扩一位
			// 后面默认全是0
			tempResult := make([]int, len(digits)+1)
			tempResult[i] = 1
			return tempResult
		}

		temp := digits[i] + 1
		if temp >= 10 {
			digits[i] = temp % 10
			continue
		} else {
			digits[i] = temp
			break
		}
	}
	return digits
}
