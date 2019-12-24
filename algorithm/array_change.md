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

1. 指向底层数组的指针。( `Golang` 语言是没有操作原始内存的指针的，所以 `unsafe` 包提供相关的对内存指针的操作，一般情况下非专业人员勿用)
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
	array = append(array, 1, 1, 1, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1, 1, 1, 1, 1, 1, 1, 1, 1)
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

具体可参考标准库 `runtime` 下的 `slice.go` 文件。

## 一、实现可变长数组

我们来实现一个简单的，存放整数的，可变长的数组版本。

因为 `Golang` 的限制，不允许使用 `[n]int` 来创建一个固定大小为 `n` 的整数数组，只允许使用常量来创建大小。

所以我们这里会使用切片的部分功能来代替数组，虽然切片本身是可变长数组，但是我们不会用到它的 `append` 功能，只把它当数组用。

```go
import (
	"sync"
)

// 可变长数组
type Array struct {
	array []int      // 固定大小的数组，用满容量和满大小的切片来代替
	len   int        // 真正长度
	cap   int        // 容量
	lock  sync.Mutex // 为了并发安全使用的锁
}
```

### 1.1. 初始化数组

创建一个 `len` 个元素，容量为 `cap` 的可变长数组：

```
// 新建一个可变长数组
func Make(len, cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len large than cap")
	}

	// 把切片当数组用
	array := make([]int, cap, cap)

	// 元数据
	s.array = array
	s.cap = cap
	s.len = 0
	return s
}
```

主要利用满容量和满大小的切片来充当固定数组，结构体 `Array` 里面的字段 `len` 和 `cap` 来控制值的存取。不允许设置 `len > cap` 的可变长数组。

时间复杂度为：`O(1)`，因为分配内存空间和设置几个值是常数时间。

### 1.2. 添加元素

```go
// 增加一个元素
func (a *Array) Append(element int) {
	// 并发锁
	a.lock.Lock()
	defer a.lock.Unlock()

	// 大小等于容量，表示没多余位置了
	if a.len == a.cap {
		// 没容量，数组要扩容，扩容到两倍
		newCap := 2 * a.len

		// 如果之前的容量为0，那么新容量为1
		if a.cap == 0 {
			newCap = 1
		}

		newArray := make([]int, newCap, newCap)

		// 把老数组的数据移动到新数组
		for k, v := range a.array {
			newArray[k] = v
		}

		// 替换数组
		a.array = newArray
		a.cap = newCap

	}

	// 把元素放在数组里
	a.array[a.len] = element
	// 真实长度+1
	a.len = a.len + 1

}
```

首先添加一个元素到可变长数组里，会加锁，这样会保证并发安全。然后将值放在数组里：`a.array[a.len] = element`，然后 `len + 1`，表示真实大小又多了一个。

当真实大小 `len = cap` 时，表明位置都用完了，没有多余的空间放新值，那么会创建一个固定大小 `2*len` 的新数组来替换老数组：`a.array = newArray`，当然容量也会变大：`a.cap = newCap`。如果一开始设置的容量 `cap = 0`，那么新的容量会是从 1 开始。

添加元素中，耗时主要在老数组中的数据移动到新数组，时间复杂度为：`O(n)`。当然，如果容量够的情况下，时间复杂度会变为：`O(1)`。

如何添加多个元素：

```go
// 增加多个元素
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}
```

只是简单遍历一下，调用 `Append` 函数。其中 `...int` 是 `Golang` 的语言特征，表示多个函数变量。

### 1.3. 获取指定下标元素

```go
// 获取某个下标的元素
func (a *Array) Get(index int) int {
	// 越界了
	if a.len == 0 || index >= a.len {
		panic("index over len")
	}
	return a.array[index]
}
```

当可变长数组的真实大小为0，或者下标 `index` 超出了真实长度 `len` ，将会 `panic` 越界。

因为只获取下标的值，所以时间复杂度为 `O(1)`。

### 1.4. 获取真实长度和容量

```go
// 返回真实长度
func (a *Array) Len() int {
	return a.len
}

// 返回容量
func (a *Array) Cap() int {
	return a.cap
}
```

时间复杂度为 `O(1)`。

### 1.5. 示例

现在我们来运行完整的可变长数组的例子：

```go
package main

import (
	"fmt"
	"sync"
)

// 可变长数组
type Array struct {
	array []int      // 固定大小的数组，用满容量和满大小的切片来代替
	len   int        // 真正长度
	cap   int        // 容量
	lock  sync.Mutex // 为了并发安全使用的锁
}

// 新建一个可变长数组
func Make(len, cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len large than cap")
	}

	// 把切片当数组用
	array := make([]int, cap, cap)

	// 元数据
	s.array = array
	s.cap = cap
	s.len = 0
	return s
}

// 增加一个元素
func (a *Array) Append(element int) {
	// 并发锁
	a.lock.Lock()
	defer a.lock.Unlock()

	// 大小等于容量，表示没多余位置了
	if a.len == a.cap {
		// 没容量，数组要扩容，扩容到两倍
		newCap := 2 * a.len

		// 如果之前的容量为0，那么新容量为1
		if a.cap == 0 {
			newCap = 1
		}

		newArray := make([]int, newCap, newCap)

		// 把老数组的数据移动到新数组
		for k, v := range a.array {
			newArray[k] = v
		}

		// 替换数组
		a.array = newArray
		a.cap = newCap

	}

	// 把元素放在数组里
	a.array[a.len] = element
	// 真实长度+1
	a.len = a.len + 1

}

// 增加多个元素
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}

}

// 获取某个下标的元素
func (a *Array) Get(index int) int {
	// 越界了
	if a.len == 0 || index >= a.len {
		panic("index over len")
	}
	return a.array[index]
}

// 返回真实长度
func (a *Array) Len() int {
	return a.len
}

// 返回容量
func (a *Array) Cap() int {
	return a.cap
}

// 辅助打印
func Print(array *Array) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		// 第一个元素
		if i == 0 {
			result = fmt.Sprintf("%s%d", result, array.Get(i))
			continue
		}

		result = fmt.Sprintf("%s %d", result, array.Get(i))
	}
	result = result + "]"
	return
}

func main() {
	// 创建一个容量为3的动态数组
	a := Make(0, 3)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))

	// 增加一个元素
	a.Append(10)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))

	// 增加一个元素
	a.Append(9)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))

	// 增加多个元素
	a.AppendMany(8, 7)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
}
```

将打印出：

```go
cap 3 len 0 array: []
cap 3 len 1 array: [10]
cap 3 len 2 array: [10 9]
cap 6 len 4 array: [10 9 8 7]
```

可以看到，容量会自动翻倍。

## 二、总结

可变长数组在实际开发上，经常会使用到，其在固定大小数组的基础上，会自动进行容量扩展。

因为这一数据结构的使用频率太高了，所以，`Golang` 自动提供了这一数据类型：切片（可变长数组）。大家一般开发过程中，直接使用这一类型即可。