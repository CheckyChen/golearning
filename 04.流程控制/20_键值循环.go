package _4_流程控制

import "fmt"

//Go 语言可以使用 for range 遍历数组、切片、字符串、map 及通道（channel）。通过 for range 遍历的返回值有一定的规律：
//	①数组、切片、字符串返回索引和值。
//	②map 返回键和值。
//	③通道（channel）只返回通道内的值。

// 遍历数组、切片——获得索引和元素
func keyValueEach1() {
	for key, value := range []int{1, 2, 3, 4, 5, 6} {
		fmt.Println(key, value)
	}
}

// 遍历字符串——获得字符

//可以通过 for range 的组合，对字符串进行遍历，遍历时，key 和 value 分别代表字符串的索引（base0）和字符串中的每一个字符
func strEach() {
	str := "hello 你好"
	for k, v := range str {
		fmt.Printf("key:%d value:0x%x\n", k, v)
	}
}

// 遍历map——获得map的键和值
func mapEach() {
	m := map[string]string{
		"name":   "france",
		"age":    "24",
		"gender": "female",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

// 遍历通道（channel）——接收通道数据
func channelEach() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	//keyValueEach1()
	//strEach()
	//mapEach()
	channelEach()
}
