package backtracing

var (
	path []int   // 可能出现的某一种组合
	res  [][]int // 所有可能出现的组合
)

// n 多少个数， k 几个数的组合
func Combine(n, k int) [][]int {
	// 初始化数组
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs(n, k, 1)
	return res
}

// n 多少个数， k 几个数的组合，indexStart 每轮递归数组起点标志
func dfs(n, k, start int) {
	// 递归终止条件
	if len(path) == k {
		// 收集元素
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp) // 二维数组先行后列
		return
	}

	// 开始递归
	for i := start; i <= n; i++ {
		if n-i+1 < k-len(path) {
			break
		}
		// 压栈
		path = append(path, i)
		start++
		// 递归
		dfs(n, k, start)
		// 出栈
		path = path[:len(path)-1]
	}
}
