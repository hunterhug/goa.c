# 简单入门Golang

我们只学 `Golang` 语言的一个子集，足以开展接下来数据结构和算法的实现即可。

## 一、前言

`Golang` 语言是谷歌 `Google` 公司在2007年启动，并在2009年正式发布并 `开源` 的高级编程语言。开源地址：[https://github.com/golang/go](https://github.com/golang/go)，官网地址：[https://golang.org](https://golang.org)。

`Golang` 语言语法简单，支持多平台交叉编译（Linux/Mac/Windows），支持内存自动 `GC`（垃圾回收），支持嵌 `C/C++` 开发，并且实现了语法层面的线程调度，开发多线程程序十分方便。语法很像 `C/Python/JavaScript` 等高级编程语言。

设计这门语言的设计者有以下几位：

1. `Ken Thompson`：在贝尔实验室与 `Dennis M. Ritche` 发明了 `C` 语言和 `Unix` 操作系统，与 `Rob Pike` 发明了 `UTF-8` 编码，图灵奖得主。
2. `Rob Pike`：也参与开发了 `Unix` 操作系统，`UTF-8` 编码发明者之一。
3. `Robert Griesemer`：参与过 `V8 JavaScript` 引擎和 `Java HotSpot` 虚拟机的研发。

前两位比较知名，现在都已经退休了，其他人有兴趣可以谷歌一下。

## 二、安装并简单使用

安装 [Golang：https://golang.org/dl](https://golang.org/dl)：Windows 操作系统点击 `msi` 按提示安装，Mac 操作系统可以使用 `brew install golang` 安装。

打开命令行终端输入：

```
go version
```

显示以下结果即为成功：

```
go version go1.13 darwin/amd64
```

在任一文件夹下新建一个文件 `main.go`（`Golang` 语言编写的程序文件后缀必须都为 `.go`）：

```go
package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("init will be before hello world")
}

func main() {
	fmt.Println("hello world")
	fmt.Println("today times:" + time.Now().String())
}
```

打开命令行终端进行编译：

```
go build main.go
```

编译后会在本地文件夹下生成一个二进制文件：`main` 或者 `main.exe`（Windows系统）。

执行二进制：

```
./main
```

将会打印出以下结果：

```
init will be before hello world
hello world
today times:2019-12-09 13:14:14.383118 +0800 CST m=+0.000199077
```

## 三、如何学习一门语言

每学一门编程语言，都离不开学习它的语言特征：

1. 支持哪些 `基本数据类型`，如整数，浮点数，布尔值，字符串，支持哪些高级数据类型，如数组，结构体等。
2. `if` 判断和 `while` 循环语句是怎样的，是否有 `switch` 或者 `goto` 等语句。
3. 语言 `函数` 的定义是怎样的，如何传递函数参数，有没有 `面向对象` 的语言特征等。 
4. `package` 包管理是怎样的，如何管理一个工程，官方提供哪些标准库，如时间处理，字符串处理，HTTP 库，加密库等。
5. 有没有特殊的语言特征，其他语言没有的，比如某些语法糖。

如果迫不及待想学习 `Golang`，可以安装 `Docker` 后，打开终端执行：

```
# 拉镜像
docker pull hunterhug/gotourzh

# 后台运行
docker run -d -p 9999:9999 hunterhug/gotourzh
```

浏览器输入：[http://127.0.0.1:9999](http://127.0.0.1:9999) 更全面地学习。
