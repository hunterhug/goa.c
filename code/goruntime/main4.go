package main

import "fmt"

func main() {
	buffedChan := make(chan int, 2)
	buffedChan <- 2
	buffedChan <- 3
	close(buffedChan) // 关闭后才能for打印出，否则死锁

	//close(buffedChan) // 不能重复关闭
	//buffedChan <- 4  // 关闭后就不能再送数据了，但是之前的数据还在
	for i := range buffedChan { // 必须关闭，否则死锁
		fmt.Println(i)
	}
}
