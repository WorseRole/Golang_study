package title1

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

func IsValid(s string) bool {

	// 先切片 rnue从string切成rnue	类似于java中的string->char
	var rnues []rune = []rune(s)
	var stack []rune = []rune{}
	for i := range rnues {
		temp := rnues[i]
		if (temp == ')' || temp == '}' || temp == ']') && len(stack) == 0 {
			return false
		}

		switch temp {
		case '(', '{', '[':
			// 切片增加用 stack = append(stack, value)
			stack = append(stack, temp)
		case ')':
			last := stack[len(stack)-1]
			if last == '(' {
				// stack pop 操作
				stack = stack[0:(len(stack) - 1)]
			} else {
				return false
			}

		case ']':
			last := stack[len(stack)-1]
			if last == '[' {
				stack = stack[0:(len(stack) - 1)]
			} else {
				return false
			}

		case '}':
			last := stack[len(stack)-1]
			if last == '{' {
				stack = stack[0:(len(stack) - 1)]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
	// if len(stack) > 0 {
	// 	return false
	// }
	// return true
}
