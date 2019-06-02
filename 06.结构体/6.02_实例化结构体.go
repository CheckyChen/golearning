package main

import "fmt"

// 实例化就是根据结构体定义的格式创建一份与格式一致的内存区域，结构体实例与实例间的内存是完全独立的。
// Go 语言可以通过多种方式实例化结构体，根据实际需要可以选用不同的写法。

// 1、基本的实例化  格式 ：var t T
type Point struct {
	X, Y int
}

// 2、创建指针类型的结构体
// Go 语言中，还可以使用 new 关键字对类型（包括结构体、整型、浮点数、字符串等）进行实例化，结构体在实例化后会形成指针类型的结构体。
// 使用new的格式： ins := new(T)

// 玩家结构体
type Player struct {
	Name        string
	healthPoint int
	MagicPoint  int
}

// 在 Go 语言中，访问结构体指针的成员变量时可以继续使用点（.）。
// 这是因为 Go 语言为了方便开发者访问结构体指针的成员变量，使用了语法糖（Syntactic sugar）技术，
// 将 ins.Name 形式转换为 (*ins).Name

// 3、取结构体的地址实例化
// 在 Go 语言中，对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作。取地址格式如下：
// ins := &T{}
// 其中：T 表示结构体类型，ins 为结构体的实例，类型为 *T，是指针类型。

// 结构体定义一个命令行指令（Command），指令中包含名称、变量关联和注释等
type Command struct {
	Name    string
	Var     *int
	Comment string
}

// 取地址实例化是最广泛的一种结构体实例化方式。可以使用函数封装上面的初始化过程，代码如下：
func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}

func main() {
	var p Point
	p.X = 200
	p.Y = 100

	p1 := new(Point)
	p1.X = 100
	p1.Y = 200

	player := new(Player)
	player.Name = "张三"
	player.healthPoint = 100
	player.MagicPoint = 3000

	var version int = 1
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "版本"

	cmd1 := newCommand("version", &version, "显示版本")
	fmt.Println(cmd1)

}
