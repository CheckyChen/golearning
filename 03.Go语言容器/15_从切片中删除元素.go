package _3_Go语言容器

import "fmt"

//Go 语言并没有对删除切片元素提供专用的语法或者接口，需要使用切片本身的特性来删除元素

func test() {
	//初始化一个切片
	var arr = []string{"a", "b", "c", "d"}
	//声明要删除的位置
	var seqPosition int = 1
	//查看删除前后的元素
	fmt.Println(arr[:seqPosition], arr[seqPosition+1:])
	arr = append(arr[:seqPosition], arr[seqPosition+1:]...)
	fmt.Println(arr)
}

//Go 语言中切片元素的删除过程并没有提供任何的语法糖或者方法封装，无论是初学者学习，还是实际使用都是极为麻烦的。
//
//连续容器的元素删除无论是在任何语言中，都要将删除点前后的元素移动到新的位置。随着元素的增加，这个过程将会变得极为耗时。
//因此，当业务需要大量、频繁地从一个切片中删除元素时，如果对性能要求较高，就需要反思是否需要更换其他的容器（如双链表等能快速从删除点删除元素）。

func main() {
	test()
}
