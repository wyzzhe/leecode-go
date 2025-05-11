package list

import "fmt"

func ReverseList(head *ListNode) *ListNode {
	// 初始化pre, cur
	var pre *ListNode
	fmt.Println(pre)
	cur := head

	// 遍历链表直到cur==nil
	for cur != nil {
		// 初始化temp指向cur的下一个节点
		temp := cur.Next
		// 反转链表
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
