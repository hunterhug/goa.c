# 字典

字典是存储键值对的数据结构，把一个键和一个值映射起来，一一映射，键不能重复。如：

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