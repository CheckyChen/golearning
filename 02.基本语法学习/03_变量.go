package _2_基本语法学习

import "fmt"

// 1、变量（Variable）的功能是存储用户的数据。不同的逻辑有不同的对象类型，也就有不同的变量类型。
// 常见变量的数据类型有：整型、浮点型、布尔型、结构体等。

// 标准格式：var 变量名 变量类型

// 1.1、基本声明形式
// 声明一个整形类型的变量
var a int

// 声明一个字符串类型的变量
var b string

// 声明一个32位浮点切片类型的变量，浮点切片表示由多个浮点类型组成的数据结构
var c []float32

// 声明一个返回值为布尔类型的函数变量，这种形式一般用户回调函数，即将函数以变量的形式保存下来，在需要的时候重新调用这个函数
var d func() bool

// 声明一个结构体类型的变量，这个结构体拥有一个整形的x字段
var e struct{ x int }

// 1.2、批量声明形式
var (
	f int
	g string
	h []float32
	i func() bool
	j struct{ x int }
)

// 2、变量的初始化

// 标准格式：var 变量名 类型 = 表达式

// 2.1 普通初始化形式
var hp int = 100

var attack = 40
var defence = 20

// 强制使用float32类型，如果不声明为float32位的话，系统默认会使用float64
var damageRate float32 = 0.17

// float32() 强制类型转换
var damage = float32(attack-defence) * damageRate

// 2.2 短变量声明并初始化，格式：变量名 := 表达式
// 注意：只能用户函数内部局部变量，不能用于包级别中

func test() {

	hp1 := 100
	fmt.Println(hp1)

	var a1 int = 200
	// 重复声明了，会报错 no new variables on left side of :=
	// “:=”的左边变量已经被声明了。
	//a1 := 100
	fmt.Println(a1)

	a2, b2 := 1, 2
	fmt.Println(a2, b2)
	// 注意：在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中，
	// 即便其他变量名可能是重复声明的，编译器也不会报错，比如：b2 没有重复声明，只是被重新赋值了
	a3, b2 := 3, 4
	fmt.Println(a3, b2)
}

func swap1() {
	var a int = 100
	var b int = 200
	var t int
	fmt.Println("交换前：a=", a, "b=", b)
	t = a
	a = b
	b = t
	fmt.Println("交换后：a=", a, "b=", b)
}

func swap2() {
	a := 100
	b := 200
	a = a ^ b
	fmt.Println(a, b)
	b = b ^ a
	fmt.Println(a, b)
	a = a ^ b
	fmt.Println(a, b)
}

func swap3() {
	a := 100
	b := 200
	a, b = b, a
	fmt.Println(a, b)
}

// 将 []int 声明为 IntSlice 类型。
type IntSlice []int

// 为这个类型编写一个 Len 方法，提供切片的长度。
func (p IntSlice) Len() int { return len(p) }

// 根据提供的 i、j 元素索引，获取元素后进行比较，返回比较结果。
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }

// 根据提供的 i、j 元素索引，交换两个元素的值。
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// 3、匿名变量
//在使用多重赋值时，如果不需要在左值中接收变量，可以使用匿名变量（anonymous variable）。
//匿名变量的表现是一个下画线_，使用匿名变量时，只需要在变量声明的地方使用下画线替换即可。
//匿名变量不占用命名空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。

func GetData() (int, int) {
	return 100, 200
}

func main() {
	//test()
	//fmt.Println(damage)
	//swap1()
	//swap2()
	//swap3()

	a, _ := GetData()
	//fmt.Println(a, b)
	_, b := GetData()
	fmt.Println(a, b)

}
