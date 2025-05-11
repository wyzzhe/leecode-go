package hashmap

import "sort"

// 三数之和为0
func SumThree(nums []int) [][]int {
	// 对数组排序
	sort.Ints(nums)
	// 初始化双指针和结果数组
	res := make([][]int, 0)
	// 遍历数组
	for i := 0; i < len(nums)-2; i++ {
		// 每轮第一个数大于0，则无三数之和为0的结果
		if nums[i] > 0 {
			break
		}
		// 第二轮之后对i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 初始化左右指针
		left, right := i+1, len(nums)-1
		// i不变时执行一轮遍历，动态调整左右指针
		for left < right {
			// 记录左右指针初始指向的值
			ln, rn := nums[left], nums[right]
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 找到一个三元组之后先移动左右指针并检测去重
				for left < right && nums[left] == ln {
					left++
				}
				for left < right && nums[right] == rn {
					right--
				}
			}
		}
		// l, r := i+1, len(nums)-1
		// for l < r {
		// 	n2, n3 := nums[l], nums[r]
		// 	if nums[i]+n2+n3 == 0 {
		// 		res = append(res, []int{nums[i], n2, n3})
		// 		// 去重逻辑应该放在找到一个三元组之后，对b 和 c去重
		// 		for l < r && nums[l] == n2 {
		// 			l++
		// 		}
		// 		for l < r && nums[r] == n3 {
		// 			r--
		// 		}
		// 	} else if nums[i]+n2+n3 < 0 {
		// 		l++
		// 	} else {
		// 		r--
		// 	}
		// }
	}
	return res
}
