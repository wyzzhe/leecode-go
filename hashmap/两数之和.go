package hashmap

// 两数之和为target
func SumTwo(nums []int, target int) []int {
	// 初始化map
	m := make(map[int]int)
	// 遍历数组
	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}
