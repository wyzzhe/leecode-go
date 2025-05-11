package list

// 返回环的入口
func DetectCycle(head *ListNode) *ListNode {
	// 快慢指针
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		// 快指针走两步，慢指针走一步
		fast = fast.Next.Next
		slow = slow.Next
		// 判断是否有环
		if fast == slow {
			// 若有环，找到环的入口
			// index1在链表头
			// index2在快慢指针交点
			index1, index2 := head, fast
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index1
		}
	}
	return nil
}
