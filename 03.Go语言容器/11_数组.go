package _3_Go语言容器

import "fmt"

// 数组：一段固定长度的连续内存区域
// 在GO语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变

// 1、数组的生命
// 格式：var 数组变量名 [元素数量]T
//		数组变量名：数组声明及使用时的变量名。
//		元素数量：数组的元素数量。可以是一个表达式，但最终通过编译期计算的结果必须是整型数值。也就是说，元素数量不能含有到运行时才能确认大小的数值。
//		T 可以是任意基本类型，包括 T 为数组本身。但类型为数组本身时，可以实现多维数组。

func test() {
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"

	fmt.Println(team)
}

// 2、数组的初始化
func test2() {
	// ... 表示让编译器确定数组的大小。
	var team2 = [...]string{"hammer", "soldier", "num"}
	fmt.Println(team2)
}

//3、遍历数组
func test3() {
	var team3 = [...]int{1, 2, 3, 4, 5, 6}
	//使用 for 循环，遍历 team 数组，遍历出的键 k 为数组的索引，值 v 为数组的每个元素值。
	for k, v := range team3 {
		fmt.Println(k, v)
	}
}

func main() {
	//test()
	test3()
}
