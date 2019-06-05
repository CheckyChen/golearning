package main

import (
	"fmt"
	"time"
)

// 在编写 Socket 网络程序时，需要提前准备一个线程池为每一个 Socket 的收发包分配一个线程。
// 开发人员需要在线程数量和 CPU 数量间建立一个对应关系，以保证每个任务能及时地被分配到 CPU 上进行处理，
// 同时避免多个任务频繁地在线程间切换执行而损失效率。
//
// 虽然，线程池为逻辑编写者提供了线程分配的抽象机制。但是，如果面对随时随地可能发生的并发和线程处理需求，
// 线程池就不是非常直观和方便了。能否有一种机制：使用者分配足够多的任务，系统能自动帮助使用者把任务分配到 CPU 上，
// 让这些任务尽量并发运作。这种机制在 Go 语言中被称为 goroutine。
//
// goroutine 的概念类似于线程，但 goroutine 由 Go 程序运行时的调度和管理。
// Go 程序会智能地将 goroutine 中的任务合理地分配给每个 CPU。
//
// Go 程序从 main 包的 main() 函数开始，在程序启动时，Go 程序就会为 main() 函数创建一个默认的 goroutine。

// 1、使用普通函数创建 goroutine

// Go 程序中使用 go 关键字为一个函数创建一个 goroutine。一个函数可以被创建多个 goroutine，一个 goroutine 必定对应一个函数。
// 1) 格式
// 为一个普通函数创建 goroutine 的写法如下：
// 	go 函数名( 参数列表 )
//
// 函数名：要调用的函数名。
// 参数列表：调用函数需要传入的参数。
//
// 使用 go 关键字创建 goroutine 时，被调用函数的返回值会被忽略。
//
// 如果需要在 goroutine 中返回数据，请使用后面介绍的通道（channel）特性，通过通道把数据从 goroutine 中作为返回值传出

func running() {
	var times int
	for {
		times++
		fmt.Println("tick ", times)
		//使用 time.Sleep 暂停 1 秒后继续循环。
		time.Sleep(time.Second)
	}
}

func running2() {
	var times int
	for {
		times = times + 2
		fmt.Println("running2 tick ", times)
		time.Sleep(time.Second)
	}
}

// 2、使用匿名函数创建goroutine

// go 关键字后也可以为匿名函数或闭包启动 goroutine。
// 1) 使用匿名函数创建goroutine的格式
// 使用匿名函数或闭包创建 goroutine 时，除了将函数定义部分写在 go 的后面之外，还需要加上匿名函数的调用参数，格式如下：
// go func( 参数列表 ){
//    函数体
// }( 调用参数列表 )
//
// 其中：
//	参数列表：函数体内的参数变量列表。
//	函数体：匿名函数的代码。
//	调用参数列表：启动 goroutine 时，需要向匿名函数传递的调用参数。

func main() {
	go running()

	go running2()

	// 匿名函数实现go线程，注意自调用()
	go func() {
		var times int
		for {
			times = times + 3
			fmt.Println("noNameFunc tick ", times)
			time.Sleep(time.Second)
		}
	}()

	var input string
	// 接受用户的输入和go线程并行执行
	fmt.Scanln(&input)

}
