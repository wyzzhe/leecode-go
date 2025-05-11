package list

func RemoveElementList(head *ListNode, val int) *ListNode {
	// 使用虚拟头节点
	dummyHead := &ListNode{}
	dummyHead.Next = head
	// 定义临时指针
	cur := dummyHead
	// 开始遍历链表找到val并删除
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			// 找到val则删除节点
			cur.Next = cur.Next.Next
		} else {
			// 没找到则移动临时指针
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
