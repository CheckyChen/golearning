package main

import "fmt"

// Go 语言中使用接口断言（type assertions）将接口转换成另外一个接口，也可以将接口转换为另外的类型。接口的转换在开发中非常常见，使用也非常频繁。

// 1、类型断言的格式
// t:= i.(T)
// 其中，i 代表接口变量，T 代表转换的目标类型，t 代表转换后的变量

// 如果 i 没有完全实现 T 接口的方法，这个语句将会触发宕机。触发宕机不是很友好，因此上面的语句还有一种写法：
// t,ok := i.(T)
// 这种写法下，如果发生接口未实现时，将会把 ok 置为 false，t 置为 T 类型的 0 值。正常实现时，ok 为 true。这里 ok 可以被认为是：i 接口是否实现 T 类型的结果。

// 2、将接口转换为其他接口
// 实现某个接口的类型同时实现了另外一个接口，此时可以在两个接口间转换

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct{}

func (b *bird) Fly() {
	fmt.Println("鸟在飞")
}

func (b *bird) Walk() {
	fmt.Println("鸟在走")
}

type pig struct{}

func (p *pig) Walk() {
	fmt.Println("猪在走")
}

// 3、将接口转换为其他类型
// 可以实现将接口转换为普通的指针类型
func interfaceTest() {
	p1 := new(pig)
	var a Walker = p1

	p2 := a.(*pig)
	// 连个结果的值都是一样的
	fmt.Printf("p1=%p  p2=%p", p1, p2)
}

// 总结
//
// 接口和其他类型的转换可以在 Go 语言中自由进行，前提是已经完全实现。
// 接口断言类似于流程控制中的 if。但大量类型断言出现时，应使用更为高效的类型分支 switch 特性。

func main() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	for name, obj := range animals {
		f, isFlyer := obj.(Flyer)
		w, isWalk := obj.(Walker)

		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalk)

		if isFlyer {
			f.Fly()
		}

		if isWalk {
			w.Walk()
		}
	}

	interfaceTest()
}
