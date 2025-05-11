package stringLee

func KMP(haystack string, needle string) int {
	// 模式串为空返回0
	if len(needle) == 0 {
		return 0
	}
	j := 0
	next := make([]int, len(needle))
	// 先对needle求前缀表
	getNext(next, needle)
	// 遍历文本串
	for i := 0; i < len(haystack); i++ {
		// 冲突时j指向前缀表中[冲突前一位下标]的值
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		// 字符匹配则继续向前比较
		if haystack[i] == needle[j] {
			j++
		}
		// 模式串全部匹配完毕，返回匹配的位置
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}

	return -1
}

// 求前缀表
func getNext(next []int, s string) {
	// 初始化j为前缀指针
	j := 0
	// 初始化next数组
	next[0] = j
	for i := 1; i < len(s); i++ {
		// ij 指向字符不相同，前缀指针j循环返回前一位前缀表的值的位置（每次都尝试可能的最优解）
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}

		// ij 指向字符相同，记录j的索引，即前后缀相等位数
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
