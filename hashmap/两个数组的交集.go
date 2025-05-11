package hashmap

// 数组方法，适用于传入的数组数字在1000以内
func IntersectionArray(nums1 []int, nums2 []int) []int {
	count1 := make([]int, 1001)
	count2 := make([]int, 1001)
	res := make([]int, 0)
	for _, v := range nums1 {
		count1[v] = 1
	}
	// 添加数字的操作不在nums2的循环中处理，防止重复
	for _, v := range nums2 {
		count2[v] = 1
	}
	for i := 0; i <= 1000; i++ {
		if count1[i]+count2[i] == 2 {
			res = append(res, i)
		}
	}
	return res
}

// hashmap方法，适用于传入数组的数字相当大(自动去重)
// 当传入 [0, 1000000001] 时使用数组会导致大量的空间浪费
func IntersectionHashMap(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0) // 用map模拟set
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			// struct{}{}代表空结构体，用作set[v]存在的标记符
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		// nums2中的数字存在nums1中，加入结果集，并清空该set值
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
