package list

func ExchangeList(head *ListNode) *ListNode {
	// 初始化虚拟头节点
	dummyHead := &ListNode{}
	dummyHead.Next = head
	// 初始化cur节点
	cur := dummyHead
	// 遍历链表并两两交换
	for cur.Next != nil && cur.Next.Next != nil {
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = temp.Next.Next
		cur.Next.Next = temp
		cur = cur.Next.Next
	}
	return dummyHead.Next
}
