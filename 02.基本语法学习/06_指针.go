package _2_基本语法学习

import "fmt"

// 指针（pointer）概念在 Go 语言中被拆分为两个核心概念：
// 	类型指针，允许对这个指针类型的数据进行修改。传递数据使用指针，而无须拷贝数据。类型指针不能进行偏移和运算。
// 	切片，由指向起始元素的原始指针、元素数量和容量组成。

// 受益于这样的约束和拆分，Go 语言的指针类型变量拥有指针的高效访问，但又不会发生指针偏移，
// 从而避免非法修改关键性数据问题。同时，垃圾回收也比较容易对不会发生偏移的指针进行检索和回收

// 1、指针地址和指针类型
// 每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言使用
// & 符号放在变量前面对变量进行“取地址”操作。
// 格式： ptr := &v  // v的类型为T，其中v代表被取地址的变量，被取地址的v使用ptr
// 变量进行接收，ptr的类型就为*T,称作T的指针类型，*代表指针

func pointerTest1() {
	var cat int = 1
	var str string = "banana"
	// 使用 fmt.Printf 的动词%p输出 cat 和 str 变量取地址后的指针值，指针值带有0x的十六进制前缀。
	// 提示：变量、指针和地址三者的关系是：每个变量都拥有地址，指针的值就是地址。
	fmt.Printf("%p %p", &cat, &str)
}

// 2、从指针中获取指针指向的值
// 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
// 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
// 	对变量进行取地址（&）操作，可以获得这个变量的指针变量。
// 	指针变量的值是指针地址。
// 	对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
func pointerTest2() {
	// 准备一个字符串
	var house = "Malibu Point 10880,90256"
	// 对字符串进行取地址操作
	ptr := &house
	// 打印ptr的类型 *string
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的值 十六进制
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值
	value := *ptr
	// 打印value的类型 string
	fmt.Printf("value type: %T\n", value)
	// 打印value的值
	fmt.Printf("value: %s\n", value)
}

// 3、使用指针修改值
func swap(a, b *int) {
	// 取a指针的值，并赋给临时变量t
	t := *a
	// 将b指针的值,赋值给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}

func swap1(a, b *int) {
	a, b = b, a
}

func main() {
	//pointerTest1()
	//pointerTest2()

	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)

	a, b := 3, 4
	swap1(&a, &b)
	fmt.Println(a, b)
}
