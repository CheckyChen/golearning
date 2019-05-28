package _4_流程控制

import "fmt"

// 在 Go 语言中的 switch，不仅可以基于常量进行判断，还可以基于表达式进行判断

// 基本写法
//
// Go 语言改进了 switch 的语法设计，避免人为造成失误。Go 语言的 switch 中的每一个 case 与 case 间是独立的代码块，
// 不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行

func switchTest() {
	var a string = "world1"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		// 一个switch只能有一个default
		fmt.Println(3)
	}
}

// 1) 一分支多值

func switchTest1() {
	var a string = "mom"
	switch a {
	case "mom", "daddy":
		fmt.Println("family")
	}
}

// 2) 分支表达式
func switchTest2() {
	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}
}

// 跨越case的fallthrough,兼容C语言的case设计
func switchTest3() {
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}

func main() {
	//switchTest()
	//switchTest1()
	//switchTest2()
	switchTest3()

}
