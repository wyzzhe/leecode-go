package acm

import "fmt"

// 多组n⾏数据，每⾏先输⼊⼀个整数N，然后在同⼀⾏内输⼊M个整数,每组输出之间输出⼀个空⾏。
func APlusBNM() {
	var N, M, num, sum int
	for {
		_, err := fmt.Scan(&N)
		if err != nil {
			break
		}

		for i := 0; i < N; i++ {
			sum = 0
			fmt.Scan(&M)

			for j := 0; j < M; j++ {
				fmt.Scan(&num)
				sum += num
			}

			// 输出当前组的数字和
			fmt.Println(sum)

			// 如果不是当前组的最后一个数字，输出空行
			if i < N-1 {
				fmt.Println()
			}
		}
	}
}
