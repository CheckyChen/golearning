package _4_流程控制

import (
	"fmt"
)

// goto 语句通过标签进行代码间的无条件跳转。
// goto 语句可以在快速跳出循环、避免重复退出上有一定的帮助。
// Go 语言中使用 goto 语句能简化一些代码的实现过程。

// 使用goto退出多层循环

func gotoTest() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//使用 goto 语句跳转到指明的标签处
				goto breakThere
			}
		}
	}
	return
breakThere:
	fmt.Println("goto breakThere")
}

func main() {
	gotoTest()
}
