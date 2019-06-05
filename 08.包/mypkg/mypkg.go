package mypkg

import "fmt"

var myVar = 100

const myConst = "hello"

type myStruct struct{}

// 以上标识符的首字母都是小写，这些标识符可以在包内使用，包外时无法访问他们的

var MyVar = 100

const MyConst = "hello"

type MyStruct struct{}

// 以上标识符的首字母都是大写，这些标识符可以在包外使用，包外时可以直接访问他们

func init() {
	fmt.Println("mypkg init function")
}
