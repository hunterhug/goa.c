# 包、变量和函数

## 一、举个例子

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

我们看到 `Golang` 语言只有小括号和大括号，不需要使用逗号来分隔代码，只有一种循环 `for`。

接下来我们会分析这个例子。

## 二、工程管理：包机制

每一个大型的软件工程项目，都需要进行工程管理。工程管理的一个环节就是代码层次的管理。

包，也称为库，如代码的一个包，代码的一个库，英文：`Library` 或者 `Package`。比如，我们常常听到某程序员说：嘿，X哥，我知道 `Github` 上有一个更好用的数据加密库，几千棵星呢。

在高级编程语言层次，也就是代码本身，各种语言发明了包（`package`）机制来更好的管理代码，将代码按功能分类归属于不同的包。

`Golang` 语言目前的包管理新机制叫 `go mod`。

我们的项目结构是：

```
├── diy
│   └── diy.go
└── main.go
```

每一个 `*.go` 源码文件，必须属于一个包，假设包名叫 `diy` ，在代码最顶端必须有 `package diy`，在此之前不能有其他代码片段，如 `diy/diy.go` 文件中：

```go
// 包名
package diy

// 结构体
type Diy struct {
	A int64   // 大写导出成员
	b float64 // 小写不可以导出
}
```

作为执行入口的源码，则强制包名必须为 `main`，入口函数为 `func main()`，如 `main.go` 文件中：

```go
// Golang程序入口的包名必须为 main
package main // import "golang"

// 导入其他地方的包，包通过 go mod 机制寻找
import (
	"fmt"
	"golang/diy"
)
```


在入口文件 `main.go` 文件夹下执行以下命令：

```
go mod int
```

该命令会解析 `main.go` 文件的第一行 `package main // import "golang"`，注意注释 `//` 后面的 `import "golang"`，会生成 `go.mod` 文件：

```
module golang

go 1.13
```

`Golang` 编译器会将这个项目认为是包 `golang`，这是整个项目最上层的包，而底下的文件夹 `diy` 作为 `package diy`，包名全路径就是 `golang/diy`。

接着，`main.go` 为了导入包，使用 `import ()`：

```go
// 导入其他地方的包，包通过 go mod 机制寻找
import (
	"fmt"
	"golang/diy"
)
```

可以看到导入了官方的包 `fmt` 和我们自已定义的包 `golang/diy`，官方的包会自动寻找到，不需要任何额外处理，而自己的包会在当前项目往下找。

在包 `golang/diy` 中，我们定义了一个结构体和函数：

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

结构体和函数会在后面的章节介绍，现在只需知道只有大写字母开头的结构体或函数，才能在其他包被人引用。

最后，`Golang` 的程序入口统一在包 `main` 中的 `main` 函数，执行程序时是从这里开始的：

```go
package main
import "fmt"

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

有个必须注意的事情是函数 `init()` 会在每个包被导入之前执行，如果导入了多个包，那么会根据包导入的顺序先后执行 `init()`，再回到执行函数 `main()`。


## 三、变量

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

可以看到 `var i int64`，数据类型是在变量的后面而不是前面，这是 `Golang` 语言与其他语言最大的区别之一。

同时，作为一门静态语言，`Golang` 在编译前还会检查哪些变量和包未被引用，强制禁止游离的变量和包，从而避免某些人类低级错误。如：

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

提示声明变量未使用，这是 `Golang` 语言与其他语言最大的区别之一。

变量定义后，如果没有赋值，那么存在默认值。我们也可以定义常量，只需加关键字 `const`，如：

```go
    const s  = 2
