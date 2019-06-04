package main

import "fmt"

// 空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口。从实现的角度看，任何值都满足这个接口的需求。因此空接口类型可以保存任何值，也可以从空接口中取出原值。
// 提示
// 空接口类型类似于 C# 或 Java 语言中的 Object、C语言中的 void*、C++ 中的 std::any。在泛型和模板出现前，空接口是一种非常灵活的数据抽象保存和使用的方法。

// 空接口的内部实现保存了对象的类型和指针。使用空接口保存一个数据的过程会比直接用数据对应类型的变量保存稍慢。因此在开发中，应在需要的地方使用空接口，而不是在所有地方使用空接口。

// 1、将值保存到空接口类型中
func nilInterfaceTest() {
	var any interface{}
	any = 1
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	any = false
	fmt.Println(any)
}

// 2、从空接口获取值
func nilInterfaceTest1() {
	var i int = 1
	var a interface{} = i
	fmt.Println(a)

	// 不能将空接口类型转化为其他类型
	// var b int = a
}

// 3、空接口的值比较
// 空接口在保存不同的值后，可以和其他变量值一样使用==进行比较操作。空接口的比较有以下几种特性

// 1)类型不同的空接口的比较结果不相同
func nilInterfaceEquals() {
	var a interface{} = 100
	var b interface{} = "hi"
	fmt.Println(a == b)
}

// 2)不能比较空接口中的动态值
// 当接口中保存有动态类型的值是，运行时将触发错误：
func nilInterfaceDynamical() {
	var c interface{} = []int{10}
	var d interface{} = []int{20}
	// 会报错：panic: runtime error: comparing uncomparable type []int
	fmt.Println(c == d)
}

// ***************类型的可比较性*******************
// 类  型	 		说  明
// map				宕机错误，不可比较
// 切片（[]T）		宕机错误，不可比较
// 通道（channel）	可比较，必须由同一个 make 生成，也就是同一个通道才会是 true，否则为 false
// 数组（[容量]T）	可比较，编译期知道两个数组是否一致
// 结构体			可比较，可以逐个比较结构体的值
// 函数				可比较
// ***********************************************

func main() {

	nilInterfaceTest()

	nilInterfaceEquals()

	nilInterfaceDynamical()
}
