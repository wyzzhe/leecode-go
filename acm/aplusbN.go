package acm

import "fmt"

// 若⼲⾏输⼊，遇到0终⽌，每⾏第⼀个数为N，表示本⾏后⾯有N个数
func APlusBN() {
	var n, a int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		if n == 0 {
			break
		}
		sum := 0
		for n > 0 {
			_, err := fmt.Scan(&a)
			if err != nil {
				break
			}
			sum += a
			n--
		}
		fmt.Println(sum)
	}
}
