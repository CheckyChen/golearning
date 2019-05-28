package _2_基本语法学习

import "fmt"

//Go 语言中现阶段没有枚举，可以使用 const 常量配合 iota 模拟枚举

type Weapon int

func enum() {
	const (
		Arrow Weapon = iota // 开始生成枚举值，默认为0
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	// 输出所有的枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)

	// 使用枚举类型并赋值
	var weapon Weapon = Blower
	fmt.Println(weapon)
}

// iota 不仅只生成每次增加 1 的枚举值。我们还可以利用 iota 来做一些强大的枚举常量值生成器
func enum2() {
	const (
		// iota 使用了一个移位操作，每次将上一次的值左移一位，以做出每一位的常量值。
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)

	fmt.Printf(" %d %d %d\n", FlagRed, FlagGreen, FlagBlue)

	//将枚举值按二进制格式输出，可以清晰地看到每一位的变化。
	fmt.Printf(" %b %b %b\n", FlagRed, FlagGreen, FlagBlue)
}

// 将枚举值转化为字符串
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

//定义 ChipType 类型的方法 String()，返回字符串。
func (c ChipType) String() string {
	//使用 switch 语句判断当前的 ChitType 类型的值，返回对应的字符串。
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}

func main() {
	enum()
	enum2()
	//输出 CPU 的值并按整型格式输出。
	fmt.Printf("%s  %d  %s", CPU, GPU, GPU)
}
