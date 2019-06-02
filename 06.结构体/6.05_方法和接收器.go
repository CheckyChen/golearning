package main

import (
	"fmt"
	"math"
)

// Go 语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收器（Receiver）
// 如果将特定类型理解为结构体或“类”时，接收器的概念就类似于其他语言中的 this 或者 self

// 在 Go 语言中，接收器的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法
// 在面向对象的语言中，类拥有的方法一般被理解为类可以做的事情。
// 在 Go 语言中“方法”的概念与其他语言一致，只是 Go 语言建立的“接收器”强调方法的作用对象是接收器，也就是类实例，而函数没有作用对象

// 1、为结构体添加方法
// 使用背包作为“对象”，将物品放入背包的过程作为“方法”，通过面向过程的方式和 Go 语言中结构体的方式来理解“方法”的概念

type Bag struct {
	items []int
}

// 1)面向过程实现方法
func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

// 2）Go语言结构体方法
func (b *Bag) Insert(itemid int) {
	b.items = append(b.items, itemid)
}

// 2、接收器，方法作用的目标
// 格式:
// func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
//    函数体
// }
// 对各部分的说明：
//   接收器变量：接收器中的参数变量名在命名时，官方建议使用接收器类型名的第一个小写字母，而不是 self、this 之类的命名。例如，Socket 类型的接收器变量应该命名为 s，Connector 类型的接收器变量应该命名为 c 等。
//   接收器类型：接收器类型和参数类似，可以是指针类型和非指针类型。
//   方法名、参数列表、返回参数：格式与函数定义一致。

// 1) 理解指针类型的接收器
// 指针类型的接收器由一个结构体的指针组成，更接近于面向对象中的 this 或者 self。
// 由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的。
// 在下面的例子，使用结构体定义一个属性（Property），为属性添加 SetValue() 方法以封装设置属性的过程，
// 通过属性的 Value() 方法可以重新获得属性的数值。使用属性时，通过 SetValue() 方法的调用，可以达成修改属性值的效果。
type Property struct {
	value int
}

func (p *Property) SetValue(v int) {
	p.value = v
}

func (p *Property) Value() int {
	return p.value
}

// 2) 理解非指针类型的接收器
// 当方法作用于非指针接收器时，Go 语言会在代码运行时将接收器的值复制一份。
// 在非指针接收器的方法中可以获取接收器的成员值，但修改后无效。
// 点（Point）使用结构体描述时，为点添加 Add() 方法，这个方法不能修改 Point 的成员 X、Y 变量，
// 而是在计算后返回新的 Point 对象。Point 属于小内存对象，在函数返回值的复制过程中可以极大地提高代码运行效率，
// 详细过程请参考下面的代码

type Point struct {
	X int
	Y int
}

// 非指针接收器的加方法
func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

// 在计算机中，小对象由于值复制时的速度较快，所以适合使用非指针接收器。
// 大对象因为复制性能较低，适合使用指针接收器，在接收器和参数间传递时不进行复制，只是传递指针

// 3、二维矢量模拟玩家移动
// 在游戏中，一般使用二维矢量保存玩家的位置。使用矢量运算可以计算出玩家移动的位置。
// 本例子中，首先实现二维矢量对象，接着构造玩家对象，最后使用矢量对象和玩家对象共同模拟玩家移动的过程

// 1) 实现二维矢量结构
// 矢量是数学中的概念，二维矢量拥有两个方向的信息，同时可以进行加、减、乘（缩放）、距离、单位化等计算。
// 在计算机中，使用拥有 X 和 Y 两个分量的 Vec2 结构体实现数学中二维向量的概念。详细实现请参考下面的代码。

type Vec2 struct {
	X, Y float32
}

// 加
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

// 减
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// 乘
func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// 距离
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y

	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// 插值
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y + v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}
	return Vec2{0, 0}
}

// 2）实现玩家对象
// 玩家对象负责存储玩家的当前位置、目标位置和速度。使用 MoveTo() 方法为玩家设定移动的目标，使用 Update() 方法更新玩家位置。
// 在 Update() 方法中，通过一系列的矢量计算获得玩家移动后的新位置

type Player struct {
	currPos   Vec2    //当前位置
	targetPos Vec2    // 目标位置
	speed     float32 // 移动速度
}

// 移动到某个点，就是设置目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// 获取当前位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// 是否到达
func (p *Player) IsArrived() bool {
	// 如果当前位置距离目标位置的距离不超过步长，则说明到达了
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

// 逻辑更新
func (p *Player) Update() {
	if !p.IsArrived() {
		// 当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()
		// 添加速度矢量生成新的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))
		// 移动完成后，更新当前位置
		p.currPos = newPos

	}
}

// 创建一个新玩家
func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}

func main() {
	bag := new(Bag)
	Insert(bag, 1001)
	bag.Insert(2001)
	fmt.Println(bag)

	p := new(Property)
	p.SetValue(100)
	fmt.Println(p.Value())

	p1 := Point{1, 1}
	p2 := Point{2, 1}
	result := p1.Add(p2)
	fmt.Println(result)

	player := NewPlayer(0.5)
	player.MoveTo(Vec2{3, 1})
	// 如果玩家没有达到目标位置就一直循环
	for !player.IsArrived() {
		player.Update()
		fmt.Println(player.Pos())
	}

}
