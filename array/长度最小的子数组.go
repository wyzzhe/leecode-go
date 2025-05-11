package array

func MinSubArrayLen(nums []int, target int) int {
	// 初始化头指针和sum、result
	i := 0
	sum, result := 0, len(nums)+1
	subLength := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			subLength = j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == len(nums)+1 {
		return 0
	} else {
		return result
	}
}
