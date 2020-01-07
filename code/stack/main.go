package main

import (
	"fmt"
	"sync"
)

// 数组栈，后进先出
type ArrayStack struct {
	array []string   // 底层切片
	size  int        // 栈的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

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

// 栈大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

// 入栈
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	stack.array = append(stack.array, v)

	// 栈中元素数量+1
	stack.size = stack.size + 1
}

// 出栈
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

// 栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

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
