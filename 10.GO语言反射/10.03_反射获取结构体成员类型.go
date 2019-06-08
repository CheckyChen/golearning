package main

import (
	"fmt"
	"reflect"
)

// 任意值通过 reflect.TypeOf() 获得反射对象信息后，如果它的类型是结构体， 可以通过反射值对象（reflect.Type）的 NumField() 和 Field()
// 方法获得结构体成员的详细信息。与成员获取相关的 reflect.Type 的方法如下表所示。

// Field(i int) StructField						根据索引，返回索引对应的结构体字段的信息。当值不是结构体或索引超界时发生宕机
// NumField() int								返回结构体成员字段数量。当类型不是结构体或索引超界时发生宕机
// FieldByName(name string) (StructField, bool)	根据给定字符串返回字符串对应的结构体字段的信息。没有找到时 bool 返回 false，当类型不是结构体或索引超界时发生宕机
// FieldByIndex(index []int) StructField		多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。没有找到时返回零值。当类型不是结构体或索引超界时 发生宕机
// FieldByNameFunc( match func(string) bool) (StructField,bool)	根据匹配函数匹配需要的字段。当值不是结构体或索引超界时发生宕机

// 1、结构体字段类型
//reflect.Type 的 Field() 方法返回 StructField 结构，这个结构描述结构体的成员信息，
// 通过这个信息可以获取成员与结构体的关系，如偏移、索引、是否为匿名字段、结构体标签（Struct Tag）等，
// 而且还可以通过 StructField 的 Type 字段进一步获取结构体成员的类型信息。StructField 的结构如下

//type StructField struct {
//    Name 		string     // 字段名
//    PkgPath 	string     // 字段路径
//    Type      Type       // 字段反射类型对象
//    Tag       StructTag  // 字段的结构体标签
//    Offset    uintptr    // 字段在结构体中的相对偏移
//    Index     []int      // Type.FieldByIndex中的返回的索引值
//    Anonymous bool       // 是否为匿名字段
//}

// 字段说明如下：
// 	Name：为字段名称。
// 	PkgPath：字段在结构体中的路径。
// 	Type：字段本身的反射类型对象，类型为 reflect.Type，可以进一步获取字段的类型信息。
// 	Tag：结构体标签，为结构体字段标签的额外信息，可以单独提取。
// 	Index：FieldByIndex 中的索引顺序。
// 	Anonymous：表示该字段是否为匿名字段。

// 2、获取成员反射信息
// 反射访问结构体成员类型及信息：

func main() {
	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
	}
	ins := cat{Name: "mini", Type: 1}

	typeOfCat := reflect.TypeOf(ins)

	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)

		fmt.Printf("name: %v tag:'%v'\n", fieldType.Name, fieldType.Tag)
	}

	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}
