package _4_流程控制

import (
	"fmt"
)

// 普通写法
func ifelseTest() {
	a := 10
	if a < 100 {
		fmt.Println("a小于100")
	} else if a == 100 {
		fmt.Println("a等于100")
	} else {
		fmt.Println("a大于100")
	}
}

// 特殊写法
func ifelseTest1() {
	// connect()为一个带有返回值的函数，执行connect函数后，将返回值赋值给 err
	// 再判断 err!=""时，进入if分支
	if err := connect(); err != "" {
		fmt.Println(err)
	}
}

func connect() string {
	return "err"
}

// 在编程中，变量在其实现了变量的功能后，作用范围越小，所造成的问题可能性越小，每一个变量代表一个状态，
// 有状态的地方，状态就会被修改，函数的局部变量只会影响一个函数的执行，但全局变量可能会影响所有代码的执行状态，
// 因此限制变量的作用范围对代码的稳定性有很大的帮助。

func main() {
	ifelseTest()
	ifelseTest1()
}
