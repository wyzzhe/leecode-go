package stringLee

import "strings"

func RepeatedSubStringKMP(s string) bool {
	// 判断字符串不为空
	if len(s) == 0 {
		return false
	}
	next := make([]int, len(s))
	// 查找前缀表，传入模式串
	getNextReSubString(next, s)
	// 最长相同前后缀的长度能整除则是重复子串
	if next[len(s)-1] != 0 && len(s)%(len(s)-next[len(s)-1]) == 0 {
		return true
	}

	return false
}

func getNextReSubString(next []int, s string) {
	// 初始化前缀指针
	j := 0
	// 初始化next数组
	next[0] = j
	// 遍历字符串
	for i := 1; i < len(s); i++ {
		// 前后缀指针值不相等且前缀指针不为0，前缀指针循环回退
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		// 前后缀指针值相等，前缀指针后移
		if s[i] == s[j] {
			j++
		}
		// 前缀表赋值为前缀指针的位置
		next[i] = j
	}
}

// 移动匹配
func RepeatedSubString(s string) bool {
	if len(s) == 0 {
		return false
	}
	t := s + s
	return strings.Contains(t[1:len(t)-1], s)
}
