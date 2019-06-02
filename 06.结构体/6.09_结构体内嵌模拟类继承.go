package main

import "fmt"

// Go 语言的结构体内嵌特性就是一种组合特性，使用组合特性可以快速构建对象的不同特性

type Flying struct{}

func (f *Flying) Fly() {
	fmt.Println("飞行咯")
}

type Walkable struct{}

func (w *Walkable) Walk() {
	fmt.Println("行走")
}

type Human struct {
	Walkable // 人类能行走
}

type Bird struct {
	Walkable // 鸟类能行走
	Flying   // 鸟类也能飞行
}

func main() {
	b := new(Bird)
	fmt.Println("Bird:")
	b.Fly()
	b.Walk()

	h := new(Human)
	fmt.Println("Human:")
	h.Walk()
}
