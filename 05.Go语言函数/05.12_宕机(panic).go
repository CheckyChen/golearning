package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// 1、手动触发宕机
//
// Go 语言可以在程序中手动触发宕机，让程序崩溃，这样开发者可以及时地发现错误，同时减少可能的损失。
//
// Go 语言程序在宕机时，会将堆栈和 goroutine 信息输出到控制台，所以宕机也可以方便地知晓发生错误的位置。
// 如果在编译时加入的调试信息甚至连崩溃现场的变量值、运行状态都可以获取，那么如何触发宕机呢？例如下面的代码：

func panicTest() {
	//panic() 的参数可以是任意类型
	panic("crash")
}

// panic函数的申明
//func panic(v interface{})

// 2、在运行依赖的必备资源缺失时主动触发宕机

// regexp 是 Go 语言的正则表达式包，正则表达式需要编译后才能使用，而且编译必须是成功的，表示正则表达式可用。
//
// 编译正则表达式函数有两种，具体如下：
// 1) func Compile(expr string) (*Regexp, error)
// 编译正则表达式，发生错误时返回编译错误，Regexp 为 nil，该函数适用于在编译错误时获得编译错误进行处理，同时继续后续执行的环境。

// 2) func MustCompile(str string) *Regexp
// 当编译正则表达式发生错误时，使用 panic 触发宕机，该函数适用于直接适用正则表达式而无须处理正则表达式错误的情况。

func MustCompile(str string) *regexp.Regexp {
	// 调用 Compile() 是编译正则表达式的入口函数，该函数返回编译好的正则表达式对象和错误。
	regexp, error := regexp.Compile(str)
	// 如果编译有误，则使用panic()触发宕机
	if error != nil {
		panic(`regexp: Compile(` + strconv.Quote(str) + `):` + error.Error())
	}
	return regexp
}

// 3、在宕机时触发延迟执行语句
// 当 panic() 触发的宕机发生时，panic() 后面的代码将不会被运行，但是在 panic() 函数前面已经运行过的 defer 语句依然会在宕机发生时发生作用，参考下面代码：

func panicDefer() {
	// 宕机前，defer语句会优先被执行，这个特性可以用来在宕机发生前进行宕机信息处理
	defer fmt.Println("宕机后执行1")
	defer fmt.Println("宕机后执行2")
	panic("宕机了")
}
func main() {
	//panicTest()
	panicDefer()
}
