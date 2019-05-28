package _4_流程控制

import "fmt"

// break 语句可以结束 for、switch 和 select 的代码块

func breakTest() {
OuterLoop:

	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				// 跳出外层循环，也就是第一层循环外
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}

func main() {
	breakTest()
}
