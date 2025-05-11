package acm

import "fmt"

// 多组测试数据，每个测试实例包括2个整数M，K（2<=k<=M<=1000)。M=0，K=0代表输⼊结束。
func ReadMK() {
	var m, k int
	for {
		_, err := fmt.Scanf("%d %d", &m, &k)
		if err != nil || (m == 0 && k == 0) {
			break
		}
		sum := m + m/k
		fmt.Println(sum)
	}
}
