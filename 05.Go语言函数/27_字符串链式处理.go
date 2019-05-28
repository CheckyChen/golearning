package main

import (
	"fmt"
	"strings"
)

// 使用 SQL 语言从数据库中获取数据时，可以对原始数据进行排序（sort by）、分组（group by）和去重（distinct）等操作。
// SQL 将数据的操作与遍历过程作为两个部分进行隔离，这样操作和遍历过程就可以各自独立地进行设计，这就是常见的数据与操作分离的设计。
//
// 对数据的操作进行多步骤的处理被称为链式处理。
// 本例中使用多个字符串作为数据集合，然后对每个字符串进行一系列的处理，用户可以通过系统函数或者自定义函数对链式处理中的每个环节进行自定义。

func StringProccess(list []string, chain []func(string) string) {
	// 遍历每一个字符串
	for index, str := range list {
		// 第一个需要处理的字符串
		result := str
		// 遍历每一个处理链
		for _, proc := range chain {
			// 输入一个字符串进行处理，返回数据作为下一个处理链的输入
			result = proc(result)
		}
		// 将结果放回切片
		list[index] = result
	}
}

// 移除前缀的处理函数
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	// 处理函数链
	chain := []func(string) string{
		// 1.移除前缀
		removePrefix,
		// 2.去掉空格
		strings.TrimSpace,
		// 3.转化为大写
		strings.ToUpper,
	}
	// 处理字符串
	StringProccess(list, chain)
	// 输出处理好的字符串
	for _, str := range list {
		fmt.Println(str)
	}
}
