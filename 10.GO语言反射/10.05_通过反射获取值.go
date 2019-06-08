package main

import (
	"fmt"
	"reflect"
)

// 反射不仅可以获取值的类型信息，还可以动态地获取或者设置变量的值。Go 语言中使用 reflect.Value 获取和设置变量的值。

// 1、使用反射值对象包装任意值

// Go 语言中，使用 reflect.ValueOf() 函数获得值的反射值对象（reflect.Value）。书写格式如下：
// value := reflect.ValueOf(rawValue)
//
// reflect.ValueOf 返回 reflect.Value 类型，包含有 rawValue 的值信息。
// reflect.Value 与原值间可以通过值包装和值获取互相转化。reflect.Value 是一些反射操作的重要类型，如反射调用函数。

// 2、从反射值对象获取被包装的值

// Go 语言中可以通过 reflect.Value 重新获得原始值。
// 1) 从反射值对象（reflect.Value）中获取值的方法
// 可以通过下面几种方法从反射值对象 reflect.Value 中获取原值
// Interface() interface {}		将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
// Int() int64					将值以 int 类型返回，所有有符号整型均可以此方式返回
// Uint() uint64				将值以 uint 类型返回，所有无符号整型均可以此方式返回
// Float() float64				将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
// Bool() bool					将值以 bool 类型返回
// Bytes() []bytes				将值以字节数组 []bytes 类型返回
// String() string				将值以字符串类型返回

func main() {
	var a int = 1024
	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)

	// 获取interface{}类型的值，通过类型断言转换
	var getA int = valueOfA.Interface().(int)

	// 获取64位的值，强制类型转换为int类型
	var getA2 int = int(valueOfA.Int())
	fmt.Println(getA, getA2)

	var b int = 234
	valueOfB := reflect.ValueOf(b)
	var getB int = valueOfB.Interface().(int)
	var getB2 int = int(valueOfB.Int())
	fmt.Println(getB, getB2)

	c := 1000
	valueOfC := reflect.ValueOf(c)
	getC := valueOfC.Interface().(int)
	getC2 := int(valueOfC.Int())
	fmt.Println(getC, getC2)
}
