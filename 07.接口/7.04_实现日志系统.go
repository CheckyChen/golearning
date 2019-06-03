package main

import (
	"errors"
	"fmt"
	"os"
)

// 日志可以用于查看和分析应用程序的运行状态。日志一般可以支持输出多种形式，如命令行、文件、网络等
// 1、日志对外接口
//
//本例中定义一个日志写入器接口（LogWriter），要求写入设备必须遵守这个接口协议才能被日志器（Logger）注册。日志器有一个写入器的注册方法（Logger 的 RegisterWriter() 方法）。
//
//日志器还有一个 Log() 方法，进行日志的输出，这个函数会将日志写入到所有已经注册的日志写入器（LogWriter）中，详细代码实现请参考下面的代码。

// 声明日志写入器接口
type LogWriter interface {
	Write(data interface{}) error
}

// 日志器
type Logger struct {
	// 这个日志器用到的日志写入器
	writerList []LogWriter
}

// 注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

// 将一个data类型的数据写入日志
func (l *Logger) Log(data interface{}) {
	// 遍历所有注册的写入器
	for _, writer := range l.writerList {
		writer.Write(data)
	}
}

// 创建日志器的实例
func NewLogger() *Logger {
	return &Logger{}
}

// *********文件写入器**************

// 声明一个文件写入器
type fileWriter struct {
	file *os.File
}

// 设置文件写入器写入文件名
func (f *fileWriter) SetFile(fileName string) (err error) {
	// 如果当前文件已经打开，关闭当前一个文件
	if f.file != nil {
		f.file.Close()
	}
	// 创建一个文件，并保存文件句柄
	f.file, err = os.Create(fileName)
	// 如果创建过程出现错误，则返回错误
	return err
}

func (f *fileWriter) Write(data interface{}) error {

	if f.file == nil {
		return errors.New("没有找到文件")
	}
	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)
	// 将数据以字节数组写入文件中
	_, err := f.file.Write([]byte(str))
	return err
}

func newFileWriter() *fileWriter {
	return &fileWriter{}
}

//*********命令行写入器***************
type consoleWriter struct{}

// 实现LogWriter的Write()方法
func (c *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

// 创建命令行写入器实例
func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}

// 2、使用日志
func createLogger() *Logger {
	// 创建日志器
	l := NewLogger()
	// 创建一个命令行写入器
	cw := newConsoleWriter()
	// 将命令行写入器注册到日志器中
	l.RegisterWriter(cw)
	// 创建一个文件写入器
	fw := newFileWriter()
	// 设置文件名
	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}
	// 注册文件写入器到日志器中
	l.RegisterWriter(fw)
	return l
}

func main() {
	l := createLogger()
	l.Log("这是一行日志")
}
