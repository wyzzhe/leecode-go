package acm

import "fmt"

// 多行输入，每行两个整数
func APlusB() {
	var a, b int
	for {
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err != nil {
			break
		}
		fmt.Println(a + b)
	}
}
