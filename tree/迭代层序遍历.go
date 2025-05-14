package tree

import "container/list"

// 前序遍历
func IterPreOrderTraversal(root *TreeNode) []int {
	// 初始结果数组
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

// 后序遍历
func IterEOrderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	// 初始化双向链表作为栈接收节点
	// 注：双向链表的首节点只作为虚拟节点首尾相连作用，不存储实际值
	st := list.New()
	// 根节点入栈
	st.PushBack(root)

	// 栈不为空时先出栈
	for st.Len() > 0 {
		// 出栈
		node := st.Remove(st.Back()).(*TreeNode)

		// 入结果列表
		res = append(res, node.Val)

		// 左右子树不为空入栈
		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}
	reverse(res)
	return res
}

func reverse(nums []int) []int {
	// 左右指针
	left, right := 0, len(nums)-1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	return nums
}

// 中序遍历
func IterNOrderTraversal(root *TreeNode) []int {
	res := []int{}
	// 双向链表模拟栈
	st := list.New()
	// 指针
	cur := root

	// 遍历树
	for cur != nil || st.Len() > 0 {
		if cur != nil {
			// 入栈
			st.PushBack(cur)
			cur = cur.Left
		} else {
			// 出栈
			node := st.Remove(st.Back()).(*TreeNode)
			// 入列表
			res = append(res, node.Val)
			cur = node.Right
		}
	}
	return res
}
