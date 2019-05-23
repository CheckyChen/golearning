package main

import "fmt"

// 闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，
// 在闭包中可以继续使用这个自由变量。
// 因此，简单的说：函数 + 引用环境 = 闭包

// 一个函数类型就像结构体一样，可以被实例化。函数本身不存储任何信息，只有与引用环境结合后形成的闭包才具有“记忆性”。
// 函数是编译期静态的概念，而闭包是运行期动态的概念。

// 1、在闭包内部修改引用的变量
func closureTest() {
	// 准备一个字符串
	str := "hello world"

	// 创建一个匿名函数
	foo := func() {
		// 在匿名函数中访问 str
		str = "hello dude"
		fmt.Println("inner", str)
	}

	// 在匿名函数中并没有定义 str，str 的定义在匿名函数之前，此时，str 就被引用到了匿名函数中形成了闭包。
	foo()
	// 执行闭包，此时 str 发生修改，变为 hello dude。
	fmt.Println("outer", str)
}

// 2、闭包的记忆效应
// 被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，
// 闭包本身就如同变量一样拥有了记忆效应。

// 提供一个值，每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

// 3、闭包实现生成器
// 闭包的记忆效应进程被用于实现类似于设计模式中工厂模式的生成器。

// 创建一个玩家生成器
func playerGen(name string) func() (string, int) {
	hp := 100
	return func() (string, int) {
		return name, hp
	}
}

func main() {
	closureTest()

	// 创建一个累加器，初始值为1,返回的accumulator是类型为func()int 的函数变量
	accumulator := Accumulate(1)
	// 累加1，并打印出来
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)

	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", accumulator2)

	// 创建一个玩家生成器
	generator := playerGen("大头儿子")
	// 调用这个玩家生成器，可以获取到玩家的名称和血量
	name, hp := generator()
	fmt.Println(name, hp)
}
