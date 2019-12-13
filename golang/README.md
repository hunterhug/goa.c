# 一小时入门Golang

我们只学 Golang 语言的一个子集，足以开展接下来数据结构和算法的实现即可。如果想更全面的学习，请安装 docker 后，打开终端输入：

```
# 拉镜像
docker pull hunterhug/gotourzh

# 后台运行
docker run -d -p 9999:9999 hunterhug/gotourzh
```

打开浏览器输入：[127.0.0.1:9999](http://127.0.0.1:9999) ，即可学习。

## 一、前言

Golang 语言是谷歌公司在2009年正式发布并开源的高级编程语言，开源地址：[https://github.com/golang/go](https://github.com/golang/go)，官网地址：[https://golang.org](https://golang.org)。

Golang 语言语法简单，支持多平台交叉编译（Linux/Mac/Windows），支持内存自动 `GC`（垃圾回收），支持嵌 `C/C++` 开发，并且实现了语法层面的线程调度，开发多线程程序十分方便。语法很像 `C/Python/JavaScript` 等高级编程语言。

入门 Golang 语言。可以不要求你拥有其他编程语言经验，但如果已经学会 `Java/C/Python` 等计算机编程语言，你会觉得相对亲切。阅读完本章并理解，需要一小时。

Golang 语言语法十分简单，你可以只使用函数式编程（类似 C 语言），也可以使用面向接口编程（类似面向对象语言 Java/C++）。

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

```
go build main.go
```

编译后会在本地文件夹下生成一个二进制文件：`main` 或者 `main.exe`（Windows系统）。

执行二进制：

```
./main
```

将会打印出以下结果：

```go
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

现在我们来建立一个完整的程序 `main.go`：

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
	p := true                             // bool
	a := 3                                // int
	b := 6.0                              // float64
	c := "hi"                             // string
	d := [3]string{"1", "2", "3"}         // array，基本不用到
	e := []int64{1, 2, 3}                 // slice
	f := map[string]int64{"a": 3, "b": 4} // map
	fmt.Printf("type:%T:%v\n", p, p)
	fmt.Printf("type:%T:%v\n", a, a)
	fmt.Printf("type:%T:%v\n", b, b)
	fmt.Printf("type:%T:%v\n", c, c)
	fmt.Printf("type:%T:%v\n", d, d)
	fmt.Printf("type:%T:%v\n", e, e)
	fmt.Printf("type:%T:%v\n", f, f)

	// 切片放值
	e[0] = 9
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

	// 死循环语句
	a = 0
	for {
		if a >= 10 {
			fmt.Println("out")
			// 退出循环
			break
		}

		a = a + 1
		if a > 5 {
			continue
		} else {
			fmt.Println(a)
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

	s := make([]int64, 5)
	s1 := make([]int64, 0, 5)
	m1 := make(map[string]int64, 5)
	m2 := make(map[string]int64)
	fmt.Printf("%#v,cap:%#v,len:%#v\n", s, cap(s), len(s))
	fmt.Printf("%#v,cap:%#v,len:%#v\n", s1, cap(s1), len(s1))
	fmt.Printf("%#v,len:%#v\n", m1, len(m1))
	fmt.Printf("%#v,len:%#v\n", m2, len(m2))

	var ll []int64
	fmt.Printf("%#v\n", ll)
	ll = append(ll, 1)
	fmt.Printf("%#v\n", ll)
	ll = append(ll, 2, 3, 4, 5, 6)
	fmt.Printf("%#v\n", ll)
	ll = append(ll, []int64{7, 8, 9}...)
	fmt.Printf("%#v\n", ll)

	fmt.Println(ll[0:2])
	fmt.Println(ll[:2])
	fmt.Println(ll[0:])
	fmt.Println(ll[:])
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
```

会显示一些打印结果：

```
init hello world
1 2 3
hello world
type:bool:true
type:int:3
type:float64:6
type:string:hi
type:[3]string:[1 2 3]
type:[]int64:[1 2 3]
type:map[string]int64:map[a:3 b:4]
5 true
0 false
a>0
0
1
2
3
4
5
^Csignal: interrupt
zhujiangdeMac-mini:example2 sachsen$ go run main.go 
init hello world
1 2 3
hello world
type:bool:true
type:int:3
type:float64:6
type:string:hi
type:[3]string:[1 2 3]
type:[]int64:[1 2 3]
type:map[string]int64:map[a:3 b:4]
5 true
0 false
a>0
1
2
3
4
5
out
i=9
i=10
0 9
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
[]int64{0, 0, 0, 0, 0},cap:5,len:5
[]int64{},cap:5,len:0
map[string]int64{},len:0
map[string]int64{},len:0
[]int64(nil)
[]int64{1}
[]int64{1, 2, 3, 4, 5, 6}
[]int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
[1 2]
[1 2]
[1 2 3 4 5 6 7 8 9]
[1 2 3 4 5 6 7 8 9]
```

我们看到 Golang 语言只有小括号和大括号，不需要使用逗号来分隔代码，只有一种循环 `for`，接下来我们会分析这个栗子。

### 4.1 工程管理：包机制

每一个大型的软件工程项目，都需要进行工程管理。工程管理的一个环节就是代码层次的管理。

包，也称为库，如代码的一个包，代码的一个库，英文：`Library` 或者 `Package`。

比如，我们常常听到某程序员说：嘿，X哥，我知道 `Github` 上有一个更好用的数据加密库，几千颗星呢。

在高级编程语言层次，也就是代码本身，各种语言发明了包（`package`）机制来更好的管理代码，将代码按功能分类归属于不同的包。

Golang 语言目前的包管理新机制叫 `go mod`，之前的老机制 `GOPATH` 方式可参考该文章：[Golang高阶：Golang1.5到Golang1.12包管理](https://www.lenggirl.com/language/gomod.html)。

每一个 `*.go` 源码文件，必须属于一个包，假设包名叫 `diy` ，在代码最顶端必须有 `package diy`，在此之前不能有其他代码片段。作为执行入口的源码，则强制包名必须为 `main`，入口函数为 `func main()`。

我们的项目结构是：

```
├── diy
│   └── diy.go
└── main.go
```

在入口文件 `main.go` 文件夹下执行：

```
go mod int
```

该命令将解析 `main.go` 文件 `package main // import "golang"`，会生成 `go.mod` 文件：

```
module golang

go 1.13
```

这样 Golang 编译器将会把这个项目认为是包 `golang`，是最上层的包，而底下的文件夹 `diy` 作为 `package diy`，包名全路径就是 `golang/diy`。

接着，`main.go` 为了导入包，使用 `import`：

```go
// 导入其他地方的包，包通过 go mod 机制寻找
import (
	"fmt"
	"golang/diy"
)
```

导入了官方的包 `fmt` 和我们自已定义的包 `golang/diy`，官方的包会自动寻找到，不需要任何额外处理。

在包 `diy` 中，我们定义了一个结构体和函数：

```go
// 结构体
type Diy struct {
	A int64   // 大写导出成员
	b float64 // 小写不可以导出
}

// 小写函数，不能导出，只能在同一包下使用
func sum(a, b int64) int64 {
	return a + b
}
```

对于包中小写的函数或者结构体中小写的字段，不能导出，其他包不能使用它，`Golang` 用它实现了私有或公有控制，毕竟有些包的内容我们不想在其他包中被使用，类似 `Java` 的 `private` 关键字。

Golang 的程序入口统一在包 `main` 中的 `main` 函数：

```go
// init函数在main函数之前执行
func init() {
	// 声明并初始化三个值
	var i, j, k = 1, 2, 3
	// 使用格式化包打印
	fmt.Println("init hello world")
	fmt.Println(i, j, k)
}

// 程序入口必须为 main 函数
func main() {
}
```

函数 `init` 会在每个包初始化的时候执行，然后再执行入口函数 `main`。

### 4.2 变量

Golang语言可以先声明变量，再赋值，也可以直接创建一个带值的变量。如：


```go
// 声明并初始化三个值
var i, j, k = 1, 2, 3

// 声明后再赋值
var i int64
i = 3

// 直接赋值，创建一个新的变量
j := 5
```

可以看到 `var i int64`，数据类型是在变量的后面而不是前面，这是 Golang 语言与其他语言最大的区别之一。

同时，作为一门静态语言，Golang 在编译前还会检查哪些变量和包未被引用，强制禁止游离的变量和包，从而避免某些人类低级错误。如：

```go
package main

func main(){
    a := 2
}
```

如果执行将会报错：

```
go run main.go

./main.go:26:2: cannot declared and not used
```

提示声明变量未使用，这是 Golang 语言与其他语言最大的区别之一。

变量定义后，如果没有赋值，那么存在默认值。变量有作用范围，称为作用域，在每一对大括号创建的变量，只在该大括号代码片段里起作用，比如函数内定义的变量无法在函数外使用。


### 4.3 基本数据类型

我们再来看看基本的数据类型有那些：

```go
	// 定义基本数据类型
	p := true                             // bool
	a := 3                                // int
	b := 6.0                              // float64
	c := "hi"                             // string
	d := [3]string{"1", "2", "3"}         // array，基本不用到
	e := []int64{1, 2, 3}                 // slice
	f := map[string]int64{"a": 3, "b": 4} // map
	fmt.Printf("type:%T:%v\n", p, p)
	fmt.Printf("type:%T:%v\n", a, a)
	fmt.Printf("type:%T:%v\n", b, b)
	fmt.Printf("type:%T:%v\n", c, c)
	fmt.Printf("type:%T:%v\n", d, d)
	fmt.Printf("type:%T:%v\n", e, e)
	fmt.Printf("type:%T:%v\n", f, f)
```


数据类型基本有整数，浮点数，字符串，布尔值，数组，切片和 map 。

布尔值：`bool`，整数：`int` (默认类型，一般视操作系统位数=int32或int64)，`int32`，`int64`。浮点数：`float32`，`float64`(默认类型，更大的精度)，字符：`string`。

没声明具体变量类型的时候，会自动识别类型，把整数认为是 `int` 类型，把带小数点的认为是 `float64` 类型，如：

```go
	a := 3                                // int
	b := 6.0                              // float64
```

所以当你需要使用确切的 `int64` 或 `float32` 类型时，你需要这么做：

```go
    var a int64 = 3
    var b float32 = 6.0
```

Golang 有数组类型的提供，但是一般不使用，因为数组不可变长，当你把数组大小定义好了，就再也无法变更大小。所以 Golang 语言造出了可变长数组：切片(`slice`)，将数组的容量大小去掉就变成了切片。

```go
	d := [3]string{"1", "2", "3"}         // array，基本不用到
	e := []int64{1, 2, 3}                 // slice
```

切片可以像数组一样按下标取值，放值，也可以追加值：

```go
	// 切片放值
	e[0] = 9
	// 切片增加值
	e = append(e, 3)
```

切片追加一个值 `3` ß进去需要使用 `append` 关键字，然后将结果再赋给自己本身，这是 Golang 语言与其他语言最大的区别之一。

同时，因为键值对 `map` 开发使用频率极高，所以 `Golang` 自动提供了这一数据类型，这是 Golang 语言与其他语言最大的区别之一。：

```go
	// 增加map键值
	f["f"] = 5

	// 查找map键值
	v, ok := f["f"]
	fmt.Println(v, ok)
	v, ok = f["ff"]
	fmt.Println(v, ok)
```

结构如 `map[string]int64` 表示键为字符串 `string`，值为整数 `int64`，然后你可以将 `f = 5` 这种关系进行绑定，需要时可以拿出键 `f` 对应的值。

### 4.4 判断和循环语言

Golang 只有一种判断和一种循环：`if` 和 `for`。

判断语句如：

```go
	// 判断语句
	if a > 0 {
		fmt.Println("a>0")
	} else {
		fmt.Println("a<=0")
	}
```

判断 `a > 0` 条件不需要加小括号。

循环语句：

```go
	// 循环语句
	for i := 9; i <= 10; i++ {
		fmt.Printf("i=%d\n", i)
	}
```

其中 `i` 是局部变量，循环第一次前被赋予了值 `9`，然后判断是否满足 `i<=10` 条件，如果满足那么进入循环打印，每一次循环后会加`1`，也就是 `i++`，然后继续判断是否满足条件。

你也可以死循环：

```go
	// 死循环语句
	a = 0
	for {
		if a >= 10 {
			fmt.Println("out")
			// 退出循环
			break
		}

		a = a + 1
		if a > 5 {
			continue
		} 
		
        fmt.Println(a)
	}
```

死循环直接 `for`，后面不需要加条件，然后当 `a>=10` 时跳出循环可以使用 `break`，表示跳出 `for {}`，对于 `a > 5`，我们不想打印出值，可以使用 `continue` 跳过后面的语句 `fmt.Println(a)`，提前再一次进入循环。

切片和 `map` 都可以使用循环来遍历数据:

```go
	// 循环切片
	for k, v := range e {
		fmt.Println(k, v)
	}

	// 循环map
	for k, v := range f {
		fmt.Println(k, v)
	}
```

切片遍历出来的结果为：数据下标，数据，`map`遍历出来的结果为：数据的键，数据的值:

```
0 1
1 2
2 3
3 3

a 3
b 4
f 5
```

### 4.5 函数

我们可以把经常使用的代码片段封装成一个函数，方便复用：

```go
// 函数，两个数相加
func sum(a, b int64) int64 {
	return a + b
}
```


Golang 定义函数使用的关键字是 `func`，后面带着函数名 `sum(a, b int64) int64`，表示函数 `sum` 传入两个 `int64` 整数`a` 和 `b`，输出值也是一个 `int64` 整数。

使用时：

```go
	// 定义 int64 变量
	var h, i int64 = 4, 6

	// 使用函数
	sum := sum(h, i)
	fmt.Printf("sum(h+i),h=%v,i=%v,%v\n", h, i, sum)
```

将函数外的变量 `h`，`i` 传入函数 `sum` 作为参数，是一个值拷贝的过程，会拷贝 `h` 和 `i` 的数据到参数 `a` 和 `b`，这两个变量是函数 `sum` 内的局部变量，和函数外的变量没有关系。

### 4.5 结构体

有了基本的数据类型，还远远不够，所以 Golang 支持我们定义自己的数据类型，结构体：

```go
// 结构体
type Diy struct {
	A int64   // 大写导出成员
	b float64 // 小写不可以导出
}
```

结构体的名字为 `Diy`，使用 `type 结构体名字 struct` 来定义。

结构体里面有一些成员 `A` 和 `b` ，和变量定义一样，类型 `int64 ` 和 `float64` 放在后面，不需要任何符号分隔，只需要换行即可。结构体里面小写的成员，在包外无法使用，也就是不可导出。

使用结构体时：

```go
    // 新建结构体，值
    g := diy.Diy{
        A: 2,
        //b: 4.0, // 小写成员不能导出
    }

    // 新建结构体，引用
    k := &diy.Diy{
        A: 2,
    }

    // 新建结构体，引用
    m := new(diy.Diy)
    m.A = 2
```

可以按照基本数据类型的样子使用结构体，上述创立的 `g` 是一个值类型的结构体，你也可以使用结构体值前面加一个 `&` 或者使用`new`来创建一个引用类型的结构体，如 `k` 和 `m` 。

引用和值类型有何区别的，我们知道传参数进函数的时候，参数是值拷贝，函数里的变量是局部变量，就算修改了函数里传入的变量，函数外也发现不了。

但引用类型的变量，传入函数时，虽然也是传值，但这个时候传的值是一个引用，这个引用指向了函数外的结构体，使用这个引用在函数里修改结构体的值，外面函数也会发现。

如果传入的不是引用类型的结构体，而是值类型的结构体，那么会拷贝一份结构体，该结构体和原来的结构体就没有关系了。

结构体可以和函数绑定，也就是说这个函数只能被该结构体使用，这种函数称为方法：

```go
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
```

只不过在以前函数的基础上 `func Set(a int64, b float64)`，变成了 `func (diy *Diy) Set(a int64, b float64)`，然后在函数里面，可以使用结构体变量 `diy` 里面的成员。

上面表示值类型的结构体 `diy Diy` 可以使用 `Set2` 方法，引用类型的结构体 `diy *Diy` 可以使用 `Set` 方法。

如果是这样的话，我们每次使用结构体方法时，都要注意结构体是值还是引用类型，幸运的是 Golang 操碎了心，每次使用一个结构体调用方法，都会自动将结构体进行类型转换，以适配方法。比如下面：

```go
	// 新建结构体，值
	g := diy.Diy{
		A: 2,
		//b: 4.0, // 小写成员不能导出
	}

	g.Set(1, 1)
	fmt.Printf("type:%T:%v\n", g, g) // 结构体值变化

	g.Set2(3, 3)
	fmt.Printf("type:%T:%v\n", g, g) // 结构体值未变化

	// 新建结构体，引用
	k := &diy.Diy{
		A: 2,
	}
	k.Set(1, 1)
	fmt.Printf("type:%T:%v\n", k, k) // 结构体值变化
	k.Set2(3, 3)
	fmt.Printf("type:%T:%v\n", k, k) // 结构体值未变化
```

结构体 `g` 是值类型，本来不能调用 `Set` 方法，但是 Golang 帮忙转换了，我们毫无感知，然后值类型就变成了引用类型。同理，`k` 是引用类型，照样可以使用 `Set2` 方法。

前面我们也说过，函数传入引用，函数里修改该引用对应的值，函数外也会发现。结构体的方法也是一样，不过范围扩散了结构体本身，方法里可以修稿结构体本身，但是如果结构体是值，那么修改后，外面的世界是发现不了的。

切片，`map` 两种基本类型，都是引用类型，所以进行传递调用函数或方法时，要协调好使用方法。


### 4.6 关键字 new 和 make

关键字 `new` 主要用来创建一个引用类型的结构体，只有结构体可以用。

关键字 `make` 是用来创建和初始化一个切片或者 `map`。我们可以直接赋值来使用：

```go
	e := []int64{1, 2, 3}                 // slice
	f := map[string]int64{"a": 3, "b": 4} // map
```

但是这种直接赋值相对粗暴，因为我们使用时可能不知道数据在哪里，数据有多少。

所以，我们在创建切片和 `map` 时，可以指定容量大小。看示例：

```go
	s := make([]int64, 5)
	s1 := make([]int64, 0, 5)
	m1 := make(map[string]int64, 5)
	m2 := make(map[string]int64)
	fmt.Printf("%#v,cap:%#v,len:%#v\n", s, cap(s), len(s))
	fmt.Printf("%#v,cap:%#v,len:%#v\n", s1, cap(s1), len(s1))
	fmt.Printf("%#v,len:%#v\n", m1, len(m1))
	fmt.Printf("%#v,len:%#v\n", m2, len(m2))
```

运行后：

```go
[]int64{0, 0, 0, 0, 0},cap:5,len:5
[]int64{},cap:5,len:0
map[string]int64{},len:0
map[string]int64{},len:0

```

切片可以使用 `make([],占用容量大小，全部容量大小)` 来定义，你可以创建一个容量大小为 `5`，但是实际占用容量为 `0` 的切片，比如 `make([]int64, 0, 5)`，你预留了 `5` 个空间，这样当你切片 `append` 时，不会因为容量不足而内部去分配空间，节省了时间。

如果你省略了后面的参数如 `make([]int64, 5)`，那么其等于 `make([]int64, 5，5)`，因为这时全部容量大小就等于占用容量大小。内置语言 `cap` 和 `len` 可以查看全部容量大小，已经占用的容量大小。

同理，`map` 也可以指定容量，使用 `make([],容量大小)`，但是它没有所谓的占用容量，它去掉了这个特征，因为我们使用切片，可能需要五个空白的初始值，但是 `map` 没有键的情况下，预留初始值也没作用。省略容量大小，表示创建一个容量为 `0` 的键值结构，当赋值时会自动分配空间。

### 4.7 slice 和 map 的特殊说明

键值结构 `map` 使用前必须初始化，如：

```go
 m := map[string]int64{}
 m1 = make(map[string]int64)
``` 

如果不进行初始化，作为引用类型，它是一个 `nil` 空引用，你使用空引用，往 `map` 里赋值，将会报错。


切片结构 `slice` 不需要初始化，因为添加值时是使用 `append` 操作，内部会自动初始化，如：

```go
	var ll []int64
	fmt.Printf("%#v\n", ll)
	ll = append(ll, 1)
	fmt.Printf("%#v\n", ll)
```

打印：

```
[]int64(nil)
[]int64{1}
```

同时切片有以下特征：


```go
	ll = append(ll, 2, 3, 4, 5, 6)
	fmt.Printf("%#v\n", ll)
	ll = append(ll, []int64{7, 8, 9}...)
	fmt.Printf("%#v\n", ll)

	fmt.Println(ll[0:2])
	fmt.Println(ll[:2])
	fmt.Println(ll[0:])
	fmt.Println(ll[:])
```

内置语法 `append` 可以传入多个值，将多个值追加进切片。并且可以将另外一个切片，如 `[]int64{7, 8, 9}...`，用三个点表示遍历出里面的值，把一个切片中的值追加进另外一个切片。

在切片后面加三个点 `...` 表示虚拟的创建若干变量，将切片里面的值赋予这些变量，再将变量传入函数。

我们取切片的值，除了可以通过下标取一个值，也可以取范围：`[下标起始:下标截止(不包括取该下标的值)]`，如 `[0:2]`，表示取出下标为 `0和1` 的值，总共有两个值，再比如 `[0:4]`，表示取出下标为 `0，1，2，3` 的值。如果下标取值，下标超出实际容量，将会报错。

如果下标起始等于下标 `0`，那么可以省略，如 `[:2]`，如果下标截止省略，如 `[2:]` 表示从下标 `2` 开始，取后面所有的值。这个表示 `[:]` 本身没有作用，它就表示切片本身。

### 4.8 内置语法和函数，方法的区别

函数是代码片段的一个封装，方法是将函数和结构体绑定。

但是 Golang 里面有一些内置语法，不是函数，也不是方法，比如 `append`，`cap`，`len`，`make`，这是一种语法特征。

语法特征是高级语言提供的，内部帮你隐藏了如何分配内存等细节。

### 4.9 使用库函数

官方提供了很多库给我们用，是封装好的轮子，比如包 `fmt`，我们多次使用它来打印数据。

我们可以查看到其里面的实现：

```go
package fmt

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```

函数 `Printf` 使用到了另外一个函数，而函数又使用了另外的函数。

我们在某些时候，可以使用官方库或别人写的库，毕竟轮子重造需要时间。

同时，如果想开发速度提高，建议安装 IDE(集成开发环境)，如 `Goland`，请自行百度。

## 五、总结

我们只学习了 Golang 语言的一个子集，很多关键特征如 `interface` 还有 `go func()` 以及 `chan` 都没有讲解，请参照最开始的步骤自行学习。

后面的算法分析和实现，会使用 Golang 来举例。