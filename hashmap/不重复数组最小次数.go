package hashmap

// minOperations 函数计算使数组中所有元素不同的最小操作次数
func MinOperations(nums []int, k int) int {
	// 创建一个映射，用于记录每个数值出现的次数
	countMap := make(map[int]int)
	// 遍历数组，记录每个数值的出现次数
	for _, num := range nums {
		countMap[num]++
	}

	// 初始化操作次数为0
	result := 0
	// 遍历数组中的每个元素
	for i, num := range nums {
		// 如果当前元素出现的次数大于1，需要进行操作
		if countMap[num] > 1 {
			// 减少当前数值的出现次数
			countMap[num]--
			// 初始化新数值为当前数值加k
			newNum := num + k
			// 找到一个未被使用的数值
			for countMap[newNum] > 0 {
				newNum += k
			}
			// 更新数组中的元素为新数值
			nums[i] = newNum
			// 更新映射中新数值的出现次数
			countMap[newNum]++
			// 增加操作次数
			result += (newNum - num) / k
		}
	}
	// 返回最小操作次数
	return result
}
