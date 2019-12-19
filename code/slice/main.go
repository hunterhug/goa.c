package main

import "fmt"

func main() {
	// 创建一个容量为2的切片
	array := make([]int, 0, 2)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)

	// 虽然 append 但是没有赋予原来的变量 array
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)

	fmt.Println("-------")

	// 赋予回原来的变量
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1,1,1,1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1,1,1,1,1,1,1,1,1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
}
