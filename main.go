package main

import (
	"fmt"

	"github.com/wyzzhe/leecode-go/acm"
	"github.com/wyzzhe/leecode-go/array"
	"github.com/wyzzhe/leecode-go/backtracing"
	"github.com/wyzzhe/leecode-go/hashmap"
	"github.com/wyzzhe/leecode-go/list"
	"github.com/wyzzhe/leecode-go/stackandqueue"
	"github.com/wyzzhe/leecode-go/stringLee"
	"github.com/wyzzhe/leecode-go/tree"
)

func main() {
	runArray()
}

func runAcm() {
	flag := "ReadMK"
	switch flag {
	case "APlusB":
		acm.APlusB()
	case "APlusBMutiple":
		acm.APlusBMutiple()
	case "APlusBCond":
		acm.APlusBCond()
	case "APlusBN":
		acm.APlusBN()
	case "APlusBLn":
		acm.APlusBLn()
	case "APlusBNM":
		acm.APlusBNM()
	case "ReadString":
		acm.ReadString()
	case "ReadInt":
		acm.ReadInt()
	case "ReadNA":
		acm.ReadNA()
	case "ReadMK":
		acm.ReadMK()
	}
}

func runArray() {
	flag := "kthLargestNumbers"
	switch flag {
	case "SplitSearch":
		index := array.SplitSearch([]int{1, 2, 3, 4, 5}, 4)
		fmt.Println(index)
	case "RemoveElement":
		index := array.RemoveElementArray([]int{1, 2, 4, 4, 5}, 4)
		fmt.Println(index)
	case "SortedSquares":
		index := array.SortedSquares([]int{-5, -3, 1, 2, 3})
		fmt.Println(index)
	case "MinSubArrayLen":
		index := array.MinSubArrayLen([]int{1, 2, 7, 3, 1}, 10)
		fmt.Println(index)
	case "GenerateMatrix":
		index := array.GenerateMatrix(3)
		fmt.Println(index)
	case "kthLargestNumbers":
		result := array.KthLargestNumbers([]string{"3", "6", "7", "10"}, 3)
		fmt.Println(result)
	}
}

func runList() {
	flag := "DetectCycle"
	head := &list.ListNode{}
	if flag == "DetectCycle" {
		head = &list.ListNode{Val: 1}
		head.Next = &list.ListNode{Val: 2}
		head.Next.Next = &list.ListNode{Val: 3}
		head.Next.Next.Next = &list.ListNode{Val: 4}
		head.Next.Next.Next.Next = head.Next.Next
		printList(head)
	} else {
		head = &list.ListNode{Val: 1}
		head.Next = &list.ListNode{Val: 2}
		head.Next.Next = &list.ListNode{Val: 3}
		head.Next.Next.Next = &list.ListNode{Val: 4}
		head.Next.Next.Next.Next = &list.ListNode{Val: 5}
		printList(head)
	}
	switch flag {
	case "RemoveElementList":
		result := list.RemoveElementList(head, 3)
		printList(result)
	case "DesignList":
		newList := list.Constructor() // 初始化链表
		newList.AddAtHead(2)          // 头插
		newList.AddAtTail(4)          // 尾插
		newList.AddAtIndex(0, 1)      // index插
		newList.AddAtIndex(2, 3)      // index插
		newList.DeleteAtIndex(3)      // 删除
		newList.PrintList()
	case "ReverseList":
		result := list.ReverseList(head)
		printList(result)
	case "ExchangeList":
		result := list.ExchangeList(head)
		printList(result)
	case "RemoveNthFromEnd":
		result := list.RemoveNthFromEnd(head, 1)
		printList(result)
	case "DetectCycle":
		result := list.DetectCycle(head)
		printList(result)
	}
}

func runHashmap() {
	flag := "MinOperations"
	switch flag {
	case "IsAnagram":
		resultTrue := hashmap.IsAnagram("abcd", "cdab")
		resultFalse := hashmap.IsAnagram("abcd", "aecd")
		fmt.Println(resultTrue, resultFalse)
	case "Intersection":
		result := hashmap.IntersectionArray([]int{1, 1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}, []int{3, 3, 7, 9, 12})
		fmt.Println(result)
	case "SumTwo":
		result := hashmap.SumTwo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1)
		fmt.Println(result)
	case "PlusFour":
		nums1 := []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}
		nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		nums3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		nums4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := hashmap.PlusFour(nums1, nums2, nums3, nums4)
		fmt.Println(result)
	case "SumThree":
		result := hashmap.SumThree([]int{1, 3, 5, 6, -9, -5, 0, 2})
		fmt.Println(result)
	case "SumFour":
		result := hashmap.SumFour([]int{1, 3, 5, 6, -9, -5, 0, 2}, 9)
		fmt.Println(result)
	case "MinOperations":
		// 读取输入的n和k
		// var n int
		var k int
		k = 2
		// fmt.Scan(&n)
		// 创建一个长度为n的数组来存储输入的数值
		// nums := make([]int, n)
		nums := []int{2, 3, 1, 3, 4, 3}
		// 读取数组中的数值
		// for i := 0; i < n; i++ {
		// 	fmt.Scan(&nums[i])
		// }
		// fmt.Scan(&k)

		// 调用minOperations函数计算最小操作次数
		result := hashmap.MinOperations(nums, k)
		// 输出结果
		fmt.Println(result)
	}
}

