package main

// 数组栈，先进先出
type ArrayStack struct {
	array []int64 // 底层切片
	len   int64   // 切片的真实长度，辅助字段
	size  int64   // 队列的元素数量
}

// 链表栈，先进先出
type LinkStack struct {
	root *LinkNode // 链表起点
	size int64     // 队列的元素数量
}

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value int64
}

func (stack *ArrayStack) Insert(v int64) {

}
