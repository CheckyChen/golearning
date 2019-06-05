package cls2

import (
	"../../../08.包/8.05_工厂模式自动注册/base"
	"fmt"
)

type Class2 struct{}

// 实现 Class 接口的 Do 方法
func (c *Class2) Do() {
	fmt.Println("class2 Do 方法")
}

func init() {
	base.Register("class2", func() base.Class {
		return new(Class2)
	})
}
