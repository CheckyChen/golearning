package main

import "fmt"

// Go 语言的 switch 不仅可以像其他语言一样实现数值、字符串的判断，还有一种特殊的用途——判断一个接口内保存或实现的类型

// 类型断言的书写格式
//
// switch 实现类型分支时的写法格式如下：
// switch 接口变量.(type) {
//    case 类型1:
//        // 变量是类型1时的处理
//    case 类型2:
//        // 变量是类型2时的处理
//    …
//    default:
//        // 变量不是所有case中列举的类型时的处理
// }
// 对各个部分的说明：
// 	接口变量：表示需要判断的接口类型的变量。
// 	类型1、类型2……：表示接口变量可能具有的类型列表，满足时，会指定 case 对应的分支进行处理。

// 1、使用类型分支判断基本类型
func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, " is int")
	case string:
		fmt.Println(v, " is string")
	case bool:
		fmt.Println(v, " is bool")
	}
}

// 2、使用类型分支判断接口类型

// 电子支付方式
type Alipay struct{}

// 为电子支付方式添加CanUseFaceID()方法，表示能用刷脸支付方式
func (a *Alipay) CanUseFaceID() {}

// 现金支付方式
type Cash struct{}

// 为Cash添加Stolen()方法，表示现金支付式会出现偷窃情况
func (c *Cash) Stolen() {}

// 具备刷脸特性的接口
type ContainCanUseFaceID interface {
	CanUseFaceID()
}

// 具备被偷特性的接口
type ContainStolen interface {
	Stolen()
}

func print(payMethod interface{}) {
	switch payMethod.(type) {
	case ContainCanUseFaceID:
		fmt.Printf("%T 能刷脸\n", payMethod)
	case ContainStolen:
		fmt.Printf("%T 可能会被偷\n", payMethod)
	}
}

func main() {
	printType(1024)
	printType("123")
	printType(false)

	print(new(Alipay))
	print(new(Cash))
}
