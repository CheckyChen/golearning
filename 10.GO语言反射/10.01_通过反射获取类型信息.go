package main

import (
	"fmt"
	"reflect"
)

// 在 Go 程序中，使用 reflect.TypeOf() 函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息

func reflactTest() {
	var a int
	typeOfA := reflect.TypeOf(a)

	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}

// 理解反射的类型（Type）与种类（Kind）

// 在使用反射时，需要首先理解类型（Type）和种类（Kind）的区别。编程中，使用最多的是类型，但在反射中，
// 当需要区分一个大品种的类型时，就会用到种类（Kind）。例如，需要统一判断类型中的指针时，使用种类（Kind）信息就较为方便。
// 1) 反射种类（Kind）的定义
// Go 程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，
// 以及使用 type 关键字定义的类型，这些类型的名称就是其类型本身的名称。例如使用 type A struct{} 定义结构体时，
// A 就是 struct{} 的类型。
//
// 种类（Kind）指的是对象归属的品种，在 reflect 包中有如下定义：

//type Kind uint
//
//const (
//	Invalid       Kind = iota // 非法类型
//	Bool                      // 布尔型
//	Int                       // 有符号整型
//	Int8                      // 有符号8位整型
//	Int16                     // 有符号16位整型
//	Int32                     // 有符号32位整型
//	Int64                     // 有符号64位整型
//	Uint                      // 无符号整型
//	Uint8                     // 无符号8位整型
//	Uint16                    // 无符号16位整型
//	Uint32                    // 无符号32位整型
//	Uint64                    // 无符号64位整型
//	Uintptr                   // 指针
//	Float32                   // 单精度浮点数
//	Float64                   // 双精度浮点数
//	Complex64                 // 64位复数类型
//	Complex128                // 128位复数类型
//	Array                     // 数组
//	Chan                      // 通道
//	Func                      // 函数
//	Interface                 // 接口
//	Map                       // 映射
//	Ptr                       // 指针
//	Slice                     // 切片
//	String                    // 字符串
//	Struct                    // 结构体
//	UnsafePointer             // 底层指针
//)

// Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。
//
// type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。

// 2) 从类型对象中获取类型名称和种类的例子
// Go 语言中的类型名称对应的反射获取方法是 reflect.Type 中的 Name() 方法，返回表示类型名称的字符串。
//
// 类型归属的种类（Kind）使用的是 reflect.Type 中的 Kind() 方法，返回 reflect.Kind 类型的常量。
//
// 下面的代码中会对常量和结构体进行类型信息获取。

type Enum int

const Zero Enum = 0

func reflectTest2() {
	type cat struct{}
	typeOfCat := reflect.TypeOf(cat{})

	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	typeOfA := reflect.TypeOf(Zero)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}

func main() {
	reflactTest()
	reflectTest2()
}
