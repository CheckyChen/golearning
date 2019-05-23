package main

// Go 语言的 defer 语句会将其后面跟随的语句进行延迟处理。在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，
// 也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。
import (
	"fmt"
	"os"
	"sync"
)

func deferTest() {
	fmt.Println("defer begin")
	// 这行会在最后执行
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end defer")
}

// 1、使用延迟执行语句在函数退出是释放资源
// 处理业务或逻辑中涉及成对的操作是一件比较烦琐的事情，比如打开和关闭文件、接收请求和回复请求、加锁和解锁等。
// 在这些操作中，最容易忽略的就是在每个函数退出处正确地释放和关闭资源。
// defer 语句正好是在函数退出时执行的语句，所以使用 defer 能非常方便地处理资源释放问题。

// 1)使用延迟并发解锁
var (
	valueByKey = make(map[string]int)
	// 保证使用映射的并发安全的互锁
	valueByKeyGuard sync.Mutex
)

func readValueNormal(key string) int {
	valueByKeyGuard.Lock()
	v := valueByKey[key]
	valueByKeyGuard.Unlock()
	return v
}

func readValueByDefer(key string) int {
	valueByKeyGuard.Lock()
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

// 2)使用延迟释放文件句柄
func fileSize(filename string) int64 {

	// 根据文件名打开文件
	f, err := os.Open(filename)
	// 如果文件打开错误，返回文件大小为0
	if err != nil {
		return 0
	}
	// 获取文件状态信息
	info, err := f.Stat()
	// 如果获取文件信息错误，关闭文件并返回文件大小为0
	if err != nil {
		f.Close()
		return 0
	}
	// 获取文件
	size := info.Size()
	f.Close()
	return size
}

// 使用defer对代码进行简化
func fileSizeDefer(filename string) int64 {

	f, err := os.Open(filename)

	if err != nil {
		return 0
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		// defer 空中触发，调用关闭文件
		return 0
	}

	size := info.Size()
	// defer 空中触发，调用关闭文件
	return size
}

// defer 后的语句（f.Close()）将会在函数返回前被调用，自动释放资源

func main() {
	//deferTest()

	fmt.Println(fileSizeDefer("test.txt"))
}
