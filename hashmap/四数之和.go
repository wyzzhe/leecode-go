package hashmap

import "sort"

// 四数之和为target
func SumFour(nums []int, target int) [][]int {
	// 对数组排序
	sort.Ints(nums)
	// 初始化结果数组
	res := [][]int{}
	// 从k开始遍历数组
	for k := 0; k < len(nums)-3; k++ {
		// 对k剪枝
		if nums[k] > target && target > 0 {
			break
		}
		// 第二轮开始对k去重
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		// 从i开始遍历
		for i := k + 1; i < len(nums)-2; i++ {
			// 对i剪枝
			if nums[i] > target && target > 0 {
				break
			}
			// 第二轮开始对i去重
			if i > k+1 && nums[i] == nums[i-1] {
				continue
			}
			// 初始化左右指针和指向的值
			left, right := i+1, len(nums)-1
			// 动态移动左右指针并收集符合条件的数组
			for left < right {
				// 记录左右指针初始指向的值
				ln, rn := nums[left], nums[right]
				if nums[k]+nums[i]+nums[left]+nums[right] == target {
					res = append(res, []int{nums[k], nums[i], nums[left], nums[right]})
					// 找到一个四元组之后移动左右指针并检测去重
					for left < right && nums[left] == ln {
						left++
					}
					for left < right && nums[right] == rn {
						right--
					}
				} else if nums[k]+nums[i]+ln+rn < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}
