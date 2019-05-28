package _3_Go语言容器

import (
	"container/list"
	"fmt"
)

//列表是一种非连续存储的容器，由多个节点组成，节点通过一些变量记录彼此之间的关系。列表有多种实现方法，如单链表、双链表等。
//在 Go 语言中，将列表使用 container/list 包来实现，内部的实现原理是双链表。列表能够高效地进行任意位置的元素插入和删除操作。

//初始化列表
//
//list 的初始化有两种方法：New 和声明。两种方法的初始化效果都是一致的。
//
//1) 通过 container/list 包的 New 方法初始化 list
//变量名 := list.New()
//
//2) 通过声明初始化list
//var 变量名 list.List

//在列表中插入元素
//
//双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。
//提示:
//这两个方法都会返回一个 *list.Element 结构。如果在以后的使用中需要删除插入的元素，则只能通过 *list.Element 配合 Remove() 方法进行删除，这种方法可以让删除更加效率化，也是双链表特性之一。

func listTest() {
	l := list.New()
	l.PushBack("first")
	l.PushFront(67)

	fmt.Println(l)

	//InsertAfter(v interface {}, mark * Element) * Element	在 mark 点之后插入元素，mark 点由其他插入函数提供
	//InsertBefore(v interface {}, mark * Element) *Element	在 mark 点之前插入元素，mark 点由其他插入函数提供
	//PushBackList(other *List)	添加 other 列表元素到尾部
	//PushFrontList(other *List)	添加 other 列表元素到头部
}

//从列表中删除元素
//
//列表的插入函数的返回值会提供一个 *list.Element 结构，这个结构记录着列表元素的值及和其他节点之间的关系等信息。从列表中删除元素时，需要用到这个结构进行快速删除。
func removeList() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 使用
	l.Remove(element)
}

//遍历列表元素
//遍历双链表需要配合 Front() 函数获取头元素，遍历时只要元素不为空就可以继续进行。每一次遍历调用元素的 Next
func eachList() {
	list := list.New()
	// 在末尾添加元素
	list.PushBack("first")
	// 在头部添加元素
	list.PushFront(2)

	for i := list.Front(); i != nil; i.Next() {
		fmt.Println(i.Value)
	}
}
func main() {
	//listTest()
	//removeList()
	eachList()
}
