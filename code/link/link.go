package main

import (
	"fmt"
)

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
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2

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
}
