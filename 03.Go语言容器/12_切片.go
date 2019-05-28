package _3_Go语言容器

import "fmt"

// 切片（Slice）是一个拥有相同类型元素的可变长度的序列。
// Go 语言切片的内部结构包含地址、大小和容量。切片一般用于快速地操作一块数据集合

// 1、从数组或切片生成新的切片
// 切片默认指向一段连续内存区域，可以是数组，也可以是切片本身。
// 格式：slice[开始位置:结束位置]

func test() {
	var a = [...]int{1, 2, 3, 4}
	//取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取。
	fmt.Println(a[1:3])
}

func test2() {
	var arr [30]int
	for i := 0; i < 30; i++ {
		arr[i] = i + 1
	}

	fmt.Println(arr[15:30])
	fmt.Println(arr[20:])
	fmt.Println(arr[:2])
	// 原有切片
	fmt.Println(arr[:])
	//重置切片
	a := arr[0:0]
	fmt.Println(a)
}

// 声明切片
func test3() {

	var strList []string
	var intList []int
	var emptyList = []int{}

	fmt.Println(len(strList), len(intList), len(emptyList))
	//声明但未使用的切片的默认值是 nil。strList 和 numList 也是 nil，所以和 nil 比较的结果是 true。
	fmt.Println(strList == nil)
	fmt.Println(intList == nil)
	//numListEmpty 已经被分配到了内存，但没有元素，因此和 nil 比较时是 false。
	fmt.Println(emptyList == nil)
}

//2、使用 make() 函数构造切片
//格式：make([]T, size, cap)
//	T：切片的元素类型。
//  size：就是为这个类型分配多少个元素。
//  cap：预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。
func test4() {
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	// cap参数的值不会影响切片的长度
	fmt.Println(len(a), len(b))
}

//使用 make() 函数生成的切片一定发生了内存分配操作。
//但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，
//设定开始与结束位置，不会发生内存分配操作。

func main() {
	test4()
}
