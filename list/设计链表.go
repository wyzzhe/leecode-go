package list

import "fmt"

type myLinkedList struct {
	dummyHead *ListNode // 虚拟头节点
	Size      int       // 链表大小
}

func Constructor() myLinkedList {
	newNode := &ListNode{ // 创建新节点
		Val:  -999,
		Next: nil,
	}
	return myLinkedList{ // 返回链表
		dummyHead: newNode,
		Size:      0,
	}
}

// 获得第n个节点的值
func (mll *myLinkedList) Get(index int) int {
	if mll == nil || index < 0 || index >= mll.Size { // 如果索引无效则返回-1
		return -1
	}
	// cur指向真正头节点
	cur := mll.dummyHead.Next
	for i := 0; i < index; i++ { // 遍历到索引所在的节点
		cur = cur.Next
	}
	return cur.Val // 返回节点值
}

func (mll *myLinkedList) AddAtHead(val int) {
	// 构造新加的节点
	newNode := &ListNode{Val: val}
	// 新加的节点的下一个节点指向真实头节点
	newNode.Next = mll.dummyHead.Next
	mll.dummyHead.Next = newNode
	mll.Size++
}

func (mll *myLinkedList) AddAtTail(val int) {
	// 构造新加的节点
	newNode := &ListNode{Val: val}
	// 初始化cur节点指向虚拟头节点
	cur := mll.dummyHead
	// 遍历到尾节点
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	mll.Size++
}

func (mll *myLinkedList) AddAtIndex(index int, val int) {
	if index < 0 { // 如果索引小于0，索引设置为0
		index = 0
	} else if index > mll.Size { // 索引大于链表长度，直接返回
		return
	}
	// 构造新节点
	newNode := &ListNode{Val: val}
	// 初始化cur节点指向虚拟头节点
	cur := mll.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	// 新节点的下一个节点指向index节点(cur的下一个节点)
	newNode.Next = cur.Next
	// index前一个节点(cur节点)的下一个节点指向新节点
	cur.Next = newNode
	mll.Size++
}

func (mll *myLinkedList) DeleteAtIndex(index int) {
	// 索引无效直接返回
	if index < 0 || index > mll.Size {
		return
	}

	// 初始化cur节点指向虚拟头节点
	cur := mll.dummyHead

	// 遍历至index
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	// index前一个节点(cur节点)的下一个节点指向index节点的下一个节点
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	mll.Size--
}

func (mll *myLinkedList) PrintList() {
	// 初始化cur节点指向虚拟头节点
	cur := mll.dummyHead

	// 遍历链表，打印链表
	for cur.Next != nil {
		fmt.Printf("%d -> ", cur.Next.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}
