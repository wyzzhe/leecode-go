package array

import (
	"sort"
)

func compare(a, b string) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return false
}

func KthLargestNumbers(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i], nums[j])
	})

	return nums[k-1]
}
