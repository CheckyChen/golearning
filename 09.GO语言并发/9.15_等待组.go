package main

import (
	"fmt"
	"net/http"
	"sync"
)

// 除了可以使用通道（channel）和互斥锁进行两个并发程序间的同步外，还可以使用等待组进行多个任务的同步，
// 等待组可以保证在并发环境中完成指定数量的任务

// 等待组有下面几个方法可用
// ①(wg * WaitGroup) Add(delta int)	等待组的计数器 +1
// ②(wg *WaitGroup) Done()				等待组的计数器 -1
// ③(wg *WaitGroup) Wait()				当等待组计数器不等于 0 时阻塞直到变 0。

// 等待组内部拥有一个计数器，计数器的值可以通过方法调用实现计数器的增加和减少。
// 当我们添加了 N 个并发任务进行工作时， 就将等待组的计数器值增加 N。每个任务完成时，这个值减 1。
// 同时，在另外一个 goroutine 中等待这个等待组的计数器值为 0 时，表示所有任务已经完成。

func main() {
	var wg sync.WaitGroup

	var urls = []string{
		"https://www.github.com/",
		"https://www.baidu.com/",
		"https://www.google.com/",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			// 使用defer，表示函数完成时将等待组值减1
			defer wg.Done()
			// 使用http访问提供的url
			// Get() 函数会一直阻塞直到网站响应或者超时
			_, err := http.Get(url)
			fmt.Println(url, err)
		}(url)
	}
	// 等待所有的任务完成
	wg.Wait()

	fmt.Println("完成")
}
