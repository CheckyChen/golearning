package main

import "fmt"

// 结构体内嵌初始化时，将结构体内嵌的类型作为字段名像普通结构体一样进行初始化

// 轮子
type Wheel struct {
	Size int
}

// 引擎
type Engine struct {
	Power int    //功率
	Type  string //类型
}

// 车
type Car struct {
	Wheel
	Engine
}

// 初始化内嵌匿名结构体
//
// 在前面描述车辆和引擎的例子中，有时考虑编写代码的便利性，会将结构体直接定义在嵌入的结构体中。
// 也就是说，结构体的定义不会被外部引用到。在初始化这个被嵌入的结构体时，就需要再次声明结构才能赋予数据。具体请参考下面的代码。

type Car2 struct {
	Wheel // 结构体定义时，这个内嵌的结构体后面不能带逗号
	Engin struct {
		Power int
		Type  string
	}
}

func main() {
	c := Car{
		Wheel: Wheel{
			Size: 18,
		},
		Engine: Engine{
			Type:  "1.4T",
			Power: 143,
		},
	}

	fmt.Printf("%+v\n", c)

	c2 := Car2{
		Wheel: Wheel{
			Size: 14,
		},
		Engin: struct {
			Power int
			Type  string
		}{Power: 143, Type: "1.4T"},
	}

	fmt.Printf("%+v\n", c2)
}
