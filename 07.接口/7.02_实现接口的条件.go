package main

import "fmt"

// 接口定义后，需要实现接口，调用方才能正确编译通过并使用接口。接口的实现需要遵循两条规则才能让接口可用

// 1、接口被实现的条件一：接口的方法与实现接口的类型方法格式一致

// 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义结构，用于实现DataWriter
type file struct{}

// 实现DataWriter接口的WriteData方法
func (d *file) WriteData(data interface{}) error {
	fmt.Println("WriteData:", data)
	return nil
}

// 1) 函数名不一致导致的报错
// 这样定义本身没有错误，但是如果在 main 函数中，将 *file 类型赋值给 DataWriter 时，
// 需要检查 *file 类型是否完全实现了 DataWriter 接口。
// 显然，编译器因为没有找到 DataWriter 需要的 WriteData() 方法而报错
func (d *file) WriteDataX(data interface{}) error {
	fmt.Println("写数据：", data)
	return nil
}

// 2) 实现接口的方法签名不一致导致的报错
// 这次未实现 DataWriter 的理由变为（错误的 WriteData() 方法类型）发现 WriteData(int)error，期望 WriteData(interface{})error
//func (f *file) WriteData(data int) error {
//	fmt.Println("写数据：", data)
//	return nil
//}

// 2、接口被实现的条件二：接口中所有方法均被实现

// 当一个接口中有多个方法时，只有这些方法都被实现了，接口才能被正确编译并使用。
type DataWriter2 interface {
	WriteData2(data interface{}) error
	IsWrite() bool
}

type file2 struct{}

func (f *file2) WriteData2(data interface{}) error {
	fmt.Println("file2 write data:", data)
	return nil
}

func (f *file2) IsWrite() bool {
	fmt.Println("file2 can write")
	return true
}

func main() {
	f := new(file)
	var writer DataWriter
	writer = f
	writer.WriteData("data")

	f2 := new(file2)
	var writer2 DataWriter2
	writer2 = f2
	writer2.WriteData2("data")
}

// Go 语言的接口实现是隐式的，无须让实现接口的类型写出实现了哪些接口。这个设计被称为非侵入式设计。
// 实现者在编写方法时，无法预测未来哪些方法会变为接口。一旦某个接口创建出来，要求旧的代码来实现这个接口时，就需要修改旧的代码的派生部分，这一般会造成雪崩式的重新编译。
// 提示
//   传统的派生式接口及类关系构建的模式，让类型间拥有强耦合的父子关系。这种关系一般会以“类派生图”的方式进行。经常可以看到大型软件极为复杂的派生树。随着系统的功能不断增加，这棵“派生树”会变得越来越复杂。
//   对于 Go 语言来说，非侵入式设计让实现者的所有类型均是平行的、组合的。如何组合则留到使用者编译时再确认。因此，使用 GO 语言时，不需要同时也不可能有“类派生图”，开发者唯一需要关注的就是“我需要什么？”，以及“我能实现什么？”
