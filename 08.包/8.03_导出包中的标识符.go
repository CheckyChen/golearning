package main

// 在 Go 语言中，如果想在一个包里引用另外一个包里的标识符（如类型、变量、常量等）时，必须首先将被引用的标识符导出，
// 将要导出的标识符的【首字母大写】就可以让引用者可以访问这些标识符了。

// 1、导出包内标识符
var myVar = 100

const myConst = "hello"

type myStruct struct{}

// 以上标识符的首字母都是小写，这些标识符可以在包内使用，包外时无法访问他们的

var MyVar = 100

const MyConst = "hello"

type MyStruct struct{}

// 以上标识符的首字母都是大写，这些标识符可以在包外使用，包外时可以直接访问他们

// 2、导出结构体及接口成员
type MyStruct1 struct {
	// 包外可以访问的字段
	ExportedField int
	// 仅限包内访问的字段
	privateField int
}
type MyInterface interface {
	// 包外可以访问的方法
	ExportedMethod()
	// 仅限包内访问的方法
	privateMethod()
}

// 在代码中，MyStruct 的 ExportedField 和 MyInterface 的 ExportedMethod() 可以被包外访问。
