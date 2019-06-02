package main

// Go 语言的类型或结构体没有构造函数的功能。结构体的初始化过程可以使用函数封装实现。

// 1、多种方式创建和初始化结构体——模拟构造函数重载

type Cat struct {
	Color string
	Name  string
}

// 定义用名字构造猫结构的函数，返回Cat指针
func newCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// 定义用毛色构造猫结构的函数，返回Cat指针
func newCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

// 2、带有父子关系的结构体的构造和初始化——模拟父级构造调用
//定义 BlackCat 结构，并嵌入了 Cat 结构体。BlackCat 拥有 Cat 的所有成员，实例化后可以自由访问 Cat 的所有成员
type BlackCat struct {
	Cat // 嵌入Cat, 类似于派生
}

// 构造基类
func newCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// 构造子类
func newBlackCat(color string) *BlackCat {
	cat := &BlackCat{} // 实例化 BlackCat 结构，此时 Cat 也同时被实例化
	cat.Color = color
	return cat
}

// 总之：Go 语言中没有提供构造函数相关的特殊机制，用户根据自己的需求，将参数使用函数传递到结构体构造参数中即可完成构造函数的任务
