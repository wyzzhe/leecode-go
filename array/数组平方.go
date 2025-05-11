package array

func SortedSquares(nums []int) []int {
	// 初始化新数组和新数组尾部的指针
	ans := make([]int, len(nums))
	k := len(nums) - 1
	// 指针i和j指向的值的平方比较大小，大的放入新数组
	for i, j := 0, len(nums)-1; i <= j; {
		lsq, rsq := nums[i]*nums[i], nums[j]*nums[j]
		if lsq > rsq {
			ans[k] = lsq
			i++
		} else {
			ans[k] = rsq
			j--
		}
		k--
	}
	return ans
}
