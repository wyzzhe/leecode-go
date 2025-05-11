package acm

import "fmt"

// 多组测试⽤例，第⼀⾏为正整数n, 第⼆⾏为n个正整数，n=0时，结束输⼊，每组输出结果的下⾯都输出⼀个空⾏
func ReadInt() {
	var n int
	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil || n == 0 {
			break
		}
		nums := make([]int, n)

		for i := 0; i < n; i++ {
			// 读取一个整数，存放在数组中
			fmt.Scanf("%d", &nums[i])
		}
		for i := 0; i < n; i++ {
			fmt.Println(nums[i])
		}
		// fmt.Println(result)
		fmt.Println()
	}
}
