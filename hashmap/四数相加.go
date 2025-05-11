package hashmap

// 四数相加为0
func PlusFour(nums1, nums2, nums3, nums4 []int) int {
	// 初始化m存放ab之和，count存放abcd之和数量
	m := make(map[int]int)
	count := 0
	// 先遍历a，b数组存入ab之和
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	// 遍历c，d数组，判断abcd之和是否为0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			if countVal, ok := m[0-(v3+v4)]; ok {
				// ab 与 cd相加等于0时，ab有countVal种搭配，cd每出现一次搭配种数翻一倍，会有(ab)*(cd)种搭配
				count += countVal
			}
		}
	}
	return count
}
