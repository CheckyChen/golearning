package main

import (
	"fmt"
	"sync/atomic"
)

// 当多线程并发运行的程序竞争访问和修改同一块资源时，会发生竞态问题。

var seq int64

func GenID() int64 {
	atomic.AddInt64(&seq, 1)
	return seq
}

func main() {
	for i := 0; i < 10; i++ {
		go GenID()
	}

	fmt.Println(GenID())
}
