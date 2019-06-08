package main

import (
	"fmt"
	"reflect"
)

// 反射值对象（reflect.Value）提供对结构体访问的方法，通过这些方法可以完成对结构体任意值的访问

// Field(i int) Value				根据索引，返回索引对应的结构体成员字段的反射值对象。当值不是结构体或索引超界时发生宕机
// NumField() int					返回结构体成员字段数量。当值不是结构体或索引超界时发生宕机
// FieldByName(name string) Value	根据给定字符串返回字符串对应的结构体字段。没有找到时返回零值，当值不是结构体或索引超界时发生宕机
// FieldByIndex(index []int) Value	多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的值。 没有找到时返回零值，当值不是结构体或索引超界时发生宕机
// FieldByNameFunc(match func(string) bool) Value	根据匹配函数匹配需要的字段。找到时返回零值，当值不是结构体或索引超界时发生宕机

type dummy struct {
	a int
	b string

	// 嵌入字段
	float32
	bool

	next *dummy
}

func main() {
	d := reflect.ValueOf(dummy{
		next: &dummy{},
	})

	// 获取成员字段数量
	fmt.Println("NumField", d.NumField())

	// 获取索引为2的字段
	floatField := d.Field(2)
	// 输出类型
	fmt.Println("Field", floatField.Type())

	// 输出字段名为 b 的类型
	fmt.Println("FieldByName(\"b\").Type", d.FieldByName("b").Type())

	// 根据索引查找值中, next字段的int字段的值，
	// 先找到 d 的第四个字段  然后再熊第四个字段中获取 第一个字段的类型
	fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 0}).Type())
}
