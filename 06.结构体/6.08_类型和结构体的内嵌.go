package main

import "fmt"

// 结构体允许其成员字段在声明时没有字段名而只有类型，这种形式的字段被称为类型内嵌或匿名字段类型内嵌
type Data struct {
	int
	string
	bool
	float32
}

// 说明：类型内嵌其实仍然拥有自己的字段名，只是字段名就是其类型本身而已，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个
func GenerateData() *Data {
	return &Data{1, "name", false, 3.1415926}
}

// 结构体实例化后，如果匿名的字段类型为结构体，那么可以直接访问匿名结构体里的所有成员，这种方式被称为结构体内嵌
// 基础颜色
type BasicColor struct {
	R, G, B float32
}

// 完整颜色定义
type Color struct {
	// 将基础颜色作为成员
	Basic BasicColor
	// 透明度
	Alpha float32
}

type Color2 struct {
	BasicColor
	Alpha float32
}

func main() {
	fmt.Println(GenerateData())

	var c Color
	c.Basic.R = 1 // 这种写法太累，冗余了，可以用Color2结构体代替
	c.Basic.G = 1
	c.Basic.B = 0
	c.Alpha = 0.6
	// 显示整个结构体内容
	fmt.Printf("%+v", c)

	fmt.Println("")
	var c2 = new(Color2)
	c2.R = 1
	c2.G = 1
	c2.B = 0
	c2.Alpha = 0.7
	fmt.Printf("%+v", c2)
}

//***********结构内嵌特性*****************
//
// Go语言的结构体内嵌有如下特性:

// 1) 内嵌的结构体可以直接访问其成员变量
// 	  嵌入结构体的成员，可以通过外部结构体的实例直接访问。如果结构体有多层嵌入结构体，结构体实例访问任意一级的嵌入结构体成员时都只用给出字段名，
//    而无须像传统结构体字段一样，通过一层层的结构体字段访问到最终的字段。例如，ins.a.b.c的访问可以简化为ins.c。

// 2) 内嵌结构体的字段名是它的类型名
//    内嵌结构体字段仍然可以使用详细的字段进行一层层访问，内嵌结构体的字段名就是它的类型名
