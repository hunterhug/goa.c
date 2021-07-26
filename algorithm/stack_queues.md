# 栈和队列

## 一、栈 Stack 和队列 Queue

我们日常生活中，都需要将物品排列，或者安排事情的先后顺序。更通俗地讲，我们买东西时，人太多的情况下，我们要排队，排队也有先后顺序，有些人早了点来，排完队就离开了，有些人晚一点，才刚刚进去人群排队。

数据是有顺序的，从数据 `1` 到数据 `2`，再到数据 `3`，和日常生活一样，我们需要放数据，也需要排列数据。

在计算机的世界里，会经常听见两种结构，`栈（stack）` 和 `队列 (queue)`。它们是一种收集数据的有序集合（`Collection`），只不过删除和访问数据的顺序不同。

1. 栈：先进后出，先进队的数据最后才出来。在英文的意思里，`stack` 可以作为一叠的意思，这个排列是垂直的，你将一张纸放在另外一张纸上面，先放的纸肯定是最后才会被拿走，因为上面有一张纸挡住了它。
2. 队列：先进先出，先进队的数据先出来。在英文的意思里，`queue` 和现实世界的排队意思一样，这个排列是水平的，先排先得。

我们可以用数据结构：`链表`（可连续或不连续的将数据与数据关联起来的结构），或 `数组`（连续的内存空间，按索引取值） 来实现 `栈（stack）` 和 `队列 (queue)`。

数组实现：能快速随机访问存储的元素，通过下标 `index` 访问，支持随机访问，查询速度快，但存在元素在数组空间中大量移动的操作，增删效率低。

链表实现：只支持顺序访问，在某些遍历操作中查询速度慢，但增删元素快。

## 二、实现数组栈 ArrayStack

数组形式的下压栈，后进先出:

主要使用可变长数组来实现。

```go
// 数组栈，后进先出
type ArrayStack struct {
	array []string   // 底层切片
	size  int        // 栈的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}
```

我们来分析它的各操作。

### 2.1.入栈

```go
// 入栈
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	stack.array = append(stack.array, v)

	// 栈中元素数量+1
	stack.size = stack.size + 1
}
```

将元素入栈，会先加锁实现并发安全。

入栈时直接把元素放在数组的最后面，然后元素数量加 1。性能损耗主要花在切片追加元素上，切片如果容量不够会自动扩容，底层损耗的复杂度我们这里不计，所以时间复杂度为 `O(1)`。

### 2.2.出栈

```go
func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素
	v := stack.array[stack.size-1]

	// 切片收缩，但可能占用空间越来越大
	//stack.array = stack.array[0 : stack.size-1]

	// 创建新的数组，空间占用不会越来越大，但可能移动元素次数过多
	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray

	// 栈中元素数量-1
	stack.size = stack.size - 1
	return v
}
```

元素出栈，会先加锁实现并发安全。

如果栈大小为0，那么不允许出栈，否则从数组的最后面拿出元素。

元素取出后:

1. 如果切片偏移量向前移动 `stack.array[0 : stack.size-1]`，表明最后的元素已经不属于该数组了，数组变相的缩容了。此时，切片被缩容的部分并不会被回收，仍然占用着空间，所以空间复杂度较高，但操作的时间复杂度为：`O(1)`。
2. 如果我们创建新的数组 `newArray`，然后把老数组的元素复制到新数组，就不会占用多余的空间，但移动次数过多，时间复杂度为：`O(n)`。

最后元素数量减一，并返回值。

### 2.3.获取栈顶元素

```go
// 获取栈顶元素
func (stack *ArrayStack) Peek() string {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素值
	v := stack.array[stack.size-1]
	return v
}
```

获取栈顶元素，但不出栈。和出栈一样，时间复杂度为：`O(1)`。

### 2.4.获取栈大小和判定是否为空

```go
// 栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
```

一目了然，时间复杂度都是：`O(1)`。

### 2.5.示例

```go
func main() {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
}
```

输出：

```go
size: 3
pop: hen
pop: dog
size: 1
pop: drag
```

## 三、实现链表栈 LinkStack

链表形式的下压栈，后进先出：

```go
// 链表栈，后进先出
type LinkStack struct {
	root *LinkNode  // 链表起点
	size int        // 栈的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value string
}
```

我们来分析它的各操作。

### 3.1.入栈

```go
// 入栈
func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 如果栈顶为空，那么增加节点
	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
	} else {
		// 否则新元素插入链表的头部
		// 原来的链表
		preNode := stack.root

		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v

		// 原来的链表链接到新元素后面
		newNode.Next = preNode

		// 将新节点放在头部
		stack.root = newNode
	}

	// 栈中元素数量+1
	stack.size = stack.size + 1
}
```

将元素入栈，会先加锁实现并发安全。

如果栈里面的底层链表为空，表明没有元素，那么新建节点并设置为链表起点：`stack.root = new(LinkNode)`。

