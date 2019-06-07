package main

import "fmt"

// 在无缓冲通道的基础上，为通道增加一个有限大小的存储空间形成带缓冲通道。带缓冲通道在发送时无需等待接收方接收即可完成发送过程，
// 并且不会发生阻塞，只有当存储空间满时才会发生阻塞。同理，如果缓冲通道中有数据，接收时将不会发生阻塞，
// 直到通道中没有数据可读时，通道将会再度阻塞

// 1、创建带缓冲的通道

// 格式：通道实例 := make(chan int,缓冲大小)
//	通道类型：和无缓冲通道用法一致，影响通道发送和接收的数据类型。
//	缓冲大小：决定通道最多可以保存的元素数量。
//	通道实例：被创建出的通道实例。

// 2、阻塞条件

// 带缓冲通道在很多特性上和无缓冲通道是类似的。无缓冲通道可以看作是长度永远为 0 的带缓冲通道。因此根据这个特性，
// 带缓冲通道在下面列举的情况下依然会发生阻塞：
// 带缓冲通道被填满时，尝试再次发送数据时发生阻塞。
// 带缓冲通道为空时，尝试接收数据时发生阻塞。

// 为什么Go语言对通道要限制长度而不提供无限长度的通道？

// 我们知道通道（channel）是在两个 goroutine 间通信的桥梁。使用 goroutine 的代码必然有一方提供数据，一方消费数据。
// 当提供数据一方的数据供给速度大于消费方的数据处理速度时，如果通道不限制长度，那么内存将不断膨胀直到应用崩溃。
// 因此，限制通道的长度有利于约束数据提供方的供给速度，供给数据量必须在消费方处理量+通道长度的范围内，才能正常地处理数据。

func main() {
	// 创建一个3个元素缓冲大小的整型通道
	ch := make(chan int, 3)
	fmt.Println(len(ch))

	// 查看当前通道的大小
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(len(ch))
}
