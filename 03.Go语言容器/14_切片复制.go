package _3_Go语言容器

import "fmt"

//使用 Go 语言内建的 copy() 函数，可以迅速地将一个切片的数据复制到另外一个切片空间中，copy() 函数的使用格式如下：
//copy( destSlice, srcSlice []T) int
// srcSlice 为数据来源切片。
// destSlice 为复制的目标。目标切片必须分配过空间且足够承载复制的元素个数。来源和目标的类型一致，copy 的返回值表示实际发生复制的元素个数。

func test() {

	const elementCount = 1000

	sourceData := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		sourceData[i] = i
	}

	refData := sourceData
	copyData := make([]int, elementCount)

	copy(copyData, sourceData)
	sourceData[0] = 999
	fmt.Println(refData[0])

	fmt.Println(copyData[0], copyData[elementCount-1])
	//将 srcData 的局部数据复制到 copyData 中,只替换掉了copyData中位置0、1的数据，其他位置的数据不变
	copy(copyData, sourceData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d,", copyData[i])
	}
}

func main() {
	test()
}
