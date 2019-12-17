# 链表，栈和队列

讲数据结构就离不开讲链表。因为数据结构是用来组织数据的，如何将一个数据关联到另外一个数据呢？链表可以将数据和数据之间关联起来，从一个数据指向另外一个数据。

## 一、链表

链表，是由一个个数据节点组成的，它是一个递归结构，要么它是空的，要么它存在一个指向另外一个数据节点的引用。

链表，可以说是最基础的数据结构。

最简单的链表如下：

```go
package main

import "fmt"

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	node := new(LinkNode)
	node.Data = 2
	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1
	
    // 按顺序打印数据
	nowNode := node
	for {
		if nowNode != nil {
			fmt.Println(nowNode.Data)
			nowNode = nowNode.NextNode
		} else {
			break
		}
	}
}
```

结构体 `LinkNode` 有两个字段，一个字段存放数据 `Data`，另一个字典指向下一个节点 `NextNode` 。这种从一个数据节点指向下一个数据节点的结构，都可以叫做链表。

有些书籍，把链表做了很细的划分，比如单链表，双链表，循环单链表，循环双链表，其实没有必要强行分类，链表就是从一个数据指向另外一个数据，一种将数据和数据关联起来的结构而已。

好吧，我们还是要知道是什么。

1. 单链表，就是链表是单向的，像我们上面这个结构一样，可以一直往下找到下一个数据节点，它只有一个方向，它不能往回找。
2. 双链表，每个节点既可以找到它之前的节点，也可以找到之后的节点，是双向的。
3. 循环链表，就是它一直往下找数据节点，最后回到了自己那个节点，形成了一个回路。循环单链表和循环双链表的区别就是，一个只能一个方向走，一个两个方向都可以走。

我们来实现一个循环链表，参见 `Golang` 标准库 `container/ring`：：

```go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ring implements operations on circular lists.
package ring

// A Ring is an element of a circular list, or ring.
// Rings do not have a beginning or end; a pointer to any ring element
// serves as reference to the entire ring. Empty rings are represented
// as nil Ring pointers. The zero value for a Ring is a one-element
// ring with a nil Value.
//
type Ring struct {
	next, prev *Ring
	Value      interface{} // for use by client; untouched by this library
}

func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// Next returns the next ring element. r must not be empty.
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// Prev returns the previous ring element. r must not be empty.
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
// in the ring and returns that ring element. r must not be empty.
//
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// New creates a ring of n elements.
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// Link connects ring r with ring s such that r.Next()
// becomes s and returns the original value for r.Next().
// r must not be empty.
//
// If r and s point to the same ring, linking
// them removes the elements between r and s from the ring.
// The removed elements form a subring and the result is a
// reference to that subring (if no elements were removed,
// the result is still the original value for r.Next(),
// and not nil).
//
// If r and s point to different rings, linking
// them creates a single ring with the elements of s inserted
// after r. The result points to the element following the
// last element of s after insertion.
//
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		// Note: Cannot use multiple assignment because
		// evaluation order of LHS is not specified.
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// Unlink removes n % r.Len() elements from the ring r, starting
// at r.Next(). If n % r.Len() == 0, r remains unchanged.
// The result is the removed subring. r must not be empty.
//
func (r *Ring) Unlink(n int) *Ring {
	if n <= 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// Len computes the number of elements in ring r.
// It executes in time proportional to the number of elements.
//
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

// Do calls function f on each element of the ring, in forward order.
// The behavior of Do is undefined if f changes *r.
func (r *Ring) Do(f func(interface{})) {
	if r != nil {
		f(r.Value)
		for p := r.Next(); p != r; p = p.next {
			f(p.Value)
		}
	}
}
```
## 二、数组和链表

数组是编程语言作为一种基本类型提供出来的，相同数据类型的元素按一定顺序排列的集合。

它的作用只有一种：存放数据，让你很快能找到存的数据。如果你不去额外改进它，它就只是存放数据而已，它不会将一个数据节点和另外一个数据节点关联起来。比如建立一个切片（可变长数组） `array`:

```go
package main

import "fmt"


//  打印出：
//  [0 0 0 0 0]
//  [8 9 7 0 0]
//  7
func main() {
	array := make([]int64, 5)
	fmt.Println(array)
	array[0] = 8
	array[1] = 9
	array[2] = 7
	fmt.Println(array)
	fmt.Println(array[2])
}
```

我们可以通过下标 `0，1，2` 来获取到数组中的数据，下标 `0，1，2` 就表示数据的位置，排第一位，排第二位，我们也可以把指定位置的数据替换成另外一个数据。它是一种将 `内存地址` 和 `数据` 映射起来的内置语法结构，数据和数据之间是没有关联的。

这一数据类型，是被编程语言高度抽象封装，底层其实是将数据和虚拟内存进行绑定， `下标` 会转换成 `虚拟内存地址`，然后去寻找`这个地址上的内存值`，这个过程是特别快的，时间复杂度为 `O(1)`。（当然这个虚拟内存不是真正的内存，每个程序启动都会有一个虚拟内存空间来映射真正的内存，这是计算机组成的内容，和数据结构也有点关系，我们会在另外的高级专题讲，这里就不展开了。）

用数组也可以实现链表，比如定义一个数组 `make([]Value, 5)`，值类型为一个结构体 `Value` ：

