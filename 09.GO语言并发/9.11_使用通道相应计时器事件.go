package main

import (
	"fmt"
	"time"
)

// Go 语言中的 time 包提供了计时器的封装。
// 由于 Go 语言中的通道和 goroutine 的设计，定时任务可以在 goroutine 中通过同步的方式完成，
// 也可以通过在 goroutine 中异步回调完成。这里将分两种用法进行例子展示。

// 定点计时
// 计时器（Timer）的原理和倒计时闹钟类似，都是给定多少时间后触发。
// 打点器（Ticker）的原理和钟表类似，钟表每到整点就会触发。这两种方法创建后会返回 time.Ticker 对象和 time.Timer 对象，
// 里面通过一个 C 成员，类型是只能接收的时间通道（<-chan Time），使用这个通道就可以获得时间触发的通知。
//
// 下面代码创建一个打点器，每 500 毫秒触发一起；创建一个计时器，2 秒后触发，只触发一次。

func Timer() {
	// 创建一个打点器，每500毫秒触发一次
	timer := time.NewTicker(time.Microsecond * 500)
	// 常见一个计时器，2秒后触发
	stopper := time.NewTimer(time.Second * 2)

	var i int
	for {
		select {

		case <-stopper.C: //计时器到了
			fmt.Println("停止")
			goto StopLab

		case <-timer.C: //打点器到了
			i++
			fmt.Println("打点器：", i)
		}
	}

StopLab:
	fmt.Println("完成")

}
func main() {
	exit := make(chan int)
	fmt.Println("开始")

	// 过1秒后，调用匿名函数
	// time.AfterFunc() 函数是在 time.After 基础上增加了到时的回调，方便使用
	time.AfterFunc(time.Second, func() {
		fmt.Println("1秒钟之后")
		exit <- 0
	})
	// 等待结束
	data := <-exit
	fmt.Println(data)

	Timer()
}
