/*
	左偏树
*/
package main

import (
	"fmt"
)

type LeftistHeap struct {
	Root *Node
}

type Node struct {
	Data       int
	Distance   int
	LeftChild  *Node
	RightChild *Node
}

func New() *LeftistHeap {
	h := new(LeftistHeap)
	return h
}

func (n *Node) Dist() int {
	if n == nil {
		return -1 // 空节点距离为-1
	}
	return n.Distance
}

func (h *LeftistHeap) Push(data int) {
	newNode := new(Node)
	newNode.Data = data

	h.Root = h.Root.Merge(newNode)
}

func (h *LeftistHeap) Pop() int {
	if h.Root == nil {
		return -1 // pop完
	}

	data := h.Root.Data
	h.Root = h.Root.LeftChild.Merge(h.Root.RightChild)
	return data
}

// 合并两棵左偏树
func (A *Node) Merge(B *Node) *Node {

	// 一棵树为空返回另外一棵树
	if A == nil {
		return B
	}

	if B == nil {
		return A
	}

	leftHeap := A
	rightHeap := B

	// 使左堆做为合并后的根节点
	if A.Data > B.Data {
		leftHeap = B
		rightHeap = A
	}

	// 递归：左堆的右子树和右堆进行合并，作为左堆右子树
	leftHeap.RightChild = leftHeap.RightChild.Merge(rightHeap)

	// 树翻转左右，确保左儿子距离大于右子
	if leftHeap.RightChild.Dist() > leftHeap.LeftChild.Dist() {
		leftHeap.LeftChild, leftHeap.RightChild = leftHeap.RightChild, leftHeap.LeftChild
	}

	if leftHeap.RightChild == nil {
		leftHeap.Distance = 0
	} else {
		leftHeap.Distance = leftHeap.RightChild.Dist() + 1
	}

	return leftHeap
}

// 递归先序排序
func (n *Node) Display() {
	if n == nil {
		fmt.Println("null")
		return
	}
	fmt.Println(n.Data)
	fmt.Printf("Node:%d,Left child：", n.Data)
	if n.LeftChild != nil {
		n.LeftChild.Display()
	} else {
		fmt.Print("null")
	}
	fmt.Println()
	fmt.Printf("Node:%d,Right child：", n.Data)
	if n.RightChild != nil {
		n.RightChild.Display()
	} else {
		fmt.Print("null")
	}
	fmt.Println()
}

func (h *LeftistHeap) Display() {
	h.Root.Display()
}

func main() {
	n := New()
	n.Display()

	fmt.Println("---")

	n.Push(3)
	n.Push(1)
	n.Push(5)
	n.Push(8)

	n.Display()

	fmt.Println(n.Pop())
	fmt.Println(n.Pop())
	fmt.Println(n.Pop())
	fmt.Println(n.Pop())
	fmt.Println(n.Pop())
	fmt.Println(n.Pop())

}
