package main

import (
	"fmt"
	"reflect"
)

// 如果反射值对象（reflect.Value）中值的类型为函数时，可以通过 reflect.Value 调用该函数。
// 使用反射调用函数时，需要将参数使用反射值对象的切片 []reflect.Value 构造后传入 Call() 方法中，
// 调用完成时，函数的返回值通过 []reflect.Value 返回

func add(a, b int) int {
	return a + b
}

func main() {
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	// 构造函数参数, 传入两个整型值
	parameterList := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(12)}
	// 反射调用函数
	result := funcValue.Call(parameterList)
	// 获取第一个返回值, 取整数值
	fmt.Println(result[0].Int())
}

// 反射调用函数的过程需要构造大量的 reflect.Value 和中间变量，对函数参数值进行逐一检查，
// 还需要将调用参数复制到调用函数的参数内存中。调用完毕后，还需要将返回值转换为 reflect.Value，
// 用户还需要从中取出调用值。因此，反射调用函数的性能问题尤为突出，
// 不建议大量使用反射函数调用反射调用函数的过程需要构造大量的 reflect.Value 和中间变量，
// 对函数参数值进行逐一检查，还需要将调用参数复制到调用函数的参数内存中。调用完毕后，
// 还需要将返回值转换为 reflect.Value，用户还需要从中取出调用值。因此，反射调用函数的性能问题尤为突出，
// 不建议大量使用反射函数调用
