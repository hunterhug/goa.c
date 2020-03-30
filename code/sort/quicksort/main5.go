package main

import (
	"fmt"
	"sync"
)

// 链表栈，后进先出
type LinkStack struct {
	root *LinkNode  // 链表起点
	size int        // 栈的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value int
}

// 入栈
func (stack *LinkStack) Push(v int) {
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
func (stack *LinkStack) Pop() int {
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

// 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

// 非递归快速排序
func QuickSort5(array []int) {

	// 人工栈
	helpStack := new(LinkStack)

	// 第一次初始化栈，推入下标0，len(array)-1，表示第一次对全数组范围切分
	helpStack.Push(len(array) - 1)
	helpStack.Push(0)

	// 栈非空证明存在未排序的部分
	for !helpStack.IsEmpty() {
		// 出栈，对begin-end范围进行切分排序
		begin := helpStack.Pop() // 范围区间左边
		end := helpStack.Pop()   // 范围

		// 进行切分
		loc := partition(array, begin, end)

		// 右边范围入栈
		if loc+1 < end {
			helpStack.Push(end)
			helpStack.Push(loc + 1)
		}

		// 左边返回入栈
		if begin < loc-1 {
			helpStack.Push(loc - 1)
			helpStack.Push(begin)
		}
	}
}

// 非递归快速排序优化
func QuickSort6(array []int) {

	// 人工栈
	helpStack := new(LinkStack)

	// 第一次初始化栈，推入下标0，len(array)-1，表示第一次对全数组范围切分
	helpStack.Push(len(array) - 1)
	helpStack.Push(0)

	// 栈非空证明存在未排序的部分
	for !helpStack.IsEmpty() {
		// 出栈，对begin-end范围进行切分排序
		begin := helpStack.Pop() // 范围区间左边
		end := helpStack.Pop()   // 范围

		// 进行切分
		loc := partition(array, begin, end)

		// 切分后右边范围大小
		rSize := -1
		// 切分后左边范围大小
		lSize := -1

		// 右边范围入栈
		if loc+1 < end {
			rSize = end - (loc + 1)
		}

		// 左边返回入栈
		if begin < loc-1 {
			lSize = loc - 1 - begin
		}

		// 两个范围，让范围小的先入栈，减少人工栈空间
		if rSize != -1 && lSize != -1 {
			if lSize > rSize {
				helpStack.Push(end)
				helpStack.Push(loc + 1)
				helpStack.Push(loc - 1)
				helpStack.Push(begin)
			} else {
				helpStack.Push(loc - 1)
				helpStack.Push(begin)
				helpStack.Push(end)
				helpStack.Push(loc + 1)
			}
		} else {
			if rSize != -1 {
				helpStack.Push(end)
				helpStack.Push(loc + 1)
			}

			if lSize != -1 {
				helpStack.Push(loc - 1)
				helpStack.Push(begin)
			}
		}
	}
}

// 切分函数，并返回切分元素的下标
func partition(array []int, begin, end int) int {
	i := begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
	j := end       // array[end]是数组的最后一位

	// 没重合之前
	for i < j {
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i] // 交换
			j--
		} else {
			i++
		}
	}

	/* 跳出while循环后，i = j。
	 * 此时数组被分割成两个部分  -->  array[begin+1] ~ array[i-1] < array[begin]
	 *                        -->  array[i+1] ~ array[end] > array[begin]
	 * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
	 * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件就退出！
	 */
	if array[i] >= array[begin] { // 这里必须要取等“>=”，否则数组元素由相同的值组成时，会出现错误！
		i--
	}

	array[begin], array[i] = array[i], array[begin]
	return i
}

func main() {
	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort5(list3)
	fmt.Println(list3)

	list4 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort6(list4)
	fmt.Println(list4)
}
