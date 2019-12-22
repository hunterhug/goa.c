# 字典

## 一、字典

字典是存储键值对的数据结构，把一个键和一个值映射起来，一一映射，键不能重复。在某些教程中，将这种结构称为符号表。

如：

```
键=>值

"cat"=>2
"dog"=>1
"hen"=>3
```

我们拿出键 `cat` 的值，就是 `2` 了。`Golang` 提供了这一数据结构：`map`：

```go

package main

import "fmt"

func main() {
	// 新建一个容量为4的字典 map
	m := make(map[string]int64, 4)

	// 放三个键值对
	m["dog"] = 1
	m["cat"] = 2
	m["hen"] = 3

	fmt.Println(m)

	// 查找 hen
	which := "hen"
	v, ok := m[which]
	if ok {
		// 找到了
		fmt.Println("find:", which, "value:", v)
	} else {
		// 找不到
		fmt.Println("not find:", which)
	}

	// 查找 ccc
	which = "ccc"
	v, ok = m[which]
	if ok {
		// 找到了
		fmt.Println("find:", which, "value:", v)
	} else {
		// 找不到
		fmt.Println("not find:", which)
	}
}

```

字典的实现有两种方式：哈希表 `HashTable` 和红黑树 `RBTree`。

在 `Golang` 语言中字典 `map` 的实现由哈希表实现，具体可参考标准库 `runtime` 下的 `map.go` 文件。

我们会在《查找算法》章节：散列查找和红黑树具体分析字典的两种实现方式。

## 二、实现不可重复集合 Set

一般很多编程语言库，会把不可重复集合（`Collection`）命名为 `Set`，这个 `Set` 中文直译为集合，在某些上下文条件下，我们大脑要自动过滤，集合这词指的是不可重复集合还是指统称的集合，在这里都可以看到中文博大精深。

不可重复集合 `Set` 存放数据，特点就是没有数据会重复，会去重。你放一个数据进去，再放一个数据进去，如果两个数据一样，那么只会保存一份数据。

集合 `Set` 可以没有顺序关系，也可以按值排序，算一种特殊的列表。

因为我们知道字典的键是不重复的，所以只要我们不考虑字典的值，就可以实现集合，我们来实现存整数的集合 `Set`：

```go
// 集合结构体
type Set struct {
	m            map[int]bool // 用字典来实现，因为字段键不能重复
	len          int          // 集合的大小
	sync.RWMutex              // 锁，实现并发安全
}
```

### 1.1