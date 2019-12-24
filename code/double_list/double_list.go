package main

// 双端列表，双端队列
type DoubleList struct {
	head *listNode // 指向链表头部
	tail *listNode // 指向链表尾部
	len  *listNode // 列表长度
}

// 列表节点
type listNode struct {
	pre   *listNode // 前驱节点
	next  *listNode // 后驱节点
	value string    // 值
}

func main() {

}
