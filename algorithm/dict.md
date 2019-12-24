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

我们拿出键 `cat` 的值，就是 `2` 了。

`Golang` 提供了这一数据结构：`map`，并且要求键的数据类型必须是可比较的，因为如果不可比较，就无法知道键是存在还是不存在。

`Golang` 字典的一般的操作如下：

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

字典的实现有两种方式：哈希表 `HashTable` 和红黑树 `RBTree`。`Golang` 语言中字典 `map` 的实现由哈希表实现，具体可参考标准库 `runtime` 下的 `map.go` 文件。

我们会在《查找算法》章节：散列查找和红黑树中，具体分析字典的两种实现方式。

## 二、实现不可重复集合 Set

一般很多编程语言库，会把不可重复集合（`Collection`）命名为 `Set`，这个 `Set` 中文直译为集合，在某些上下文条件下，我们大脑要自动过滤，集合这词指的是不可重复集合还是指统称的集合，在这里都可以看到中文博大精深。

不可重复集合 `Set` 存放数据，特点就是没有数据会重复，会去重。你放一个数据进去，再放一个数据进去，如果两个数据一样，那么只会保存一份数据。

集合 `Set` 可以没有顺序关系，也可以按值排序，算一种特殊的列表。

因为我们知道字典的键是不重复的，所以只要我们不考虑字典的值，就可以实现集合，我们来实现存整数的集合 `Set`：

```go
// 集合结构体
type Set struct {
	m            map[int]struct{} // 用字典来实现，因为字段键不能重复
	len          int          // 集合的大小
	sync.RWMutex              // 锁，实现并发安全
}
```

### 2.1.初始化一个集合

```go
// 新建一个空集合
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp,
	}
}
```

使用字典来初始化一个容量为 `cap` 的不可重复集合。`map` 的值我们不使用，所以使用空结构体，因为空结构体不占用内存空间。如：

```go
package main

import (
	"fmt"
	"sync"
)

func main()
    // 为什么使用空结构体
	a := struct{}{}
	b := struct{}{}
	if a == b {
		fmt.Printf("right:%p\n", &a)
	}

	fmt.Println(unsafe.Sizeof(a))
}
```

会打印出：

```go
right:0x1198a98
0
```

空结构体的内存地址都一样，并且不占用内存空间。

### 2.2.添加一个元素

```go
// 增加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{} // 实际往字典添加这个键
	s.len = s.len + 1      // 集合大小增加
}
```

首先，加并发锁，实现线程安全，然后往结构体 `s *Set` 里面的内置 `map` 添加该元素，元素作为字典的键，会自动去重。同时，集合大小加1。

时间复杂度等于字典存键值对的复杂度，哈希不冲突的时间复杂度为：`O(1)`，否则为 `O(n)`，可看哈希表实现一章。

### 2.3.删除一个元素

```go
// 移除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	s.Unlock()
	delete(s.m, item) // 实际从字典删除这个键
	s.len = s.len - 1 // 集合大小减少
}
```

同理，删除 `map` 里面的键。

### 2.3.查看元素是否在集合中

```go
// 查看是否存在元素
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}
```

### 2.4.查看集合大小

```go
// 查看集合大小
func (s *Set) Len() int {
	return s.len
}
```

### 2.5.查看集合是否为空

```go
// 集合是够为空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}
```

### 2.6.清除集合所有元素

```go
// 清除集合所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{} // 字典重新赋值
	s.len = 0                // 大小归零
}
```

将原先的 `map` 释放掉，并且重新赋一个空的 `map`。

### 2.7.将集合转化为列表

```go
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
```

### 2.8.完整例子

```go
package main

import (
	"fmt"
	"sync"
	"unsafe"
)

// 集合结构体
type Set struct {
	m            map[int]struct{} // 用字典来实现，因为字段键不能重复
	len          int              // 集合的大小
	sync.RWMutex                  // 锁，实现并发安全
}

// 新建一个空集合
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp,
	}
}

// 增加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{} // 实际往字典添加这个键
	s.len = s.len + 1      // 集合大小增加
}

// 移除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	s.Unlock()
	delete(s.m, item) // 实际从字典删除这个键
	s.len = s.len - 1 // 集合大小减少
}

// 查看是否存在元素
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// 查看集合大小
func (s *Set) Len() int {
	return s.len
}

// 清除集合所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{} // 字典重新赋值
	s.len = 0                // 大小归零
}

// 集合是够为空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// 将集合转化为列表
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func main() {
	// 初始化一个容量为5的不可重复集合
	s := NewSet(5)

	s.Add(1)
	s.Add(1)
	s.Add(2)
	fmt.Println("list of all items", s.List())

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("empty")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		fmt.Println("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items", s.List())
}
```

打印出:

```go
list of all items [1 2]
empty
2 does exist
list of all items [1]
```
