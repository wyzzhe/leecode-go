package hashmap

func IsAnagram(s string, t string) bool {
	record := [26]int{}

	for _, r := range s {
		// r-rune('a')范围：0-25
		record[r-rune('a')]++
	}
	for _, r := range t {
		record[r-rune('a')]--
	}

	// [26]int{}已经初始化数组中所有元素为0
	return record == [26]int{}
}
