package stackandqueue

// 队列实现栈
type MyNewStack struct {
	queue []int // 创建一个队列
}

// 初始化栈
func StackConstructor() MyNewStack {
	return MyNewStack{
		queue: make([]int, 0),
	}
}

// 压栈
func (s *MyNewStack) Push(v int) {
	s.queue = append(s.queue, v)
}

// 出栈
func (s *MyNewStack) Pop() int {
	// 相当于遍历区间为(0, len(s.queue)-1]，保留队列最后一个元素
	n := len(s.queue) - 1
	for n != 0 {
		// 取值
		val := s.queue[0]
		// 出队
		s.queue = s.queue[1:]
		// 入队
		s.queue = append(s.queue, val)
		n--
	}
	// 出栈
	val := s.queue[0]
	s.queue = s.queue[1:]
	return val
}

// 查看栈顶元素
func (s *MyNewStack) Top() int {
	val := s.Pop()
	s.queue = append(s.queue, val)
	return val
}

// 判断栈是否为空
func (s *MyNewStack) Empty() bool {
	return len(s.queue) == 0
}
