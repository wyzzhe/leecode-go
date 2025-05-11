package stringLee

func ReverseWords(s string) string {
	str := []byte(s)
	// 初始化双指针
	slow, fast := 0, 0
	// 去除开头空格
	for len(str) > 0 && fast < len(str) && str[fast] == ' ' {
		fast++
	}
	// 从第二个单词开始，删除单词间冗余空格
	for ; fast < len(str); fast++ {
		if fast-1 > 0 && str[fast] == str[fast-1] && str[fast] == ' ' {
			continue
		}
		str[slow] = s[fast]
		slow++
	}
	// 删除尾部冗余空格
	if slow-1 > 0 && str[slow-1] == ' ' {
		str = str[:slow-1]
	} else {
		str = str[:slow]
	}
	// 反转整个字符串
	reverse(str)
	// 反转单个单词
	i := 0
	for i < len(str) {
		j := i
		// j走到单个单词的末尾后一位
		for ; j < len(str) && str[j] != ' '; j++ {
		}
		reverse(str[i:j])
		i = j
		i++
	}
	return string(str)
}

func reverse(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
