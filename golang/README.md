# 一小时入门Golang

## 前言

Golang 语言是谷歌公司在2009年正式发布并开源的高级编程语言，开源地址：[https://github.com/golang/go](https://github.com/golang/go)，官网地址：[https://golang.org](https://golang.org)。

Golang 语言语法简单，支持多平台交叉编译，支持内存 `GC`（垃圾回收），支持嵌 `C/C++` 开发，并且实现了语法层面的线程调度，开发多线程程序十分方便。语法很像 `C/Python/JavaScript` 等高级编程语言。

作为一门静态语言，在编译前会检查哪些变量和包未被引用，强制禁止游离的变量和包，从而避免某些人类低级错误。

## 正文

入门 Golang 十分简单。可以不要求你拥有其他编程语言经验，但如果已经学会 `Java/C/Python` 等计算机编程语言，你会觉得相对亲切。阅读完该文章并理解，需要一小时。

Golang 语言语法十分简单，你可以只使用函数式编程（类似 C 语言），也可以使用面向接口编程（类似面向对象语言 Java/C++）。

### 安装并简单使用

安装 [Golang：https://golang.org/dl](https://golang.org/dl)：Windows 操作系统点击 msi 按提示安装，Mac 操作系统可以使用 `brew install golang` 安装。

打开命令行终端输入：

```
go version
```

显示以下结果即为成功：

```
go version go1.13 darwin/amd64
```

在任一文件夹下新建一个文件 `main.go`（Golang 语言编写的程序文件后缀必须都为 `.go`）：

```
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

init will be before hello world
hello world
today times:2019-12-09 13:14:14.383118 +0800 CST m=+0.000199077
```

### 深入了解

编程语言离不开几种特征：

1. 支持哪些基本数据类型，如整数，浮点数，布尔值，字符串，支持哪些高级数据类型，如数组，结构体等。
2. 判断和循环语句是怎样的，是否有 `switch` 或者 `goto` 等语句。
3. 函数是怎样的，如何传参，有没有面向对象的特征等。 
4. 包管理是怎样的，如何管理一个工程，官方提供哪些标准库函数，如时间处理，字符串处理，HTTP 库，加密库等。
5. 有没有特殊的语言特征，其他语言没有的。

现在我们来看一个完整的程序：

```

```