package main

import (
	"fmt"
	"io"
)

// 在 Go 语言中，不仅结构体与结构体之间可以嵌套，接口与接口间也可以通过嵌套创造出新的接口。
// 接口与接口嵌套组合而成了新接口，只要接口的所有方法被实现，则这个接口中的所有嵌套接口的方法均可以被调用

// 1、系统包中的嵌套组合
type Writer interface {
	Writer(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type WriteCloser interface {
	Writer
	Closer
}

// 2、在代码中使用接口嵌套组合
type device struct{}

func (d *device) Write(p []byte) (n int, err error) {
	fmt.Println("写数据：", p)
	return 0, nil
}

func (d *device) Close() error {
	fmt.Println("关闭")
	return nil
}

func main() {
	// 声明写入关闭器, 并赋予device的实例
	var wc io.WriteCloser = new(device)
	wc.Write(nil)
	wc.Close()

	// 声明写入器, 并赋予device的新实例
	var writeOnly io.Writer = new(device)
	writeOnly.Write(nil)
}
