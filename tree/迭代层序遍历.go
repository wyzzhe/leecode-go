package tree

import "container/list"

// 前序遍历
func IterPreOrder(root *TreeNode) []int {
	// 初始化栈和结果数组
	res := []int{}

	if root == nil {
		return res
	}

	// 双向链表模拟栈
	st := list.New()
	// 入栈（链表尾部）
	st.PushBack(root)

	// 栈不为空时循环出栈
	for st.Len() > 0 {
		// 节点出栈（链表尾部），node每次都指向出栈的那个节点
		node := st.Remove(st.Back()).(*TreeNode)

		// 写入结果数组
		res = append(res, node.Val)

		// 右、左子树入栈
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}

	return res
}
