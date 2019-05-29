package main

import (
	"errors"
	"fmt"
	"net"
)

// Go 语言的错误处理思想及设计包含以下特征：
// 一个可能造成错误的函数，需要返回值中返回一个错误接口（error）。如果调用是成功的，错误接口将返回 nil，否则返回错误。
// 在函数调用后需要检查错误，如果发生错误，进行必要的错误处理。
//
// Go 语言没有类似 Java 或 .NET 中的异常处理机制，虽然可以使用 defer、panic、recover 模拟，但官方并不主张这样做。
// Go 语言的设计者认为其他语言的异常机制已被过度使用，上层逻辑需要为函数发生的异常付出太多的资源。
// 同时，如果函数使用者觉得错误处理很麻烦而忽略错误，那么程序将在不可预知的时刻崩溃。
//
// Go 语言希望开发者将错误处理视为正常开发必须实现的环节，正确地处理每一个可能发生错误的函数。
// 同时，Go 语言使用返回值返回错误的机制，也能大幅降低编译器、运行时处理错误的复杂度，让开发者真正地掌握错误的处理

// 1、net 包中的例子
// net.Dial() 是 Go 语言系统包 net 即中的一个函数，一般用于创建一个 Socket 连接。
//
// net.Dial 拥有两个返回值，即 Conn 和 error。这个函数是阻塞的，因此在 Socket 操作后，
// 会返回 Conn 连接对象和 error；如果发生错误，error 会告知错误的类型，Conn 会返回空。
//
// 根据 Go 语言的错误处理机制，Conn 是其重要的返回值。因此，为这个函数增加一个错误返回，类似为 error。参见下面的代码：
func Dial(network, address string) (net.Conn, error) {
	var d net.Dialer
	return d.Dial(network, address)
}

// 在io包中的Writer接口也拥有错误返回，代码如下：
type Writer interface {
	Writer(p []byte) (n int, err error)
}

// io包中的Closer接口，只有一个错误返回，代码如下：
type Closer interface {
	Close() error
}

// 2、错误接口的定义格式
// 所有符合Error()string格式的方法，都能实现错误接口。
// Error() 方法返回错误的具体描述，使用者可以通过这个字符串知道发生了什么错误。
// error是Go系统声明的接口类型，代码如下：
type error interface {
	Error() string
}

// 3、定义一个错误
// 返回错误前，需要定义会产生哪些可能的错误。在 Go 语言中，使用 errors 包进行错误的定义，格式如下：
var err = errors.New("这是一个错误")

// 错误字符串由于相对固定，一般在包作用域声明，应尽量减少在使用时直接使用 errors.New 返回。

// 1）errors包
// Go语言的errors中对New的定义非常简单，如下：

// 创建错误对象
func New(text string) error {
	// 将errorString结构体实例化，并赋值错误描述的成员
	return &errorString{text}
}

// 错误字符串
type errorString struct {
	// 声明errorString结构体，拥有一个成员,描述错误内容
	s string
}

// 返回发生何种错误
func (e *errorString) Error() string {
	// 实现error接口的Error方法，该方法 返回成员中的错误描述
	return e.s
}

// 2）在代码中使用错误定义
var errDivisionByZero = errors.New("devision by zero")

func div(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	return dividend / divisor, nil
}

// 示例：在解析中使用自定义错误
// 使用 errors.New 定义的错误字符串的错误类型是无法提供丰富的错误信息的。
// 那么，如果需要携带错误信息返回，就需要借助自定义结构体实现错误接口。
//
// 下面代码将实现一个解析错误（ParseError），这种错误包含两个内容：文件名和行号。
// 解析错误的结构还实现了 error 接口的 Error() 方法，返回错误描述时，就需要将文件名和行号返回。

// 声明一个解析错误
type ParseError struct {
	Filename string // 文件名
	Line     int    //行号
}

// 实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprint("%s:%d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func main() {

	fmt.Println(div(12, 0))

	var e error
	// 创建一个错误实例，包含文件名和行号
	e = newParseError("main.go", 1)
	// 通过error接口查看错误描述
	fmt.Println(e.Error())
	switch detail := e.(type) {
	case *net.ParseError:
		fmt.Printf("Filename: %s Line: %d\n", detail.Text, detail.Text)
	default:
		fmt.Println("other error")
	}
}
