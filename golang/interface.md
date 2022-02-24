# 接口

在 `Golang` 世界中，有一种叫 `interface` 的东西，很是神奇。

## 一、数据类型 interface{}

如果你事前并不知道变量是哪种数据类型，不知道它是整数还是字符串，但是你还是想要使用它。

`Golang` 就产生了名为 `interface{}` 的数据类型，表示并不知道它是什么类型。举例子：

````go
package main

import (
	"fmt"
	"reflect"
)

func print(i interface{}) {
	fmt.Println(i)
}

func main() {
	// 声明一个未知类型的 a，表明不知道是什么类型
	var a interface{}
	a = 2
	fmt.Printf("%T,%v\n", a, a)

	// 传入函数
	print(a)
	print(3)
	print("i love you")

	// 使用断言，判断是否是 int 数据类型
	v, ok := a.(int)
	if ok {
		fmt.Printf("a is int type,value is %d\n", v)
	}

	// 使用断言，判断变量类型
	switch a.(type) {
	case int:
		fmt.Println("a is type int")
	case string:
		fmt.Println("a is type string")
	default:
		fmt.Println("a not type found type")
	}

	// 使用反射找出变量类型
	t := reflect.TypeOf(a)
	fmt.Printf("a is type: %s", t.Name())
}
````

输出：

```go
int,2
2
3
i love you
a is int type,value is 2
a is type int
a is type: int
```

### 1.1.基本使用

我们使用 `interface{}`，可以声明一个未知类型的变量 `a`：

```go
	// 声明一个未知类型的 a，表明不知道是什么类型
	var a interface{}
	a = 2
	fmt.Printf("%T,%v\n", a, a)
```

然后给变量赋值一个整数：`a=2`，这时 `a` 仍然是未知类型，使用占位符 `%T` 可以打印变量的真实类型，占位符 `%v` 打印值，这时 `fmt.Printf` 在内部会进行类型判断。

我们也可以将函数的参数也定为 `interface`，和变量的定义一样：

```go
func print(i interface{}) {
	fmt.Println(i)
}
```

使用时：

```go
	// 传入函数
	print(a)
	print(3)
	print("i love you")
```

我们传入 `print` 函数的参数可以是任何类型，如整数 `3` 或字符串 `i love you` 等。进入函数后，函数内变量 `i` 丢失了类型，是一个未知类型，这种特征使得我们如果想处理不同类型的数据，不需要写多个函数。

当然，结构体里面的字段也可以是 `interface{}`：

```go
type H struct {
    A interface{}
    B interface{}
}
```

### 1.2.判断具体类型

我们定义了 `interface{}`，但是实际使用时，我们有判断类型的需求。有两种方法可以进行判断。

使用断言：

````go
	// 使用断言，判断是否是 int 数据类型
	v, ok := a.(int)
	if ok {
		fmt.Printf("a is int type,value is %d\n", v)
	}
````

直接在变量后面使用 `.(int)`，有两个返回值 `v, ok` 会返回。`ok` 如果是 `true` 表明确实是整数类型，这个整数会被赋予 `v`，然后我们可以拿 `v` 愉快地玩耍了。否则，`ok` 为 `false`，`v` 为空值，也就是默认值 0。

如果我们每次都这样使用，会很难受，因为一个 `interface{}` 类型的变量，数据类型可能是 `.(int)`，可能是 `.(string)`，可以使用 `switch` 来简化：

```go
	// 使用断言，判断变量类型
	switch a.(type) {
	case int:
		fmt.Println("a is type int")
	case string:
		fmt.Println("a is type string")
	default:
		fmt.Println("a not type found type")
	}
```

在 `swicth` 中，断言不再使用 `.(具体类型)`，而是 `a.(type)`。

最后，还有一种方式，使用的是反射包 `reflect` 来确定数据类型：

```go
    // 使用反射找出变量类型
	t := reflect.TypeOf(a)
	fmt.Printf("a is type: %s", t.Name())
```

这个包会直接使用非安全指针来获取真实的数据类型：

```go
func TypeOf(i interface{}) Type {
	eface := *(*emptyInterface)(unsafe.Pointer(&i))
	return toType(eface.typ)
}
```

一般日常开发，很少使用反射包。

## 二. 接口结构 interface 

我们现在都是函数式编程，或者是结构体方法式的编程，难道没有其他语言那种面向对象，对象继承的特征吗？有，`Golang` 语言叫做面向接口编程。

