package main

import "fmt"

func main() {
	buffedChan := make(chan int, 2)
	buffedChan <- 2
	buffedChan <- 3
	for i := range buffedChan { // 必须关闭，否则死锁
		fmt.Println(i)
	}
}
