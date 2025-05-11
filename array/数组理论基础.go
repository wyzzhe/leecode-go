package array

import "fmt"

// go中的数组
//
// 数组在内存中连续存储且数据类型相同，长度在声明时确定且不可改变
// 通过下标获取数据
// 数组的元素不能删除只能覆盖
// 数组是值类型，会复制整个数组，想用函数修改原数组得传递数组的指针

// 移动数组位置相较困难

// 声明包含5个int的数组
var arr1Int [5]int           // 数组会初始化为[0 0 0 0 0]
var arr1String [5]string     // 数组会初始化为[    ]，实际为五个""空字符串
var arr1PtrString [5]*string // 数组会初始化为[<nil> <nil> <nil> <nil> <nil>]

// 声明并初始化
var arr2 = [2]int{1, 2}

// 短声明
func initArray() [3]int {
	arr3 := [3]int{1, 2, 3}
	return arr3
}

// 数组的结构体近似表示（实际Go运行时用Go实现）
// struct {
//     len  int       // 编译期已知，运行时不需要存储
//     data *ElemType // 指向连续内存块的指针
// }

// a = b  // 编译错误：类型不匹配，数组的长度是数组类型的一部分
var a [3]int
var b [5]int

func arrayBase() {
	// 自动推断数组长度
	arr4 := [...]int{1, 2, 3, 4}
	fmt.Println(arr4)

	// 只初始化特定位置的元素
	arr5 := [5]int{1: 10, 3: 30}
	// 结果: [0, 10, 0, 30, 0]
	fmt.Println(arr5)

	arr6 := [5]int{} // 数组会初始化为[0 0 0 0 0]

	fmt.Println(arr1Int)
	fmt.Println(arr1String)
	fmt.Println(arr1PtrString)
	fmt.Println(arr2)
	fmt.Println(arr6)
	fmt.Println(a)
	fmt.Println(b)

	arr3 := initArray()
	fmt.Println(arr3)
}