```go
package main

import (
	"fmt"
	"reflect"
)

// 定义一个接口，有一个方法
type A interface {
	Println()
}

// 定义一个接口，有两个方法
type B interface {
	Println()
	Printf() int
}

// 定义一个结构体
type A1Instance struct {
	Data string
}

// 结构体实现了Println()方法，现在它是一个 A 接口
func (a1 *A1Instance) Println() {
	fmt.Println("a1:", a1.Data)
}

// 定义一个结构体
type A2Instance struct {
	Data string
}

// 结构体实现了Println()方法，现在它是一个 A 接口
func (a2 *A2Instance) Println() {
	fmt.Println("a2:", a2.Data)
}

// 结构体实现了Printf()方法，现在它是一个 B 接口，它既是 A 又是 B 接口
func (a2 *A2Instance) Printf() int {
	fmt.Println("a2:", a2.Data)
	return 0
}

func main() {
	// 定义一个A接口类型的变量
	var a A

	// 将具体的结构体赋予该变量
	a = &A1Instance{Data: "i love you"}
	// 调用接口的方法
	a.Println()
	// 断言类型
	if v, ok := a.(*A1Instance); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a A1")
	}
	fmt.Println(reflect.TypeOf(a).String())

	// 将具体的结构体赋予该变量
	a = &A2Instance{Data: "i love you"}
	// 调用接口的方法
	a.Println()
	// 断言类型
	if v, ok := a.(*A1Instance); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a A1")
	}
	fmt.Println(reflect.TypeOf(a).String())

	// 定义一个B接口类型的变量
	var b B
	//b = &A1Instance{Data: "i love you"} // 不是 B 类型
	b = &A2Instance{Data: "i love you"}
	fmt.Println(b.Printf())
}

```

输出：

```go
a1: i love you
&{i love you}
*main.A1Instance
a2: i love you
not a A1
*main.A2Instance
a2: i love you
0
```

我们可以定义一个接口类型，使用 `type 接口名 interface`，这时候不再是 `interface{}`：

```go
// 定义一个接口，有一个方法
type A interface {
	Println()
}

// 定义一个接口，有两个方法
type B interface {
	Println()
	Printf() int
}
```

可以看到接口 `A` 和 `B` 是一种抽象的结构，每个接口都有一些方法在里面，只要结构体 `struct` 实现了这些方法，那么这些结构体都是这种接口的类型。如：

```go
// 定义一个结构体
type A1Instance struct {
	Data string
}

// 结构体实现了Println()方法，现在它是一个 A 接口
func (a1 *A1Instance) Println() {
	fmt.Println("a1:", a1.Data)
}

// 定义一个结构体
type A2Instance struct {
	Data string
}

// 结构体实现了Println()方法，现在它是一个 A 接口
func (a2 *A2Instance) Println() {
	fmt.Println("a2:", a2.Data)
}

// 结构体实现了Printf()方法，现在它是一个 B 接口，它既是 A 又是 B 接口
func (a2 *A2Instance) Printf() int {
	fmt.Println("a2:", a2.Data)
	return 0
}
```

我们要求结构体必须实现某些方法，所以可以定义一个接口类型的变量，然后将结构体赋值给它：

```go
	// 定义一个A接口类型的变量
	var a A
	// 将具体的结构体赋予该变量
	a = &A1Instance{Data: "i love you"}
	// 调用接口的方法
	a.Println()
```
a = &A1Instance{Data:"i love you"}而不是 a = A1Instance{Data:"i love you"}
指针类型的receiver 方法实现接口时，只有指针类型的对象实现了该接口。
对应上面的例子来说，只有&A1Instance实现了Println接口，而A1Instance根本没有实现该接口
当写成a = A1Instance{Data:"i love you"}
结构体没有实现该方法，将编译不通过，无法编译二进制。

当然也可以使用断言和反射来判断接口类型是属于哪个实际的结构体 `struct`。

```go
	// 断言类型
	if v, ok := a.(*A1Instance); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a A1")
	}
	fmt.Println(reflect.TypeOf(a).String())
```

`Golang` 很智能判断结构体是否实现了接口的方法，如果实现了，那么该结构体就是该接口类型。我们灵活的运用接口结构的特征，使用组合的形式就可以开发出更灵活的程序了。
