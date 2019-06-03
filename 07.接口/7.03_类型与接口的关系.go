package main

import (
	"fmt"
	"io"
)

// 类型和接口之间有一对多和多对一的关系
// 1、一个类型可以实现多个接口
// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现
// 网络上的两个程序通过一个双向的通信连接实现数据的交换，连接的一端称为一个 Socket。Socket 能够同时读取和写入数据，这个特性与文件类似。因此，开发中把文件和 Socket 都具备的读写特性抽象为独立的读写器概念。
//
// Socket 和文件一样，在使用完毕后，也需要对资源进行释放。
//
// 把 Socket 能够写入数据和需要关闭的特性使用接口来描述，请参考下面的代码：

type Socket struct{}

func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (s *Socket) Close() error {
	return nil
}

// Socket结构的Write()方法实现了io.Write接口：
type Writer interface {
	Write(p []byte) (n int, err error)
}

// Socket结构也实现了io.Closer接口：
type Closer interface {
	Close() error
}

// 在代码中使用Socket结构实现的Writer接口和Closer接口代码如下：
// 使用io.Writer的代码，并不知道Socket和io.Closer的存在
func usingWriter(writer io.Writer) {
	writer.Write(nil)
}

// 使用io.Closer,并不知道Socket和io.Writer的存在
func usingCloser(closer io.Closer) {
	closer.Close()
}

// 2、多个类型可以实现相同的接口
type Service interface {
	Log(string) // 写日志方法
	Start()     //开始方法
}

type Logger struct{}

func (l *Logger) Log(s string) {
	fmt.Println("写日志:", s)
}

type GameService struct {
	Logger // 嵌入日志器
}

func (g *GameService) Start() {
	fmt.Println("游戏开始...")
}

func main() {
	s := new(Socket)
	usingWriter(s)
	usingCloser(s)

	g := new(GameService)
	g.Start()
	g.Log("发生了错误哦")

}
