package stackandqueue

// 单调队列定义
type MonoQueue struct {
	queue []int
}

// 单调队列初始化
func NewMonoQueue() *MonoQueue {
	return &MonoQueue{
		queue: make([]int, 0),
	}
}

// 查看队头元素
func (m *MonoQueue) Front() int {
	return m.queue[0]
}

// 查看队尾元素
func (m *MonoQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

// 查看队列是否为空
func (m *MonoQueue) Empty() bool {
	return len(m.queue) == 0
}

// 单调队列入队
func (m *MonoQueue) Push(val int) {
	// 新加入的元素比队尾元素大，队尾元素出队
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

// 单调队列出队 -> 移除上一个窗口的第一个元素
func (m *MonoQueue) Pop(val int) {
	// 上一个窗口的第一个元素是最大值，需要被移除
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func MaxSlidingWindow(nums []int, k int) []int {
	// 初始化队列
	queue := NewMonoQueue()
	length := len(nums)
	res := make([]int, 0)
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	// 记录前k个元素的最大值
	res = append(res, queue.Front())

	// 开始滑动窗口
	for i := k; i < length; i++ {
		// 移除上一个窗口的第一个元素
		queue.Pop(nums[i-k])
		// 插入当前窗口的最后一个元素
		queue.Push(nums[i])
		// 记录每轮滑动窗口的最大值
		res = append(res, queue.Front())
	}

	return res
}
