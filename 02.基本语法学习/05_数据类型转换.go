package _2_基本语法学习

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("int8 range:", math.MaxInt8)
	fmt.Println("int16 range:", math.MaxInt16)
	fmt.Println("int32 range:", math.MaxInt32)
	fmt.Println("int64 range:", math.MaxInt64)

	var a int32 = 1047483647
	// %x 整型以十六进制输出，字母小写方式显示
	// %X 整型以十六进制输出，字母大写方式显示
	fmt.Printf("int32:0x%x %d\n", a, a)

	// 将a变量数值转换为十六进制, 发生数值截断
	b := int16(a)
	fmt.Printf("int32:0x%X %d\n", b, b)

	var c float32 = math.Pi
	// 将浮点型转为整型，小数点后的数据会丢失
	fmt.Println(int(c))
}
