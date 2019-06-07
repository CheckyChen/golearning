package main

import (
	"errors"
	"fmt"
	"time"
)

// 服务器开发中会使用RPC（Remote Procedure Call，远程过程调用）简化进程间通信的过程。RPC 能有效地封装通信过程，让远程的数据收发通信过程看起来就像本地的函数调用一样。
//
// 本例中，使用通道代替 Socket 实现 RPC 的过程。客户端与服务器运行在同一个进程，服务器和客户端在两个 goroutine 中运行。

func RPCClient(ch chan string, req string) (string, error) {
	// 往通道里发送一个请求
	ch <- req
	select {
	// 通道返回数据，并赋值给ack
	case ack := <-ch:
		return ack, nil
	// 请求超时了
	case <-time.After(time.Second):
		return "", errors.New("Time out")
	}
}

func RPCServer(ch chan string) {
	// 无限循环，侦听发送到通道的数据
	for {
		data := <-ch
		fmt.Println("服务端接收到数据:", data)

		// 通过睡眠函数让程序执行阻塞2秒的任务
		//time.Sleep(time.Second * 1)

		ch <- "收到"
	}
}

func main() {
	ch := make(chan string)
	// 开启一个 goroutine 用于侦听通道接收请求
	go RPCServer(ch)
	// 客户端往服务端发送一个请求“你好”
	recv, err := RPCClient(ch, "你好")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("客户端接收到数据:", recv)
	}
}
