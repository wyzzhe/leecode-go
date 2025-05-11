package acm

import "fmt"

// 若⼲⾏输⼊，每⾏输⼊两个整数，遇到特定条件终⽌
func APlusBCond() {
	var a, b int
	for {
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			break
		}
		if a == 0 && b == 0 {
			break
		}
		fmt.Println(a + b)
	}
}
