package array

func SplitSearch(nums []int, target int) int {
	// 初始化左右边界
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}
