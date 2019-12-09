/*
	最小堆
*/
package main

import "fmt"

type Heap struct {
	Size  int
	Elems []int
}

func NewHeap(MaxSize int) *Heap {
	h := new(Heap)
	h.Elems = make([]int, MaxSize, MaxSize)
	return h
}

func (h *Heap) Push(x int) {
	h.Size++

	// i是要插入节点的下标
	i := h.Size
	for {
		if i <= 0 {
			break
		}

		// parent为父亲节点的下标
		parent := (i - 1) / 2
		// 如果父亲节点小于等于插入的值，则说明大小没有跌倒，可以退出
		if h.Elems[parent] <= x {
			break
		}

		// 互换当前父亲节点与要插入的值
		h.Elems[i] = h.Elems[parent]
		i = parent
	}

	h.Elems[i] = x
}

func (h *Heap) Pop() int {
	if h.Size == 0 {
		return 0
	}

	// 取出根节点
	ret := h.Elems[0]

	// 将最后一个节点的值提到根节点上
	h.Size--
	x := h.Elems[h.Size]

	i := 0
	for {
		// a，b为左右两个子节点的下标
		a := 2*i + 1
		b := 2*i + 2

		// 没有左子树
		if a >= h.Size {
			break
		}

		// 有右子树，找两个子节点中较小的值
		if b < h.Size && h.Elems[b] < h.Elems[a] {
			a = b
		}

		// 父亲小直接退出
		if h.Elems[a] >= x {
			break
		}

		// 交换
		h.Elems[i] = h.Elems[a]
		i = a
	}

	h.Elems[i] = x
	return ret
}

func (h *Heap) Display() {
	fmt.Printf("Size:%d,Elems:%#v\n", h.Size, h.Elems[0:h.Size])
}

func main() {
	h := NewHeap(100)
	h.Display()

	h.Push(3)
	h.Push(6)
	h.Push(7)
	h.Push(27)
	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Display()

	fmt.Println(h.Pop())
	h.Display()
	fmt.Println(h.Pop())
	h.Display()
	fmt.Println(h.Pop())
	h.Display()
	fmt.Println(h.Pop())
	h.Display()
	fmt.Println(h.Pop())
	h.Display()
}
