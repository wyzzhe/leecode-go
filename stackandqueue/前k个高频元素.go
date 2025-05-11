package stackandqueue

import "container/heap"

// 定义小顶堆
type Heap [][2]int

// 查询小顶堆长度
func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TopKFrequent(nums []int, k int) []int {
	// 初始化hash表
	mapNum := map[int]int{}

	// 记录每个元素出现的次数
	for _, num := range nums {
		mapNum[num]++
	}

	// 初始化小顶堆
	h := &Heap{}
	heap.Init(h)

	// 所有元素入堆，堆的长度为k
	for key, value := range mapNum {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	// 按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		// k-i-1 相当于倒序, 从k-1到0
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}
