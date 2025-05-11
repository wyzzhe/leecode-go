package stackandqueue

// 切片实现栈
type MyStack []int

// 压栈
func (s *MyStack) Push(v int) {
	*s = append(*s, v)
}

// 出栈
func (s *MyStack) Pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

// 查看栈顶元素
func (s *MyStack) Peek() int {
	return (*s)[len(*s)-1]
}

// 查看栈大小
func (s *MyStack) Size() int {
	return len(*s)
}

// 查看栈是否为空
func (s *MyStack) Empty() bool {
	return s.Size() == 0
}

// 栈实现队列
type MyQueue struct {
	stackIn  *MyStack
	stackOut *MyStack
}

// 初始化队列函数
func QueueConstructor() MyQueue {
	return MyQueue{
		stackIn:  &MyStack{},
		stackOut: &MyStack{},
	}
}

// 入队
func (q *MyQueue) Push(x int) {
	q.stackIn.Push(x)
}

// 出队
func (q *MyQueue) Pop() int {
	q.fillStackOut()
	return q.stackOut.Pop()
}

// 查看队头元素
func (q *MyQueue) Peek() int {
	q.fillStackOut()
	return q.stackOut.Peek()
}

// 查看队列是否为空
func (q *MyQueue) Empty() bool {
	return q.stackIn.Empty() && q.stackOut.Empty()
}

// fillStackOut 填充队列的输出栈
func (q *MyQueue) fillStackOut() {
	// 队列的出栈为空时，才会从队列的入栈拿取元素放入队列的出栈
	if q.stackOut.Empty() {
		for !q.stackIn.Empty() {
			val := q.stackIn.Pop()
			q.stackOut.Push(val)
		}
	}
}
