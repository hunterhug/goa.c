# 可变长数组

因为数组大小是固定的，当数据元素特别多时，固定的数组无法储存这么多的值，所以可变长数组出现了，这也是一种数据结构。在 `Golang `语言中，可变长数组被内置在语言里面：切片 `slice`。

`slice` 是对底层数组的抽象和控制。它是一个结构体：

````go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
````

1. 指向底层数组的指针。( `Golang` 语言是没有指针的数据类型，所以 `unsafe` 包提供相关的对内存指针的操作，一般情况下非专业人员勿用)
2. 切片的真正长度，也就是实际元素占用的大小。
3. 切片的容量，底层固定数组的长度。

每次可以初始化一个固定容量的切片，切片内部维护一个固定大小的数组。当 `append` 新元素时，固定大小的数组不够时会自动扩容，如：

```go
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
```

输出：

```go
cap 2 len 0 array: []
cap 2 len 0 array: []
cap 2 len 0 array: []
cap 2 len 0 array: []
-------
cap 2 len 1 array: [1]
cap 2 len 2 array: [1 1]
cap 4 len 3 array: [1 1 1]
cap 8 len 7 array: [1 1 1 1 1 1 1]
cap 16 len 16 array: [1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1]
```

我们可以看到 `Golang` 的切片无法原地 `append`，每次添加元素时返回新的引用地址，必须把该引用重新赋予之前的切片变量。并且，当容量不够时，会自动倍数递增扩容。

具体可查看标准库 `runtime` 下的 `slice.go` 文件。

## 一、实现可变长数组

我们来实现一个简单的可变长数组版本。

首先初始化一个切片时：