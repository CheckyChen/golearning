package main

import "fmt"

// 通道是一个引用对象，和 map 类似。map 在没有任何外部引用时，Go 程序在运行时（runtime）会自动对内存进行垃圾回收（Garbage Collection, GC）。类似的，通道也可以被垃圾回收，但是通道也可以被主动关闭。
// 格式
// 使用 close() 来关闭一个通道：
//  close(ch)
//
// 关闭的通道依然可以被访问，访问被关闭的通道将会发生一些问题。

// 1、给被关闭的通道发送数据时将会触发panic

func closeChannelTest() {

	ch := make(chan int)

	close(ch)

	// 打印通道的指针, 容量和长度
	fmt.Printf("ptr:%p cap:%d len:%d\n", ch, cap(ch), len(ch))

	// panic: send on closed channel
	ch <- 100
}

// 2、从已关闭的通道接收数据时将不会发生阻塞
//   从已经关闭的通道接收数据或者正在接收数据时，将会接收到通道类型的零值，然后停止阻塞并返回。

func closeChannelTest2() {
	ch := make(chan int, 2)

	ch <- 100
	ch <- 200

	close(ch)

	for i := 0; i < cap(ch)+1; i++ {
		v, ok := <-ch
		fmt.Println(v, ok)
	}
}
func main() {
	//closeChannelTest()
	closeChannelTest2()
}
