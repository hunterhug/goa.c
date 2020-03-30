package main

import (
	"fmt"
	"sync"
)

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value string
}

// 数组队列，先进先出
type ArrayQueue struct {
	array []string   // 底层切片
	size  int        // 队列的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

// 入队
func (queue *ArrayQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	queue.array = append(queue.array, v)

	// 队中元素数量+1
	queue.size = queue.size + 1
}

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

// 链表队列，先进先出
type LinkQueue struct {
	root *LinkNode  // 链表起点
	size int        // 队列的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

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

func main() {
	q := new(ArrayQueue)
	q.Add("1")
	q.Add("2")
	q.Add("3")
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())

	q2 := new(LinkQueue)
	q2.Add("1")
	q2.Add("2")
	q2.Add("3")
	fmt.Println(q2.Remove())
	fmt.Println(q2.Remove())
	fmt.Println(q2.Remove())
}
