package main

import (
	"fmt"
	"reflect"
)

// IsNil() bool		返回值是否为 nil。如果值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic，类似于语言层的v== nil操作
// IsValid() bool	判断值是否有效。 当值本身非法时，返回 false，例如 reflect Value不包含任何值，值为 nil 等。

func main() {
	var a *int
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())

	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())

	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())

	s := struct{}{}
	// 判断 s 中是否有某个字段
	fmt.Println("不存在的结构体成员：", reflect.ValueOf(s).FieldByName("").IsValid())
	// 判断 s 中是否存在某个方法
	fmt.Println("不存在的结构体方法：", reflect.ValueOf(s).MethodByName("").IsValid())

	m := map[int]int{}

	// 判断 m 中索引为3的值是否合法
	fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
}