否则取出老的节点：`preNode := stack.root`，新建节点：`newNode := new(LinkNode)`，然后将原来的老节点链接在新节点后面： `newNode.Next = preNode`，最后将新节点设置为链表起点 `stack.root = newNode`。

时间复杂度为：`O(1)`。

### 3.2.出栈

```go
// 出栈
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 顶部元素要出栈
	topNode := stack.root
	v := topNode.Value

	// 将顶部元素的后继链接链上
	stack.root = topNode.Next

	// 栈中元素数量-1
	stack.size = stack.size - 1

	return v
}
```

元素出栈。如果栈大小为0，那么不允许出栈。

直接将链表的第一个节点 `topNode := stack.root` 的值取出，然后将表头设置为链表的下一个节点：`stack.root = topNode.Next`，相当于移除了链表的第一个节点。

时间复杂度为：`O(1)`。

### 3.3.获取栈顶元素

```go
// 获取栈顶元素
func (stack *LinkStack) Peek() string {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 顶部元素值
	v := stack.root.Value
	return v
}
```

获取栈顶元素，但不出栈。和出栈一样，时间复杂度为：`O(1)`。

### 3.4.获取栈大小和判定是否为空

```go
// 栈大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}
```

### 3.5.示例

```go
func main() {
	linkStack := new(LinkStack)
	linkStack.Push("cat")
	linkStack.Push("dog")
	linkStack.Push("hen")
	fmt.Println("size:", linkStack.Size())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Size())
	linkStack.Push("drag")
	fmt.Println("pop:", linkStack.Pop())
}
```

输出：

```go
size: 3
pop: hen
pop: dog
size: 1
pop: drag
```

## 四、实现数组队列 ArrayQueue

队列先进先出，和栈操作顺序相反，我们这里只实现入队，和出队操作，其他操作和栈一样。

```go
// 数组队列，先进先出
type ArrayQueue struct {
	array []string   // 底层切片
	size  int        // 队列的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}
```

### 4.1.入队

```go
// 入队
func (queue *ArrayQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	queue.array = append(queue.array, v)

	// 队中元素数量+1
	queue.size = queue.size + 1
}
```

直接将元素放在数组最后面即可，和栈一样，时间复杂度为：`O(n)`。

### 4.2.出队

```
// 出队
func (queue *ArrayQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 队中元素已空
	if queue.size == 0 {
		panic("empty")
	}

	// 队列最前面元素
	v := queue.array[0]

	/*	直接原位移动，但缩容后继的空间不会被释放
		for i := 1; i < queue.size; i++ {
			// 从第一位开始进行数据移动
			queue.array[i-1] = queue.array[i]
		}
		// 原数组缩容
		queue.array = queue.array[0 : queue.size-1]
	*/

	// 创建新的数组，移动次数过多
	newArray := make([]string, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		// 从老数组的第一位开始进行数据移动
		newArray[i-1] = queue.array[i]
	}
	queue.array = newArray

	// 队中元素数量-1
	queue.size = queue.size - 1
	return v
}
```

出队，把数组的第一个元素的值返回，并对数据进行空间挪位，挪位有两种：

1. 原地挪位，依次补位 `queue.array[i-1] = queue.array[i]`，然后数组缩容：`queue.array = queue.array[0 : queue.size-1]`，但是这样切片缩容的那部分内存空间不会释放。
2. 创建新的数组，将老数组中除第一个元素以外的元素移动到新数组。

时间复杂度是：`O(n)`。

## 五、实现链表队列 LinkQueue

队列先进先出，和栈操作顺序相反，我们这里只实现入队，和出队操作，其他操作和栈一样。

```go
// 链表队列，先进先出
type LinkQueue struct {
	root *LinkNode  // 链表起点
	size int        // 队列的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value string
}
```

### 5.1.入队

```go
// 入队
func (queue *LinkQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 如果栈顶为空，那么增加节点
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {
		// 否则新元素插入链表的末尾
		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v

		// 一直遍历到链表尾部
		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}

		// 新节点放在链表尾部
		nowNode.Next = newNode
	}

	// 队中元素数量+1
	queue.size = queue.size + 1
}
```

将元素放在链表的末尾，所以需要遍历链表，时间复杂度为：`O(n)`。

### 5.2.出队

```go
// 出队
func (queue *LinkQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 队中元素已空
	if queue.size == 0 {
		panic("empty")
	}

	// 顶部元素要出队
	topNode := queue.root
	v := topNode.Value

	// 将顶部元素的后继链接链上
	queue.root = topNode.Next

	// 队中元素数量-1
	queue.size = queue.size - 1

	return v
}
```

链表第一个节点出队即可，时间复杂度为：`O(1)`。

## 附录

代码下载： [https://github.com/hunterhug/goa.c/blob/master/code/stack/main.go](https://github.com/hunterhug/goa.c/blob/master/code/stack/main.go) 。