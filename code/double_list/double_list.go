package main

// 双端列表，双端队列
type DoubleList struct {
	head *ListNode // 指向链表头部
	tail *ListNode // 指向链表尾部
	len  int       // 列表长度
}

// 列表节点
type ListNode struct {
	pre   *ListNode // 前驱节点
	next  *ListNode // 后驱节点
	value string    // 值
}

// 获取节点值
func (node *ListNode) GetValue() string {
	return node.value
}

// 获取节点前驱节点
func (node *ListNode) GetPre() *ListNode {
	return node.pre
}

// 获取节点后驱节点
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// 返回列表长度
func (list *DoubleList) Len() int {
	return list.len
}

// 添加节点到链表头部
func (list *DoubleList) AddNodeHead(v string) {
	// 新节点
	newNode := new(ListNode)
	newNode.value = v

	// 如果头结点为空，那么链表为空，直接赋值
	if list.head == nil {
		list.head = newNode // 新节点是头部
		list.tail = newNode // 新节点也是尾部
	} else {
		// 老头部
		oldHead := list.head

		// 新节点成为新头部
		list.head = newNode

		// 头结点的后驱节点是之前的老头部，老头部的前驱节点是现在的新头部
		list.head.next = oldHead
		oldHead.pre = list.head
	}

	// 列表长度+1
	list.len = list.len + 1
}

// 添加节点到链表尾部
func (list *DoubleList) AddNodeTail(v string) {
	// 新节点
	newNode := new(ListNode)
	newNode.value = v

	// 如果尾结点为空，那么链表为空，直接赋值
	if list.tail == nil {
		list.head = newNode // 新节点是头部
		list.tail = newNode // 新节点也是尾部
	} else {
		// 新节点的前驱为链表的尾部，链表尾部的后驱是新的节点
		newNode.pre = list.tail
		list.tail.next = newNode

		// 现在链表的尾部变成新的节点了
		list.tail = newNode
	}

	// 列表长度+1
	list.len = list.len + 1
}

// 返回列表链表头结点
func (list *DoubleList) First() *ListNode {
	return list.head
}

// 返回列表链表尾结点
func (list *DoubleList) Last() *ListNode {
	return list.tail
}

// 从头部开始往后找，获取第N+1个位置的节点，索引从0开始。
func (list *DoubleList) IndexFromHead(n int) *ListNode {
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}

	head := list.head

	// 往后遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		head = head.next
	}

	return head
}

// 从尾部开始往前找，获取第N+1个位置的节点，索引从0开始。
func (list *DoubleList) IndexFromTail(n int) *ListNode {
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}

	tail := list.tail

	// 往前遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		tail = tail.pre
	}

	return tail
}

func main() {

}
