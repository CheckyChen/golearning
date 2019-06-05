package main

import (
	"../08.包/mypkg"
	"../08.包/mypkg2"
	"fmt"
)

// 要引用其他包的标识符，可以使用 import 关键字，导入的包名使用双引号包围，包名是从 GOPATH 开始计算的路径，使用/进行路径分隔

// 1、默认导入的写法
// 1）单行导入
// 单行导入格式如下：

//		import "包1"
//		import "包2"

// 2) 多行导入
// 当多行导入时，包名在 import 中的顺序不影响导入效果，格式如下：

//	import(
//    "包1"
//    "包2"
//    …
//	)

// 2、导入包后自定义引用的包名
// 在默认导入包的基础上，在导入包路径前添加标识符即可形成自定义引用包，格式如下：

// import customName "path/to/package"

// 其中，path/to/package 为要导入的包路径，customName 为自定义的包名

// 3、匿名导入包——只导入包但不使用包内类型和数值
// 如果只希望导入包，而不使用任何包内的结构和类型，也不调用包内的任何函数时，可以使用匿名导入包，格式如下：

// import (
//    _ "path/to/package"
// )

// 其中，path/to/package 表示要导入的包名，下画线_表示匿名导入包。
// 匿名导入的包与其他方式导入包一样会让导入包编译到可执行文件中，同时，导入包也会触发 init() 函数调用

// 4、包在程序启动前的初始化入口：init
// 在某些需求的设计上需要在程序启动时统一调用程序引用到的所有包的初始化函数，如果需要通过开发者手动调用这些初始化函数，
// 那么这个过程可能会发生错误或者遗漏。我们希望在被引用的包内部，由包的编写者获得代码启动的通知，
// 在程序启动时做一些自己包内代码的初始化工作。

// Go 语言为以上问题提供了一个非常方便的特性：init() 函数。
//
// init() 函数的特性如下：
// 每个源码可以使用 1 个 init() 函数。
// init() 函数会在程序执行前（main() 函数执行前）被自动调用。
// 调用顺序为 main() 中引用的包，以深度优先顺序初始化。
//
// 例如，假设有这样的包引用关系：main→A→B→C，那么这些包的 init() 函数调用顺序为：
// C.init→B.init→A.init→main
// 说明：
// 同一个包中的多个 init() 函数的调用顺序不可预期。
// init() 函数不能被其他函数调用。

// 5、理解包导入后的init()函数初始化顺序
func main() {

	fmt.Println("main")
	fmt.Println("from mypkg2 :", mypkg2.ConstVar)
	fmt.Println("from mypkg :", mypkg.MyConst)

	// 代码输出如下：
	//mypkg init function
	//mypkg2 init function
	//main
	//from mypkg2 : hello
	//from mypkg : hello
}
