package acm

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 多组测试样例，每组输⼊数据为字符串，字符⽤空格分隔,输出为⼩数点后两位
func ReadString() {
	//创建⼀个bufio.Reader对象，⽤于从标准输⼊（即键盘）读取数据
	reader := bufio.NewReader(os.Stdin)
	for {
		s, _, err := reader.ReadLine()
		s_list := strings.Split(string(s), " ")
		if err != nil {
			break
		}
		for i := 0; i < len(s_list); i++ {

		}
	}
	fmt.Println(fmt.Sprintf("%.2f", 3.14159))
}