```

常量一旦定义就不能修改。

## 四、基本数据类型

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

输出：

```go
type:bool:true
type:int:3
type:float64:6
type:string:hi
type:[3]string:[1 2 3]
type:[]int64:[1 2 3]
type:map[string]int64:map[a:3 b:4]
```

数据类型基本有整数，浮点数，字符串，布尔值，数组，切片(slice) 和 字典(map) 。

1. 布尔值：`bool`。
2. 整数：`int` (默认类型，一般视操作系统位数=int32或int64)，`int32`，`int64`。
3. 浮点数：`float32`，`float64`(默认类型，更大的精度)
4. 字符：`string`。
5. 数组，切片(可变长数组)，字典(键值对结构)。

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

`Golang` 有数组类型的提供，但是一般不使用，因为数组不可变长，当你把数组大小定义好了，就再也无法变更大小。所以 `Golang` 语言造出了可变长数组：切片(`slice`)，将数组的容量大小去掉就变成了切片。切片，可以像切东西一样。自动调整大小，可以切一部分，或者把两部分拼起来。

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

切片追加一个值 `3` 进去需要使用 `append` 关键字，然后将结果再赋给自己本身，这是 `Golang` 语言与其他语言最大的区别之一，实际切片底层有个固定大小的数组，当数组容量不够时会生成一个新的更大的数组。

同时，因为日常开发中，我们经常将两个数据进行映射，类似于查字典一样，先查字母，再翻页。所以字典 `map` 开发使用频率极高，所以 `Golang` 自动提供了这一数据类型，这是 `Golang` 语言与其他语言最大的区别之一。

字典存储了一对对的键值：

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

## 五、slice 和 map 的特殊说明

键值结构字典：`map` 使用前必须初始化，如：

```go
     m := map[string]int64{}
     m1 = make(map[string]int64)
``` 

如果不对字典进行初始化，作为引用类型，它是一个 `nil` 空引用，你使用空引用，往字典里添加键值对，将会报错。

而切片结构 `slice` 不需要初始化，因为添加值时是使用 `append` 操作，内部会自动初始化，如：

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

## 六、函数

我们可以把经常使用的代码片段封装成一个函数，方便复用：

```go
// 函数，两个数相加
func sum(a, b int64) int64 {
	return a + b
}
```


`Golang` 定义函数使用的关键字是 `func`，后面带着函数名 `sum(a, b int64) int64`，表示函数 `sum` 传入两个 `int64` 整数`a` 和 `b`，输出值也是一个 `int64` 整数。

使用时：

```go
	// 定义 int64 变量
	var h, i int64 = 4, 6

	// 使用函数
	sum := sum(h, i)
	fmt.Printf("sum(h+i),h=%v,i=%v,%v\n", h, i, sum)
```

输出：

```go
sum(h+i),h=4,i=6,10
```

将函数外的变量 `h`，`i` 传入函数 `sum` 作为参数，是一个值拷贝的过程，会拷贝 `h` 和 `i` 的数据到参数 `a` 和 `b`，这两个变量是函数 `sum` 内的局部变量，两个变量相加后返回求和结果。

就算函数里面改了局部变量的值，函数外的变量还是不变的，如：

```go
package main

import "fmt"

func changeTwo(a, b int) {
	a = 6
	b = 8
}

func main() {
	a, b := 1, 2
	fmt.Println(a, b)
	changeTwo(a, b)
	fmt.Println(a, b)
}
```

输出：

```go
1 2
1 2
```

变量是有作用域的，作用域主要被约束在各级大括号 `{}` 里面，所以函数里面的变量和函数体外的变量是没有关系的，互相独立。

我们还可以实现匿名的函数如：

```go
	input := 2

	output := func(num int) int {
		num = num * 2
		return num
	}(input)

	fmt.Println(output)
```

打印出：

```go
4
```

本来函数在外部是这样的：

```go
func A(num int) int {
		num = num * 2
		return num
	}
```

现在省略了函数名，定义后直接使用：

```go
	output := func(num int) int {
		num = num * 2
		return num
	}(input)
```

`input` 是匿名函数的输入参数，匿名函数返回的值会赋予 `output`。

## 七、其他

`Golang` 会智能进行变量分析。所以有一个叫变量逃逸的说法。

也就是变量被分配到堆里面 `heap`，还是栈里面 `stack`，主要取决于后期该变量会不会继续使用，如果不会，那么就分配到栈里，如果会，分配到堆里。