```go
package main

import "fmt"

func main() {
	type Value struct {
		Data      string
		NextIndex int64
	}

	array := make([]Value, 5)   // 五个节点的变长数组
	array[0] = Value{"I", 3}    // 下一个节点的下标为3
	array[1] = Value{"Army", 4} // 下一个节点的下标为4
	array[2] = Value{"You", 1}  // 下一个节点的下标为1
	array[3] = Value{"Love", 2} // 下一个节点的下标为2
	array[4] = Value{"!", -1}   // -1表示没有下一个节点
	node := array[0]
	for {
		fmt.Println(node.Data)
		if node.NextIndex == -1 {
			break
		}
		node = array[node.NextIndex]
	}
}
```

打印出：

```go
I
Love
You
Army
!
```

获取某个下标的数据，通过该数据可以知道下一个数据的下标是什么，然后拿出该下标的数据，继续往下做。问题是，有时候需要做删除，移动等各种操作，需要大量空间移动，所以某些情况下，数组的效率很低。

数组和链表是两个不同的概念。一个是编程语言提供的基本数据类型，表示一个连续的内存空间，可通过一个索引访问数据。另一个是我们定义的数据结构，表示从一个数据节点关联到另一个数据节点。

很多其他的数据结构都由数组和链表配合实现的。

## 三、栈 Stack 和队列 Queue

我们日常生活中，都需要将物品排列，或者安排事情的先后顺序。更通俗地讲，我们买东西时，人太多的情况下，我们要排队，排队也有先后顺序，有些人早了点来，排完队就离开了，有些人晚一点，才刚刚进去人群排队。

数据是有顺序的，从数据 `1` 到数据 `2`，再到数据 `3`，和日常生活一样，我们需要放数据，也需要排列数据，所以栈和队列出现了。

在计算机的世界里，会经常听见两种结构，`栈（stack）` 和 `队列 (queue)`。这其实是一种东西，就是 `有顺序的队列`，只不过出队的顺序不同。

1. 栈：先进后出，先进队的数据最后才出来。在英文的意思里，`stack` 可以作为一叠的意思，这个排列是垂直的，你将一张纸放在另外一张纸上面，先放的纸肯定是最后才会被拿走，因为上面有一张纸挡住了它。
2. 队列：先进先出，先进队的数据先出来。在英文的意思里，`queue` 和现实世界的排队意思一样，这个排列是水平的，先排先得。

我们可以用数据结构：`链表`（可连续或不连续的将数据与数据关联起来的结构），或 `数组`（连续的内存空间，按索引取值） 来实现 `队列`（实现 `栈`，只要把出队的方向改一下即可）：

数组形式的队列:

```go

```

链表形式的队列：

```go

```

## 四、集合 Set 和列表 List

我们又经常听到 `集合` 和 `列表` 这两种数据结构，这两种数据结构其实也是用于存放数据，不过它们是更宏观的概念，代表存放数据的容器。

我们一般写算法进行数据计算，数据处理，都需要有个地方来存数据，我们可以使用封装好的数据结构：`集合` 或 `列表`。

1. 列表 `List` ：存放数据，数据按顺序排列，可以依次入队和出队，有序号关系，可以取出某序号的数据。先进先出的 `队列 (queue)` 和先进后出的 `栈（stack）` 都是列表。
2. 集合 `Set` ：存放数据，特点就是没有数据会重复，会去重。你放一个数据进去，再放一个数据进去，如果两个数据一样，那么只会保存一份数据。集合可以没有顺序关系，也可以按值排序，算一种特殊的列表。

大家也经常听说一种叫 `线性表` 的数据结构，表示具有相同特性的数据元素的有限序列，实际上是上述的 `列表`，有 `顺序表示` 或 `链式表示`。

顺序表示：指的是用一组 `地址连续的存储单元` 依次存储线性表的数据元素，称为线性表的 `顺序存储结构`。它以 `物理位置相邻` 来表示线性表中数据元素间的逻辑关系，可随机存取表中任一元素。顺序表示的又叫 `顺序表`，也就是用数组来实现的列表。

链式表示：指的是用一组 `任意的存储单元` 存储线性表中的数据元素，称为线性表的 `链式存储结构`。它的存储单元可以是连续的，也可以是不连续的。在表示数据元素之间的逻辑关系时，除了存储其本身的信息之外，还需存储一个指示其直接后继的信息，也就是用链表来实现的列表。

我们可以实现一个链式的双向列表（也叫双端队列）：既可以先进先出 ` (queue)`，也可以先进后出 `（stack）` 的数据结构。它是由双向链表来实现的，参见 `Golang` 标准库 `container/list`：

```go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package list implements a doubly linked list.
//
// To iterate over a list (where l is a *List):
//	for e := l.Front(); e != nil; e = e.Next() {
//		// do something with e.Value
//	}
//
package list

// Element is an element of a linked list.
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value interface{}
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New returns an initialized list.
func New() *List { return new(List).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// remove removes e from its list, decrements l.len, and returns e.
func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}

// move moves e to next to at and returns e.
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e

	return e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}
```

在实际工程应用上，缓存数据库 `Redis` 的 `列表List` 基本类型就是用双端队列来实现的。

我们来实现集合 `Set` ：

```go

```



## 五、总结

`链表` 和 `数组` 可以用来辅助构建各种基本数据结构。

数据结构名字特别多，所以以后的计算机生涯中，有些自己造的数据结构，或者不常见的别人造的数据结构，不知道叫什么名字是很正常的。我们只需知道常见的数据结构即可，方便与其他程序员交流。