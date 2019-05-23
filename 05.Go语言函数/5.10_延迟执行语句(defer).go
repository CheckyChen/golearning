package main

// Go 语言的 defer 语句会将其后面跟随的语句进行延迟处理。在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，
// 也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。
import (
	"fmt"
	"sync"
)

func deferTest(){
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
	valueByKeyGuard  sync.Mutex
)

func readValueNormal(key string)int{
	valueByKeyGuard.Lock()
	v:=valueByKey[key]
	valueByKeyGuard.Unlock()
	return v
}

func readValueByDefer(key string)int{
	valueByKeyGuard.Lock()
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

func main()  {
	deferTest()
}
