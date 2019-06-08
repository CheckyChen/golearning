package main

import (
	"fmt"
	"reflect"
)

// Go 程序中对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型。这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作，代码如下：

func main() {

	type cat struct{}
	//创建了cat结构体的实例，ins 是一个 *cat 类型的指针变量
	ins := &cat{}

	typeOfCat := reflect.TypeOf(ins)
	//Go 语言的反射中对所有指针变量的种类都是 Ptr，但注意，指针变量的类型名称是空，不是 *cat
	fmt.Printf("name:'%v',kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())

	typeOfCat = typeOfCat.Elem()
	//变量指向元素的类型名称和种类，得到了 cat 的类型名称（cat）和种类（struct）
	fmt.Printf("element name:'%v',element kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())
}
