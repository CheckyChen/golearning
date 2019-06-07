package main

import (
	"fmt"
	"sync"
)

// 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个 goroutine 可以访问共享资源

var (
	// 逻辑中使用的变量
	count int
	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
)

func GetCount() int {
	// 锁定
	countGuard.Lock()
	// 在函数退出时解除锁定
	defer countGuard.Unlock()

	return count
}

func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

// 一般情况下，建议将互斥锁的粒度设置得越小越好，降低因为共享访问时等待的时间。
// 这里好的习惯性地将互斥锁的变量命名为以下格式： 变量名+Guard

// 在【读多写少】的环境中，可以优先使用读写互斥锁（sync.RWMutex），它比互斥锁更加高效。
// sync 包中的 RWMutex 提供了读写互斥锁的封装。
//
// 我们将互斥锁例子中的一部分代码修改为读写互斥锁，参见下面代码：
var (
	count1      int
	count1Guard sync.RWMutex
)

func GetCOunt1() int {
	count1Guard.Lock()
	// 在函数退出时解除锁定
	defer count1Guard.RUnlock()
	return count1
}

func main() {
	for i := 0; i < 100; i++ {
		go SetCount(i)
	}
	//SetCount(1)
	fmt.Println(GetCount())
}
