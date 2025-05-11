package acm

import "fmt"

// 多组测试数据，每组数据只有⼀个整数，对于每组输⼊数据，输出⼀⾏，每组数据下⽅有⼀个空⾏。
func ReadNA() {
	var n, a int
	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil || n == 0 {
			break
		}
		for n != 0 {
			a = n % 10
			n = n / 10
		}
		// fmt.Println(result)
		fmt.Println()
		fmt.Println(a)
	}
}
