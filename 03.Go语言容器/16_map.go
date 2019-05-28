package _3_Go语言容器

import (
	"fmt"
	"sort"
	"sync"
)

//在业务和算法中需要使用任意类型的关联关系时，就需要使用到映射，如学号和学生的对应、名字与档案的对应等。
//
//Go 语言提供的映射关系容器为 map，map使用散列表（hash）实现。
//提示
//
//大多数语言中映射关系容器使用两种算法：散列表和平衡树。
//
//散列表可以简单描述为一个数组（俗称“桶”），数组的每个元素是一个列表。根据散列函数获得每个元素的特征值，将特征值作为映射的键。如果特征值重复，表示元素发生碰撞。碰撞的元素将被放在同一个特征值的列表中进行保存。散列表查找复杂度为 O(1)，和数组一致。最坏的情况为 O(n)，n 为元素总数。散列需要尽量避免元素碰撞以提高查找效率，这样就需要对“桶”进行扩容，每次扩容，元素需要重新放入桶中，较为耗时。
//
//平衡树类似于有父子关系的一棵数据树，每个元素在放入树时，都要与一些节点进行比较。平衡树的查找复杂度始终为 O(log n)。

//添加关联到 map 并访问关联和数据
//
//Go 语言中 map 的定义是这样的：
//map[KeyType]ValueType
//	KeyType为键类型。
//	ValueType是键对应的值类型。
//
//一个 map 里，符合 KeyType 和 ValueType 的映射总是成对出现。

func map_test() {
	user := make(map[string]int)
	user["age"] = 24
	fmt.Println(user["age"])

	v := user["age1"]
	//map中并没有age1这个键，所以返回的为 ValueType 的默认值。 0
	fmt.Println(v)
}

// 多个键一起声明
func map_test1() {
	m := map[string]string{
		"W": "forward",
		"A": "left",
		"D": "right",
		"S": "backward",
	}
	fmt.Println(m)
}

// 遍历map
//map 的遍历过程使用 for range 循环完成
func foreachMap() {
	scene := make(map[string]string)
	scene["name"] = "checky"
	scene["age"] = "24"
	scene["gender"] = "male"
	for k, v := range scene {
		fmt.Println(k, v)
	}
}

// 只遍历值的话，可以使用 for _,value range scene{}，将不需要的值用匿名变量形式
// 只遍历键的话，可以使用 for k range scene{}，直接将value的值忽略掉即可

// ※ 注意：遍历输出元素的顺序与填充顺序无关。不能期望 map 在遍历时返回某种期望顺序的结果。

// 如果想以特定的顺序遍历的话，可以先对map的键排序

func map_test2() {
	scene := make(map[string]string)
	scene["name"] = "checky"
	scene["age"] = "24"
	scene["gender"] = "male"
	var sceneList []string
	for k := range scene {
		sceneList = append(sceneList, k)
	}
	//	对切片进行排序
	sort.Strings(sceneList)

	fmt.Println(sceneList)
}

// map元素的删除  delete(map,key)
//   map 要删的map实例
//   key 删除的key
func deleteMapKey() {
	scene := make(map[string]int)
	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	delete(scene, "brazil")
	for k, v := range scene {
		fmt.Println(k, v)
	}
}

//清空 map 中的所有元素
//
//有意思的是，Go 语言中并没有为 map 提供任何清空所有元素的函数、方法。
// 清空 map 的唯一办法就是重新 make 一个新的 map。不用担心垃圾回收的效率，Go 语言中的并行垃圾回收效率比写一个清空函数高效多了。

// sync.Map
//sync.Map有以下特性：
//	无须初始化，直接声明即可。
//	sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用。Store 表示存储，Load 表示获取，Delete 表示删除。
//	使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值。Range 参数中的回调函数的返回值功能是：需要继续迭代遍历时，
//返回 true；终止迭代遍历时，返回 false。

func syncMapTest() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}
func main() {
	//map_test()
	//map_test1()
	//foreachMap()
	//map_test2()
	//deleteMapKey()
	syncMapTest()
}
