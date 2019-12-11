# 一小时入门Golang

## 一、前言

Golang 语言是谷歌公司在2009年正式发布并开源的高级编程语言，开源地址：[https://github.com/golang/go](https://github.com/golang/go)，官网地址：[https://golang.org](https://golang.org)。

Golang 语言语法简单，支持多平台交叉编译（Linux/Mac/Windows），支持内存自动 `GC`（垃圾回收），支持嵌 `C/C++` 开发，并且实现了语法层面的线程调度，开发多线程程序十分方便。语法很像 `C/Python/JavaScript` 等高级编程语言。

入门 Golang 语言。可以不要求你拥有其他编程语言经验，但如果已经学会 `Java/C/Python` 等计算机编程语言，你会觉得相对亲切。阅读完本章并理解，需要一小时。

Golang 语言语法十分简单，你可以只使用函数式编程（类似 C 语言），也可以使用面向接口编程（类似面向对象语言 Java/C++）。

## 二、安装并简单使用

安装 [Golang：https://golang.org/dl](https://golang.org/dl)：Windows 操作系统点击 `msi` 按提示安装，Mac 操作系统可以使用 `brew install golang` 安装。

打开命令行终端输入：

```shell script
go version
```

显示以下结果即为成功：

```shell script
go version go1.13 darwin/amd64
```

在任一文件夹下新建一个文件 `main.go`（Golang 语言编写的程序文件后缀必须都为 `.go`）：

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

```shell script
go build main.go
```

编译后会在本地文件夹下生成一个二进制文件：`main` 或者 `main.exe`（Windows系统）。

执行二进制：

```shell script
./main
```

将会打印出以下结果：

```shell script
init will be before hello world
hello world
today times:2019-12-09 13:14:14.383118 +0800 CST m=+0.000199077
```

## 三、如何学习一门语言

每学一门编程语言，都离不开学习它的语言特征：

1. 支持哪些基本数据类型，如整数，浮点数，布尔值，字符串，支持哪些高级数据类型，如数组，结构体等。
2. `if` 判断和 `while` 循环语句是怎样的，是否有 `switch` 或者 `goto` 等语句。
3. 语言函数的定义是怎样的，如何传递函数参数，有没有面向对象的语言特征等。 
4. `package` 包管理是怎样的，如何管理一个工程，官方提供哪些标准库，如时间处理，字符串处理，HTTP 库，加密库等。
5. 有没有特殊的语言特征，其他语言没有的，比如某些语法糖。

## 四、从例子中学习

现在我们来建立一个完整的程序 `all.go`：

```go
// Golang程序入口的包名必须为 main
package main // import "golang"

// 导入其他地方的包，包通过 go mod 机制寻找
import (
	"fmt"
	"golang/diy"
)

// init函数在main函数之前执行
func init() {
	// 声明并初始化三个值
	var i, j, k = 1, 2, 3
	// 使用格式化包打印
	fmt.Println("init hello world")
	fmt.Println(i, j, k)
}

// 函数，两个数相加
func sum(a, b int64) int64 {
	return a + b
}

// 程序入口必须为 main 函数
func main() {
	// 未使用的变量，不允许声明
	//cannot := 6

	fmt.Println("hello world")

	// 定义基本数据类型
	a := 3                                // int
	b := 6.0                              // float64
	c := "hi"                             // string
	d := [3]int64{1, 2, 3}                // array，基本不用到
	e := []int64{1, 2, 3}                 // slice
	f := map[string]int64{"a": 3, "b": 4} // map
	fmt.Printf("type:%T:%v\n", a, a)
	fmt.Printf("type:%T:%v\n", b, b)
	fmt.Printf("type:%T:%v\n", c, c)
	fmt.Printf("type:%T:%v\n", d, d)
	fmt.Printf("type:%T:%v\n", e, e)
	fmt.Printf("type:%T:%v\n", f, f)

	// 切片增加值
	e = append(e, 3)

	// 增加map键值
	f["f"] = 5

	// 查找map键值
	v, ok := f["f"]
	fmt.Println(v, ok)
	v, ok = f["ff"]
	fmt.Println(v, ok)

	// 判断语句
	if a > 0 {
		fmt.Println("a>0")
	} else {
		fmt.Println("a<=0")
	}

	// 循环语句
	for {
		if true {
			fmt.Println("for")
			// 退出循环
			break
		}
	}

	// 循环语句
	for i := 9; i <= 10; i++ {
		fmt.Printf("i=%d\n", i)
	}

	// 循环切片
	for k, v := range e {
		fmt.Println(k, v)
	}

	// 循环map
	for k, v := range f {
		fmt.Println(k, v)
	}

	// 定义 int64 变量
	var h, i int64 = 4, 6

	// 使用函数
	sum := sum(h, i)
	fmt.Printf("sum(h+i),h=%v,i=%v,%v\n", h, i, sum)

	// 新建结构体，值
	g := diy.Diy{
		A: 2,
		//b: 4.0, // 小写成员不能导出
	}

	// 打印类型，值
	fmt.Printf("type:%T:%v\n", g, g)

	// 小写方法不能导出
	//g.set(1,1)
	g.Set(1, 1)
	fmt.Printf("type:%T:%v\n", g, g) // 结构体值变化

	g.Set2(3, 3)
	fmt.Printf("type:%T:%v\n", g, g) // 结构体值未变化

	// 新建结构体，引用
	k := &diy.Diy{
		A: 2,
	}
	fmt.Printf("type:%T:%v\n", k, k)
	k.Set(1, 1)
	fmt.Printf("type:%T:%v\n", k, k) // 结构体值变化
	k.Set2(3, 3)
	fmt.Printf("type:%T:%v\n", k, k) // 结构体值未变化

	// 新建结构体，引用
	m := new(diy.Diy)
	m.A = 2
	fmt.Printf("type:%T:%v\n", m, m)
}
```

在相同目录下新建 `diy` 文件夹，文件下新建一个 `diy.go` 文件（名字任取）：

```go
// 包名
package diy

// 结构体
type Diy struct {
	A int64   // 大写导出成员
	b float64 // 小写不可以导出
}

// 引用结构体的方法，引用传递，会改变原有结构体的值
func (diy *Diy) Set(a int64, b float64) {
	diy.A = a
	diy.b = b
	return
}

// 值结构体的方法，值传递，不会改变原有结构体的值
func (diy Diy) Set2(a int64, b float64) {
	diy.A = a
	diy.b = b
	return
}

// 小写方法，不能导出
func (diy Diy) set(a int64, b float64) {
	diy.A = a
	diy.b = b
	return
}

// 小写函数，不能导出，只能在同一包下使用
func sum(a, b int64) int64 {
	return a + b
}
```

进入文件所在目录，打开命令行终端，执行：

```
go mod init
go run main.go


init hello world
1 2 3
hello world
type:int:3
type:float64:6
type:string:hi
type:[3]int64:[1 2 3]
type:[]int64:[1 2 3]
type:map[string]int64:map[a:3 b:4]
5 true
0 false
a>0
for
i=9
i=10
0 1
1 2
2 3
3 3
a 3
b 4
f 5
sum(h+i),h=4,i=6,10
type:diy.Diy:{2 0}
type:diy.Diy:{1 1}
type:diy.Diy:{1 1}
type:*diy.Diy:&{2 0}
type:*diy.Diy:&{1 1}
type:*diy.Diy:&{1 1}
type:*diy.Diy:&{2 0}
```

会显示一些打印结果，我们在接下来会分析这个栗子。

### 4.1 工程管理：包机制

每一个大型的软件工程项目，都需要进行工程管理。工程管理的一个环节就是代码层次的管理。

包，也称为库，如代码的一个包，代码的一个库，英文：`library` 或者 `package`。

比如，我们常常听到某程序员说：嘿，X哥，我知道 `Github` 上有一个更好用的数据加密库，几千颗星呢。

在高级编程语言层次，也就是代码本身，各种语言发明了包（`package`）机制来更好的管理代码，将代码按功能分类归属于不同的包。

Golang 语言目前的包管理新机制叫 `go mod`，之前的老机制 `GOPATH` 方式可参考该文章：[Golang高阶：Golang1.5到Golang1.12包管理](https://www.lenggirl.com/language/gomod.html)。

我们在例子代码的文件夹下，打开命令终端执行：

```
go mod int
```

该命令将解析 `package main // import "golang"`，将项目的入口定义

### 4.2 基本数据类型和变量

作为一门静态语言，Golang 在编译前会检查哪些变量和包未被引用，强制禁止游离的变量和包，从而避免某些人类低级错误。如：

```golang
package main

func main(){
    a := 2
}
```

这样执行：

```shell script
go run main.go

./all.go:26:2: cannot declared and not used
```

提示声明变量未使用。