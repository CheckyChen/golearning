package main

import "fmt"

type A struct {
	a int // 定义一个a变量
}

type B struct {
	a int // 也定义了一个a变量
}

type C struct {
	A
	B
}

func main() {
	c := new(C)

	//这样能会报错:ambiguous selector c.a
	//c.a = 1

	// 这样正常
	c.A.a = 1
	fmt.Printf("%+v\n", c)
}
