package cls1

import (
	"../../../08.包/8.05_工厂模式自动注册/base"
	"fmt"
)

type Class1 struct{}

// 实现 Class 接口的 Do 方法
func (c *Class1) Do() {
	fmt.Println("class1 Do方法")
}

func init() {
	base.Register("class1", func() base.Class {
		return new(Class1)
	})
}
