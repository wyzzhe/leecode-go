package stackandqueue

import "fmt"

func switchSeaSon(month int) string {
	switch {
	case month >= 3 && month <= 5:
		return "春季"
	case month >= 6 && month <= 8:
		return "夏季"
	case month >= 9 && month <= 11:
		return "秋季"
	case month >= 12 && month <= 2:
		return "冬季"
	default:
		return "未知季节"
	}
}

func GetSeaSon() {
	var year, month int

	// 输入年月
	fmt.Print("请输入年份")
	fmt.Scan(&year)
	fmt.Print("请输入月份")
	fmt.Scan(&month)

	// 检查月份是否有效
	if month < 1 || month > 12 {
		fmt.Println("无效的月份，请输入1到12之间的数字")
		return
	}

	season := switchSeaSon(month)
	fmt.Printf("当前是%d年%d月，%s\n", year, month, season)
}
