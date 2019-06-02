package main

import "fmt"

// Go 语言可以将类型的方法与普通函数视为一个概念，从而简化方法和函数混合作为回调类型时的复杂性。
// 这个特性和 C# 中的代理（delegate）类似，调用者无须关心谁来支持调用，系统会自动处理是否调用普通函数或类型的方法。
// 本节中，首先将用简单的例子了解 Go 语言是如何将方法与函数视为一个概念，接着会实现一个事件系统，
// 事件系统能有效地将事件触发与响应两端代码解耦。

type class struct{}

func (c *class) Do(v int) {
	fmt.Println("调用了Do函数:", v)
}

func FuncDo(v int) {
	fmt.Println("调用普通的Do函数", v)
}

// 事件系统的基本原理
/**************************

		 ------------
		 * 事件调用者 *
		 ------------
		  ^		  |
		  |		  |
		 注册    调用
		  |	      |
		  |       V
		 -------------
		 *事件相应代码*
		 -------------

****************************/
// 一个事件系统拥有如下特性：
//	①能够实现事件的一方，可以根据事件ID或名字注册对应的事件。
//	②事件发起者，会根据注册信息通知这些注册者。
//	③一个事件可以有多个实现方响应。

// 事件注册
// 事件系统需要为外部提供一个注册入口。这个注册入口传入注册的事件名称和对应事件名称的响应函数，
// 事件注册的过程就是将事件名称和响应函数关联并保存起来，详细实现请参考下面代码的 RegisterEvent() 函数

// 实例化一个通过字符串映射函数切片的map
var eventByName = make(map[string][]func(interface{}))

// 注册事件，提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	// 通过名字查找事件列表
	list := eventByName[name]
	// 在列表切片中添加函数
	list = append(list, callback)
	// 将修改的事件列表切片保存回去
	eventByName[name] = list
}

// 调用事件
func CallEvent(name string, param interface{}) {
	// 通过事件名找到事件列表
	list := eventByName[name]
	// 遍历这个事件的所有回调
	for _, callback := range list {
		callback(param)
	}
}

// 使用事件系统
// 声明角色结构体
type Actor struct{}

func (a *Actor) onEvent(param interface{}) {
	fmt.Println("角色事件", param)
}

func GloabEvent(param interface{}) {
	fmt.Println("全局事件", param)
}
func EventTest() {
	a := new(Actor)
	RegisterEvent("onSkill", a.onEvent)
	RegisterEvent("onSkill", GloabEvent)
	CallEvent("onSkill", 100)
}

func main() {
	// 申明一个函数回调
	var delegate func(int)
	// 创建结构体实例
	c := new(class)
	// 将回调设为c的Do方法
	delegate = c.Do
	// 调用
	delegate(100)

	delegate = FuncDo
	delegate(100)

	EventTest()
}
