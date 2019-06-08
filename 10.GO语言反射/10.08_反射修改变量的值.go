package main

import (
	"fmt"
	"reflect"
)

// 使用 reflect.Value 对包装的值进行修改时，需要遵循一些规则。如果没有按照规则进行代码设计和编写，轻则无法修改对象值，重则程序在运行时会发生宕机

// 1、判定及获取元素的相关方法

// Elem() Value		取值指向的元素值，类似于语言层*操作。当值类型不是指针或接口时发生宕 机，空指针时返回 nil 的 Value
// Addr() Value		对可寻址的值返回其地址，类似于语言层&操作。当值不可寻址时发生宕机
// CanAddr() bool	表示值是否可寻址
// CanSet() bool	返回值能否被修改。要求值可寻址且是导出的字段

// 2、值修改相关方法

// Setlnt(x int64)		使用 int64 设置值。当值的类型不是 int、int8、int16、 int32、int64 时会发生宕机
// SetUint(x uint64)	使用 uint64 设置值。当值的类型不是 uint、uint8、uint16、uint32、uint64 时会发生宕机
// SetFloat(x float64)	使用 float64 设置值。当值的类型不是 float32、float64 时会发生宕机
// SetBool(x bool)		使用 bool 设置值。当值的类型不是 bod 时会发生宕机
// SetBytes(x []byte)	设置字节数组 []bytes值。当值的类型不是 []byte 时会发生宕机
// SetString(x string)	设置字符串值。当值的类型不是 string 时会发生宕机

// *****注意：以上方法，在 reflect.Value 的 CanSet 返回 false 仍然修改值时会发生宕机。************
// 在已知值的类型时，应尽量使用值对应类型的反射设置值。

func setValueByReflect() {
	var a int = 1024
	valueOfA := reflect.ValueOf(&a)
	valueOfA = valueOfA.Elem()

	valueOfA.SetInt(1)
	fmt.Println(valueOfA.Int())
}

func setValueByRelect2() {

	type dog struct {
		LegCount int
	}

	// 获取dog实例地址的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})

	// 取出dog实例地址的元素
	valueOfDog = valueOfDog.Elem()

	valueLegCount := valueOfDog.FieldByName("LegCount")

	valueLegCount.SetInt(4)

	fmt.Println(valueLegCount)
}

// 值的修改从表面意义上叫可寻址，换一种说法就是值必须“可被设置”。那么，想修改变量值，一般的步骤是：
//	①取这个变量的地址或者这个变量所在的结构体已经是指针类型。
//	②使用 reflect.ValueOf 进行值包装。
//	③通过 Value.Elem() 获得指针值指向的元素值对象（Value），因为值对象（Value）内部对象为指针时，使用 set 设置时会报出宕机错误。
//	④使用 Value.Set 设置值。

func main() {
	setValueByReflect()
	setValueByRelect2()
}
