package main

import (
	"flag"
	"fmt"
)

// Go 语言支持匿名函数，即在需要使用函数时，再定义函数，匿名函数没有函数名，只有函数体，函数可以被作为一种类型被赋值给函数类型的变量，匿名函数也往往以变量方式被传递。
//
//匿名函数经常被用于实现回调函数、闭包等。

// 1、定义一个匿名函数
// 匿名函数的定义格式如下：
// func(参数列表)(返回参数列表){
//     函数体
// }
// ******* 匿名函数的定义就是没有名字的普通函数定义************

//func(data int){
//	fmt.Println("hello ",data)
//}

// 注意第“}”后的(100)，表示对匿名函数进行调用，传递参数为 100。

// 2、将匿名函数赋值给变量
func noNameTest() {
	f := func(data int) {
		fmt.Println("你好，", data)
	}
	f(100)
}

// 3、匿名函数作为回调函数
// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

// 4、使用匿名函数实现操作封装
var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	//noNameTest()

	// 使用匿名函数打印切片内容
	//visit([]int{1, 2, 3, 4}, func(v int) {
	//	fmt.Println(v)
	//})

	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}

	//第 42 行，定义命令行参数 skill，从命令行输入— skill 可以将空格后的字符串传入 skillParam 指针变量。
	//第 12 行，解析命令行参数，解析完成后，skillParam 指针变量将指向命令行传入的值。
	//第 14 行，定义一个从字符串映射到 func()的map，然后填充这个 map。
	//第 15～23 行，初始化 map 的键值对，值为匿名函数。
	//第 26 行，skillParam 是一个 *string 类型的指针变量，使用 *skillParam 获取到命令行传过来的值，并在 map 中查找对应命令行参数指定的字符串的函数。
	//第 29 行，如果在 map 定义中存在这个参数就调用；否则打印“技能没有找到”。
}
