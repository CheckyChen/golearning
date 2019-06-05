package main

import (
	"../8.05_工厂模式自动注册/base"
	_ "../8.05_工厂模式自动注册/cls1"
	_ "../8.05_工厂模式自动注册/cls2"
)

func main() {
	c1 := base.Create("class1")
	c1.Do()

	c2 := base.Create("class2")
	c2.Do()
}
