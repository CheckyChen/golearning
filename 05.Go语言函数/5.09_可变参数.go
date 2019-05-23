package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

// 所谓可变参数，是指参数数量不固定的函数形式。
// Go 语言支持可变参数特性，函数声明和调用时没有固定数量的参数，同时也提供了一套方法进行可变参数的多级传递。

// 格式：
// func 函数名(固定参数列表,v ...T){
// 		函数体
// }
//	①可变参数一般被放置在函数列表的末尾，前面是固定参数列表，当没有固定参数时，所有变量就将是可变参数。
//	②v 为可变参数变量，类型为 []T，也就是拥有多个 T 元素的 T 类型切片，v 和 T 之间由...即3个点组成。
//	③T 为可变参数的类型，当 T 为 interface{} 时，传入的可以是任意类型。

// 1、所有的参数都是可变参数：比如，fmt.Println()
func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a)
}

// 2、部分参数是可变参数：比如，fmt.Printf()
// fmt.Printf 的第一个参数为参数列表，后面的参数是可变参数
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, format, a...)
}

func testFunc(a ...string) {
	// a的类型为切片
	fmt.Println(a)
}

// 打印每个参数的类型和值
func printTypeValue(slist ...interface{}) string {
	var b bytes.Buffer
	for _, s := range slist {

		//使用 fmt.Sprintf 配合%v动词，可以将 interface{} 格式的任意值转为字符串
		str := fmt.Sprintf("%v", s)

		var typeString string

		// switch s.(type) 可以对 interface{} 类型进行类型断言，也就是判断变量的实际类型
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}

		// 拼接结果
		b.WriteString("value:")
		b.WriteString(str)
		b.WriteString(" type:")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}

// 3、在多个可变参数函数中传递参数

// 可变参数变量是一个包含所有参数的切片，如果要在多个可变参数中传递参数，可以在传递时在可变参数变量中默认添加...，
// 将切片中的元素进行传递，而不是传递可变参数变量本身
func rawPrint(rawList ...interface{}) {
	for _, a := range rawList {
		fmt.Println(a)
	}
}

func print(slist ...interface{}) {
	rawPrint(slist...)
	// 如果不加... 的话，那么传递的就是slist本身，类型为切片
	rawPrint(slist)
}

func main() {
	// 传入的值类型不受限制
	Println(5, "hello", &struct{ a int }{1}, true)

	// fmt.Printf() 函数在调用时，第一个函数始终必须传入字符串，对应参数是 format，后面的参数数量可以变化
	Printf("pure string\n")
	Printf("value: %v %f\n", true, math.Pi)

	testFunc("a", "b", "333")

	fmt.Println(printTypeValue("a", true, 33))

	print(1, 2, 3, 4)
	print(1)
}
