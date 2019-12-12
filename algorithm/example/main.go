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

	array := make([]int64, 5)
	fmt.Println(array)
	array[0] = 8
	array[1] = 9
	array [2] = 7
	fmt.Println(array)
	fmt.Println(array[2])
}
