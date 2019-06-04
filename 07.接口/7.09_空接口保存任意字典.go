package main

import "fmt"

// 空接口可以保存任何类型这个特性可以方便地用于容器的设计。下面例子使用 map 和 interface{} 实现了一个字典。
// 字典在其他语言中的功能和 map 类似，可以将任意类型的值做成键值对保存，然后进行找回、遍历操作

type Dictionary struct {
	data map[interface{}]interface{} // 键值对都为interface{}类型
}

// 根据键获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

// 设置键值
func (d *Dictionary) Set(key, value interface{}) {
	d.data[key] = value
}

// 遍历所有的键值，如果回调返回值为false，停止遍历
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	// 当 callback 为空时，退出遍历，避免后续代码访问空的 callback 而导致的崩溃
	if callback == nil {
		return
	}

	for k, v := range d.data {
		//根据 callback 的返回值，决定是否继续遍历，返回为 false 时停止遍历
		if !callback(k, v) {
			return
		}
	}
}

// 清空所有值
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

// 创建一个字典
func NewDictionay() *Dictionary {
	d := &Dictionary{}
	// 初始化map
	d.Clear()
	return d
}

func main() {
	dict := NewDictionay()
	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Don't Hungry", 24)

	// 获取值及打印值
	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite:", favorite)

	dict.Visit(func(key, value interface{}) bool {
		if value.(int) > 40 {
			fmt.Println(key, "太贵了")
			return true
		}

		fmt.Println(key, "便宜的")
		return true
	})
}
