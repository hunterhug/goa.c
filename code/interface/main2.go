package main

import (
	"fmt"
	"reflect"
)

// A 定义一个接口，有一个方法
type A interface {
	Println()
}

// B 定义一个接口，有两个方法
type B interface {
	Println()
	Printf() int
}

// A1Instance 定义一个结构体
type A1Instance struct {
	Data string
}

// Println 结构体实现了Println()方法，现在它是一个 A 接口
func (a1 *A1Instance) Println() {
	fmt.Println("a1:", a1.Data)
}

// A2Instance 定义一个结构体
type A2Instance struct {
	Data string
}

// Println 结构体实现了Println()方法，现在它是一个 A 接口
func (a2 *A2Instance) Println() {
	fmt.Println("a2:", a2.Data)
}

// Printf 结构体实现了Printf()方法，现在它是一个 B 接口，它既是 A 又是 B 接口
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
