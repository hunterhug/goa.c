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
