package main

import (
	"fmt"
	"reflect"
)

// 通过 reflect.Type 获取结构体成员信息 reflect.StructField 结构中的 Tag 被称为结构体标签（Struct Tag）。结构体标签是对结构体字段的额外信息标签。
//
// JSON、BSON 等格式进行序列化及对象关系映射（Object Relational Mapping，简称 ORM）系统都会用到结构体标签，这些系统使用标签设定字段在处理时应该具备的特殊属性和可能发生的行为。这些信息都是静态的，无须实例化结构体，可以通过反射获取到。
// 提示
// 结构体标签（Struct Tag）类似于 C# 中的特性（Attribute）。C# 允许在类、字段、方法等前面添加 Attribute，然后在反射系统中可以获取到这个属性系统。例如：
// [Conditional("DEBUG")]
// public static void Message(string msg)
// {
//    Console.WriteLine(msg)；
// }

// 1、结构体标签的格式

// Tag在结构体字段后方书写格式： `key:"value" key2:"value2"`
// 结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。

// 2、从结构体标签中获取值

// StructTag 拥有一些方法，可以进行 Tag 信息的解析和提取，如下所示：
//	func(tag StructTag)Get(key string)string
//	根据 Tag 中的键获取对应的值，例如 `key1:"value1"key2:"value2"` 的 Tag 中，可以传入“key1”获得“value1”。
//	func(tag StructTag)Lookup(key string)(value string,ok bool)
//	根据 Tag 中的键，查询值是否存在。

// 3、结构体标签格式错误导致的问题

// 编写 Tag 时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，参见下面这个例子：
func main() {
	type cat struct {
		Name string
		Type int `json: "type" id:"1000"` // 这里故意在 json: 和 "type" 之间加了个空格
	}

	typeOfCat := reflect.TypeOf(cat{})
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 输出空字符串，因为 json 键和值之间多了个 空格
		fmt.Println(catType.Tag.Get("json"))
	}
}
