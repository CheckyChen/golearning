package main

import (
	"fmt"
	"sort"
)

// Go语言中在排序时，需要使用者通过 sort.Interface 接口提供数据的一些特性和操作方法

type Interface interface {
	// 元素的数量
	Len() int
	// 小于比较
	Less(i, j int) bool
	// 交换元素
	Swap(i, j int)
}

// 使用sort.Interface接口进行排序
type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// 1、常见类型的便捷排序

// 通过实现 sort.Interface 接口的排序过程具有很强的可定制性，可以根据被排序对象比较复杂的特性进行定制。
// 例如，需要多种排序逻辑的需求就适合使用 sort.Interface 接口进行排序。
// 但大部分情况中，只需要对字符串、整型等进行快速排序。Go 语言中提供了一些固定模式的封装以方便开发者迅速对内容进行排序。

//  1) 字符串切片的便捷排序
//  sort 包中有一个 StringSlice 类型，定义如下：
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p StringSlice) Sort()              { sort.Sort(p) }

// sort 包中的 StringSlice 的代码与 MyStringList 的实现代码几乎一样。
// 因此，只需要使用 sort 包的 StringSlice 就可以更简单快速地进行字符串排序。
// 将代码1中的排序代码简化后如下所示：

func simpleStringSort() {
	names := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(names)
	for _, value := range names {
		fmt.Println(value)
	}
}

// 可以用sort.Strings()进一步简化
func simpleStringsSort() {
	names := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Strings(names)
	for _, value := range names {
		fmt.Println(value)
	}
}

// 2) 对整型切片进行排序
//
// 除了字符串可以使用 sort 包进行便捷排序外，还可以使用 sort.IntSlice 进行整型切片的排序。sort.IntSlice 的定义如下：
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p IntSlice) Sort() { sort.Sort(p) }

// 3) sort包内建的类型排序接口一览

//  类型		实现接口sort.Interface的类型			直接排序							说明
//	String		StringSlice						sort.Strings(a []string)		字符ASCII值升序
//	int			IntSlice						sort.Ints(a []int)				数值升序
//	float64		Float64Slice					sort.Float64s(a []float64)		数值升序

// 2、对结构体数据进行排序
// 除了基本类型的排序，也可以对结构体进行排序。结构体比基本类型更为复杂，排序时不能像数值和字符串一样拥有一些固定的单一原则。结构体的多个字段在排序中可能会存在多种排序的规则，例如，结构体中的名字按字母升序排列，数值按从小到大的顺序排序。一般在多种规则同时存在时，需要确定规则的优先度，如先按名字排序，再按年龄排序等。
//	1) 完整实现sort.Interface进行结构体排序
//	将一批英雄名单使用结构体定义，英雄名单的结构体中定义了英雄的名字和分类。排序时要求按照英雄的分类进行排序，相同分类的情况下按名字进行排序，详细代码实现过程如下。
type HeroKind int

const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heros []*Hero

func (s Heros) Len() int { return len(s) }

func (s Heros) Less(i, j int) bool {
	// 如果英雄的分类不一致时, 优先对分类进行排序
	if s[i].Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}
	return s[i].Name < s[j].Name
}

func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 2） 使用sort.Slice进行切片元素排序
func sliceSort() {
	heros := []*Hero{
		{"吕布", Tank},
		{"李白", Assassin},
		{"妲己", Mage},
		{"貂蝉", Assassin},
		{"关羽", Tank},
		{"诸葛亮", Mage},
	}

	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})

	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}

func main() {
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(names)

	for _, value := range names {
		fmt.Println(value)
	}
	fmt.Println("=====================")
	// 以上代码可以简化为这个方法，同样实现了排序
	simpleStringSort()
	fmt.Println("=====================")
	simpleStringsSort()

	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}
	sort.Sort(heros)
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("=====================")
	sliceSort()
}
