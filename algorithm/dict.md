# 字典

字典是存储键值对的数据结构，把一个键和一个值映射起来，一一映射，键不能重复。如：

```
键=>值

"a"=>1
"b"=>2
"c"=>1
"d"=>2
```

我们拿出键 `a` 的值，就是 `1` 了。`Golang` 提供了这一数据结构：`map`：

```go

f := map[string]int64{"a": 3, "b": 4} // map

// 查找map键值
v, ok := f["f"]
fmt.Println(v, ok)
v, ok = f["ff"]
fmt.Println(v, ok)
```

字典的实现有两种方式：哈希表 `HashTable` 和红黑树 `RBTree`。

在 `Golang` 语言中字典 `map` 的实现由哈希表实现，具体可查看标准库 `runtime` 下的 `map.go` 文件。

我们会在《查找算法》章节：散列查找和红黑树具体分析字典的两种实现方式。