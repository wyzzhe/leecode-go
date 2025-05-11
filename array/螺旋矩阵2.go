package array

func GenerateMatrix(n int) [][]int {
	// 初始化辅助变量
	startx, starty := 0, 0
	count := 1
	offset := 1
	loop := n / 2
	// 初始化矩阵，外层长度初始化为n，内层长度自动初始化为0
	res := make([][]int, n)
	// 循环初始化内层矩阵长度为n
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	// n/2代表循环几圈
	for loop > 0 {
		i, j := startx, starty
		for ; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}
		for ; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}
	// 对奇数最后一个值特殊处理
	if n%2 != 0 {
		res[n/2][n/2] = n * n
	}
	return res
}
