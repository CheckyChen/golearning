package main

import "fmt"

// 函数声明以 func 标识，后面紧接着函数名、参数列表、返回参数列表及函数体，具体形式如下：
//		func 函数名(参数列表)(返回参数列表){
// 		   函数体
//		}

// 函数名：由字母、数字、下画线组成。其中，函数名的第一个字母不能为数字。在同一个包内，函数名称不能重名。
//	包（package）是 Go 源码的一种组织方式，一个包可以认为是一个文件夹，在后续章节中将会详细讲解包的概念。

// 参数列表：一个参数由参数变量和参数类型组成，例如： func foo( a int, b string )，其中，参数列表中的变量作为函数的局部变量而存在。

// 返回参数列表：可以是返回值类型列表，也可以是类似参数列表中变量名和类型名的组合。函数在声明有返回值时，
// 必须在函数体中使用 return 语句提供返回值列表。

// 函数体：能够被重复调用的代码片段。

// 1、基本函数
func add(a int, b int) int {
	return a + b
}

// 2、参数类型的简写
// 在参数列表中，如有多个参数变量，则以逗号分隔；如果相邻变量是同类型，则可以将类型省略
func add1(a, b int) int {
	return a + b
}

// 3、函数的返回值

//1) 同一种类型返回值
//如果返回值是同一种类型，则用括号将多个返回值类型括起来，用逗号分隔每个返回值的类型。
//使用 return 语句返回时，值列表的顺序需要与函数声明的返回值类型一致
func typeTwoValue() (int, int) {
	return 1, 2
}

//2) 带有变量名的返回值
//Go 语言支持对返回值进行命名，这样返回值就和参数一样拥有参数变量名和类型。
//命名的返回值变量的默认值为类型的默认值，即数值为 0，字符串为空字符串，布尔为 false、指针为 nil 等。
func namedRetValue() (a, b int) { //对两个整型返回值进行命名，分别为 a 和 b。
	//命名返回值的变量与这个函数的布局变量的效果一致，可以对返回值进行赋值和值获取
	a = 1
	b = 2
	// 当函数使用命名返回值时，可以在 return 中不填写返回值列表，如果填写也是可行的
	// return a,b
	return
}

// *******提示**************
//同一种类型返回值和命名返回值两种形式只能二选一，混用时将会发生编译错误，例如下面的代码：
//func namedRetValues() (a, b int, int)
//编译报错提示：
//mixed named and unnamed function parameters

// 4、调用函数
// 函数在定义后，可以通过调用的方式，让当前代码跳转到被调用的函数中进行执行。调用前的函数局部变量都会被保存起来不会丢失；
// 被调用的函数结束后，恢复到被调用函数的下一行继续执行代码，之前的局部变量也能继续访问。
// 函数内的局部变量只能在函数体中使用，函数调用结束后，这些局部变量都会被释放并且失效。
// Go语言的函数调用格式如下：
// 返回值变量列表 = 函数名(参数列表)

// 5、实战
// 1） 将秒数转化为具体的时间
const (
	SecondsPerMinute = 60
	SecondsPerHour   = 60 * SecondsPerMinute
	SecondsPerDay    = 24 * SecondsPerHour
)

func resolveTime(seconds int) (day, hour, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}
func main() {
	//result := add(1, 2)
	//fmt.Println(result)

	//a, b := typeTwoValue()
	//fmt.Println(a, b)

	//a, b := namedRetValue()
	//fmt.Println(a, b)

	day, hour, minute := resolveTime(3600 * 24)
	fmt.Println(day, hour, minute)
}
