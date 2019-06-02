package main

import "fmt"

// 1、使用“键值对”初始化结构体
// 键值对的填充是可选的，不需要初始化的字段可以不填入初始化列表中。
// 结构体实例化后字段的默认值是字段类型的默认值，例如：数值为 0，字符串为空字符串，布尔为 false，指针为 nil 等。

// 1) 键值对初始化结构体的书写格式
// ins := 结构体类型名{
//    字段1: 字段1的值,
//    字段2: 字段2的值,
//    …
//}

type User struct {
	name   string
	age    int
	gender string
}

func userGender() {
	u := User{
		name: "张三",
		age:  25,
	}
	fmt.Println(u)

	//键值之间以:分隔；键值对之间以,分隔
}

// 2、使用多个值的列表初始化结构体
// Go语言可以在“键值对”初始化的基础上忽略“键”。也就是说，可以使用多个值的列表初始化结构体的字段。

// 1) 多个值列表初始化结构体的书写格式
// 多个值使用逗号分隔初始化结构体，例如：
// ins := 结构体类型名{
//    字段1的值,
//    字段2的值,
//    …
// }

type Address struct {
	Province string
	City     string
	ZipCode  int
	Phone    string
}

func address() {
	address := Address{
		"江西省",
		"九江市",
		332305,
		"13333333333",
	}
	fmt.Println(address)

	// 使用这种格式初始化时，需要注意：
	//  必须初始化结构体的所有字段。
	//  每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	//  键值对与值列表的初始化形式不能混用。
}

// 3、初始化匿名结构体
// 匿名结构体没有类型名称，无须通过type关键字定义就可以直接使用。
//1) 匿名结构体定义格式和初始化写法
//
//匿名结构体的初始化写法由结构体定义和键值对初始化两部分组成。结构体定义时没有结构体类型名，只有字段和类型定义。键值对初始化部分由可选的多个键值对组成，如下格式所示：
//ins := struct {
//    // 匿名结构体字段定义
//    字段1 字段类型1
//    字段2 字段类型2
//    …
//}{
//    // 字段值初始化
//    初始化字段1: 字段1的值,
//    初始化字段2: 字段2的值,
//    …
//}

func noNameStruct() {
	ins := struct {
		field1 string
		field2 string
		field3 int
		field4 bool
	}{
		field1: "france",
		field2: "eric",
		field3: 24,
		field4: false,
	}

	fmt.Println(ins)

	//键值对初始化部分是可选的

	ins2 := struct {
		field1 string
		field2 string
		field3 int
		field4 bool
	}{}

	fmt.Println(ins2)
}

func main() {
	noNameStruct()
}