func runString() {
	flag := "ReSubString"
	switch flag {
	case "ReverseString":
		s := []byte{'h', 'e', 'l', 'l', 'o'}
		stringLee.ReverseString(s)
		fmt.Println(string(s))
	case "ReverseWords":
		s := "   hello    world   "
		s = stringLee.ReverseWords(s)
		fmt.Println(s)
	case "KMP":
		haystack := "aabaabaaf"
		needle := "aabaaf"
		index := stringLee.KMP(haystack, needle)
		fmt.Println(index)
	case "ReSubString":
		s := "abababab"
		isReSubStringKMP := stringLee.RepeatedSubStringKMP(s)
		isReSubString := stringLee.RepeatedSubString(s)
		fmt.Println(isReSubStringKMP, isReSubString)
	}

}

func runStackAndQueue() {
	flag := "TopKFrequent"
	switch flag {
	case "stacktoqueue":
		queue := stackandqueue.QueueConstructor()
		queue.Push(1)
		queue.Push(2)
		queue.Push(3)
		queue.Pop()
		fmt.Printf("%+v\n", queue.Peek())
	case "queuetostack":
		stack := stackandqueue.StackConstructor()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Pop()
		fmt.Printf("%+v\n", stack)
	case "IsValid":
		isValid1 := stackandqueue.IsValid("{()}") // 正确
		isValid2 := stackandqueue.IsValid("()}")  // 多余右括号
		isValid3 := stackandqueue.IsValid("(()")  // 多余左括号
		isValid4 := stackandqueue.IsValid("(()}") // 不匹配括号
		fmt.Println(isValid1, isValid2, isValid3, isValid4)
	case "RemoveDuplicates":
		s := stackandqueue.RemoveDuplicates("accab")
		fmt.Println(s)
	case "EvalRPN":
		s := stackandqueue.EvalRPN([]string{"1", "2", "+", "3", "4", "+", "*"})
		fmt.Println(s)
	case "MaxSlidingWindow":
		s := stackandqueue.MaxSlidingWindow([]int{1, -1, 2, 4, 9, 5}, 3)
		fmt.Println(s)
	case "GetSeaSon":
		stackandqueue.GetSeaSon()
	case "TopKFrequent":
		result := stackandqueue.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
		fmt.Println(result)
	}
}

func runTree() {
	root := getTree()
	flag := "InvertTree"
	switch flag {
	case "RecursionTraversal":
		// 前序遍历
		res := tree.RecursionPreOrderTraversal(root)
		fmt.Printf("%#v\n", res)
		// 中序遍历
		res = tree.RecursionNTraversal(root)
		fmt.Printf("%#v\n", res)
		// 后序遍历
		res = tree.RecursionETraversal(root)
		fmt.Printf("%#v\n", res)
	case "IterTraversal":
		result := tree.IterPreOrderTraversal(root)
		fmt.Println(result)
		result = tree.IterNOrderTraversal(root)
		fmt.Println(result)
		result = tree.IterEOrderTraversal(root)
		fmt.Println(result)
	case "InvertTree":
		result := tree.InvertTree(root)
		fmt.Println("处理后的二叉树为")
		printTree(result)
	}
}

func runBackTracing() {
	flag := "Combine"
	switch flag {
	case "Combine":
		result := backtracing.Combine(4, 2)
		fmt.Println(result)
	}
}

// 辅助函数：打印链表
func printList(head *list.ListNode) {
	result := list.DetectCycle(head)
	flag := false
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		if head.Val == result.Val {
			// 第一次打印到环入口
			if !flag {
				flag = true
			} else {
				// 第二次打印到环入口
				break
			}
		}
		head = head.Next
	}
	fmt.Println("nil")
}

// 辅助函数：打印二叉树
func printTree(root *tree.TreeNode) {
	if root == nil {
		fmt.Println("二叉树为空")
		return
	}

	node := root
	fmt.Println(node.Val)
	printTree(node.Left)
	printTree(node.Right)
}

// 辅助函数：输入或构造二叉树
func getTree() *tree.TreeNode {
	acmFlag := false
	var root *tree.TreeNode
	if acmFlag {
		// acm模式二叉树输入
		// 输入第一行是n，代表有n行数据（n个节点）
		var n int
		fmt.Scan(&n)

		// nodes节点数组从下标1开始存放二叉树节点，所以数组容量为n+1
		nodes := make([]*tree.TreeNode, n+1)

		// 先初始化所有节点，下标从1到n（二叉树节点数），nodes[0]已经被初始化为0
		for i := 1; i <= n; i++ {
			nodes[i] = &tree.TreeNode{}
		}

		// 接收二叉树节点的值和左右指针下标
		var val int
		for i := 1; i <= n; i++ {
			fmt.Scan(&val)
			left, right := 0, 0
			fmt.Scan(&left, &right)
			// 二叉树节点单独对值赋值
			nodes[i].Val = val
			// 左右指针下标不为0时，左右子树节点指针赋值
			if left != 0 {
				nodes[i].Left = nodes[left]
			}
			if right != 0 {
				nodes[i].Right = nodes[right]
			}
		}
		root = nodes[1]
	} else {
		// 自动构建一个5个节点的二叉树
		n := 5
		nodes := make([]*tree.TreeNode, n+1)
		for i := 1; i <= n; i++ {
			nodes[i] = &tree.TreeNode{}
		}
		for i := 1; i <= n; i++ {
			val := i
			left, right := 2*i, 2*i+1
			nodes[i].Val = val
			if left <= n && left != 0 {
				nodes[i].Left = nodes[left]
			}
			if right <= n && right != 0 {
				nodes[i].Right = nodes[right]
			}
		}
		root = nodes[1]
		fmt.Println("自动构建的二叉树为")
		printTree(root)
	}
	return root
}
