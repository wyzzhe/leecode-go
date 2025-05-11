package acm

import "fmt"

// 若⼲⾏输⼊，每⾏包括两个整数a和b，由空格分隔，每⾏输出后接⼀个空⾏。
func APlusBLn() {
	var a, b int
	for {
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			break
		}
		fmt.Printf("%d\n\n", a+b)
	}
}
