package stackandqueue

func IsValid(s string) bool {
	// 切片实现栈
	stack := make([]rune, 0)

	// m 用于记录某个右括号对应的左括号
	m := make(map[rune]rune)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'

	// 遍历字符串中的rune
	for _, c := range s {
		// 左括号直接入栈
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			// 如果是右括号，先判断栈内是否还有元素 -> 多余右括号
			if len(stack) == 0 {
				return false
			}
			// 再判断栈顶元素是否匹配 -> 左右括号不匹配
			peek := stack[len(stack)-1]
			if peek != m[c] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		}
	}

	// 栈中不包含元素，则是有效的括号 -> 多余左括号
	return len(stack) == 0
}
