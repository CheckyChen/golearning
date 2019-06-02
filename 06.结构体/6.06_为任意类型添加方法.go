package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// Go 语言可以对任何类型添加方法。给一种类型添加方法就像给结构体添加方法一样，因为结构体也是一种类型
type MyInt int

func (m MyInt) IsZero() bool {
	return m == 0
}

func (m MyInt) Add(other int) int {
	return other + int(m)
}

// http包中的类型方法
// Go 语言提供的 http 包里也大量使用了类型方法。
// Go 语言使用 http 包进行 HTTP 的请求，使用 http 包的 NewRequest() 方法可以创建一个 HTTP 请求，
// 填充请求中的 http 头（req.Header），再调用 http.Client 的 Do 包方法，将传入的 HTTP 请求发送出去。

func HttpTest() {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://www.163.com/", strings.NewReader("key=value"))
	// 如果发现错误，打印退出
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	req.Header.Add("User-Agent", "myClient")
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	defer response.Body.Close()
}

// time包中的类型方法
// Go 语言提供的 time 包主要用于时间的获取和计算等。在这个包中，也使用了类型方法

func GetSecond() string {
	return time.Second.String()
}

func main() {
	var b MyInt
	fmt.Println(b.IsZero())

	b = 1
	fmt.Println(b.Add(2))

	HttpTest()

	fmt.Println(GetSecond())
}
