package list

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// 虚拟头节点
	dummyHead := &ListNode{}
	dummyHead.Next = head
	// 初始化快慢指针
	fast, slow := dummyHead, dummyHead
	// 快指针先走n+1
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	// 快慢指针一起走，直到快指针指向nil
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 删除节点操作
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
