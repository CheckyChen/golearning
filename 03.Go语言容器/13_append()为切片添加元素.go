package _3_Go语言容器

import "fmt"

//Go 语言的内建函数 append() 可以为切片动态添加元素。每个切片会指向一片内存空间，这片空间能容纳一定数量的元素。
//当空间不能容纳足够多的元素时，切片就会进行“扩容”。“扩容”操作往往发生在 append() 函数调用时。

func test1() {
	var numbers []int

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d ,cap: %d ,pointer: %p \n", len(numbers), cap(numbers), numbers)

		//输出发现，len() 函数并不等于 cap。
	}
}

//append() 函数除了添加一个元素外，也可以一次性添加很多元素。
func test2() {
	var a []string
	a = append(a, "a", "b", "c")
	a = append(a, "c", "d", "e")

	team := []string{"f", "g", "h"}
	//在team后面加上了...，表示将 team 整个添加到 car 的后面。
	a = append(a, team...)

	fmt.Println(a)
}

func main() {
	test2()
}
