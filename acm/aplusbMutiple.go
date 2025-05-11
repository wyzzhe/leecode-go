package acm

import "fmt"

// 多组数据，每组第一行为n，之后输入n行每行两个整数
func APlusBMutiple() {
	var n, a, b int
	for {
		_, err := fmt.Scan(&n) // 若输入1 2，则n=1，2存入Scan的缓存区
		if err != nil {
			break
		}
		for n > 0 {
			_, err := fmt.Scan(&a, &b) // 若之前输入1 2，此处a=2，2来自Scan的缓存区
			fmt.Println(a, b)
			if err != nil {
				break
			}
			fmt.Println(a + b)
			n--
		}
	}
}
