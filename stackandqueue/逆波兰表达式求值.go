package stackandqueue

import (
	"strconv"
)

func EvalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		// 字符串转化为整数
		val, err := strconv.Atoi(token)
		// 整数入栈
		if err == nil {
			stack = append(stack, val)
		} else {
			// 出现错误，字符串不是整数，是操作符
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
