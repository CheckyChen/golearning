package main

import "fmt"

// 在 Go 语言中，函数也是一种类型，可以和其他类型一样被保存在变量中
func question() {
	fmt.Println("你是谁，你要去哪？")
}

func main() {
	//将变量f声明为 func() 类型，此时 f 就被俗称为“回调函数”。此时 f 的值为 nil
	var q func()
	//将函数question作为值赋给q，此时q的值即为question函数
	q = question
	//调用q函数
	q()
}
