package _4_流程控制

import "fmt"

// for 循环可以通过 break、goto、return、panic 语句强制退出循环。for 的初始语句、条件表达式、结束语句的详细介绍如下。

// for 中的初始语句——开始循环时执行的语句
//
//初始语句是在第一次循环前执行的语句，一般使用初始语句执行变量初始化，如果变量在此处被声明，其作用域将被局限在这个 for 的范畴内。

func forTest() {
	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}
}

// for 中的条件表达式——控制是否循环的开关
//
//对每次循环开始前计算的表达式，如果表达式为 true，则循环继续，否则结束循环。条件表达式可以被忽略，被忽略条件的表达式默认形成无限循环。

// 1) 结束循环时带可执行语句的无限循环
func forTest1() {
	var i int
	for ; ; i++ {
		fmt.Println(i)
		if i > 10 {
			break
		}
	}
}

// 2) 无限循环
func forTest2() {
	var i int
	for {
		fmt.Println(i)
		if i > 5 {
			break
		}
		i++
	}
}

// 3) 只有一个循环条件的循环
func forTest3() {
	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

// 生成九九乘法表
func demo() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" %d*%d=%d", j, i, i*j)
		}
		fmt.Println()
	}
}

func main() {
	//forTest()
	//forTest1()
	//forTest2()
	//forTest3()
	demo()
}
