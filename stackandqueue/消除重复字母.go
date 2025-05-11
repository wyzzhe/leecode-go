package stackandqueue

func RemoveDuplicates(s string) string {
	// 字符串实现栈
	var stack []byte
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		// 栈不空且栈中字符匹配
		if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			// 栈为空或栈中字符不匹配
			// 入栈
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
