package stringLee

func ReverseString(s []byte) {
	// 初始化双指针
	left := 0
	right := len(s) - 1
	// 遍历字节数组
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
