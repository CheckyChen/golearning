package main

// GOPATH 是 Go 语言中使用的一个环境变量，它使用绝对路径提供项目的工作目录

// 1、使用命令行查看GOPATH信息
// C:\Users\CHECKY>go env
// set GOARCH=amd64
// set GOBIN=
// set GOCACHE=C:\Users\CHECKY\AppData\Local\go-build
// set GOEXE=.exe
// set GOFLAGS=
// set GOHOSTARCH=amd64
// set GOHOSTOS=windows
// set GOOS=windows
// set GOPATH=C:\Users\CHECKY\go
// set GOPROXY=
// set GORACE=
// set GOROOT=D:\install\Go
// set GOTMPDIR=
// set GOTOOLDIR=D:\install\Go\pkg\tool\windows_amd64
// set GCCGO=gccgo
// set CC=gcc
// set CXX=g++
// set CGO_ENABLED=1
// set GOMOD=
// set CGO_CFLAGS=-g -O2
// set CGO_CPPFLAGS=
// set CGO_CXXFLAGS=-g -O2
// set CGO_FFLAGS=-g -O2
// set CGO_LDFLAGS=-g -O2
// set PKG_CONFIG=pkg-config
// set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\CHECKY\AppData\Local\Temp\go-build916592538=/tmp/go-build -gno-record-gcc-switches

// 2、在多项目工程中使用GOPATH

// 在很多与 Go 语言相关的书籍、文章中描述的 GOPATH 都是通过修改系统全局的环境变量来实现的。然而，根据笔者多年的 Go 语言使用和实践经验及周边朋友、同事的反馈，这种设置全局 GOPATH 的方法可能会导致当前项目错误引用了其他目录的 Go 源码文件从而造成编译输出错误的版本或编译报出一些无法理解的错误提示。
// 比如说，将某项目代码保存在 /home/davy/projectA 目录下，将该目录设置为 GOPATH。随着开发进行，需要再次获取一份工程项目的源码，此时源码保存在 /home/davy/projectB 目录下，如果此时需要编译 projectB 目录的项目，但开发者忘记设置 GOPATH 而直接使用命令行编译，则当前的 GOPATH 指向的是 /home/davy/projectA 目录，而不是开发者编译时期望的 projectB 目录。编译完成后，开发者就会将错误的工程版本发布到外网。
// 因此，建议大家无论是使用命令行或者使用集成开发环境编译 Go 源码时，GOPATH 跟随项目设定。在 Jetbrains 公司的 GoLand 集成开发环境（IDE）中的 GOPATH 设置分为全局 GOPATH 和项目 GOPATH

// 建议：建议在使用GoLand开发时只填写项目 GOPATH，每一个项目尽量只设置一个 GOPATH，不使用多个 GOPATH 和全局的 GOPATH。
