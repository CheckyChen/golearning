package main

import "time"

// Go 的通道可以在声明时约束其操作方向，如只发送或是只接收。这种被约束方向的通道被称做单向通道

// 1、单向通道的声明格式
//	只能发送的通道类型为chan<-，只能接收的通道类型为<-chan，格式如下：
//	var 通道实例 chan<- 元素类型    // 只能发送通道
//	var 通道实例 <-chan 元素类型    // 只能接收通道
//
//	元素类型：通道包含的元素类型。
//	通道实例：声明的通道变量。

func chSendOnly() {
	ch := make(chan int)
	// 声明一个只能发送数据的通道哦 ，并赋值为ch
	var chSendOnly chan<- int = ch
	// 声明一个只能接受数据的通道，并赋值为ch
	var chRecvOnly <-chan int = ch

	// 只能用于发送数据的通道
	ch2 := make(chan<- int)

	// 只能用于接受数据的通道
	ch3 := make(<-chan string)
}

// 上面的例子中，chSendOnly 只能发送数据，如果尝试接收数据，将会出现如下报错：
//  invalid operation: <-chSendOnly (receive from send-only type chan<- int)
//
// 同理，chRecvOnly 也是不能发送的。
//
// 当然，使用 make 创建通道时，也可以创建一个只发送或只读取的通道：
// ch := make(<-chan int)
// var chReadOnly <-chan int = ch
// <-chReadOnly
// 上面代码编译正常，运行也是正确的。但是，一个不能填充数据（发送）只能读取的通道是毫无意义的。

// 2、time包中的单向通道

// time包中的计时器会返回一个timer实例
func timeTest() {
	timer := time.NewTimer(time.Second)
}

// time的Timer类定义如下：
// type Timer struct {
//	 C <-chan Time
//	 r runtimeTimer
// }

// C 通道的类型就是一种只能接收的单向通道。如果此处不进行通道方向约束，一旦外部向通道发送数据，将会造成其他使用到计时器的地方逻辑产生混乱。
//
// *********因此，单向通道有利于代码接口的严谨性**************
