package array

func RemoveElementArray(nums []int, val int) int {
	// 初始化慢指针
	slow := 0
	// 用快指针查找元素，慢指针代表新数组下标
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
