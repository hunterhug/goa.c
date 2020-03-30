package main

import "fmt"

func main() {
	buffedChan1 := make(chan int, 2)
	buffedChan1 <- 2
	buffedChan1 <- 3
	j, ok := <-buffedChan1
	fmt.Println(j, ok)
	j, ok = <-buffedChan1
	fmt.Println(j, ok)

	close(buffedChan1)    // 关闭后才能，否则死锁
	j, ok = <-buffedChan1 // 如果未关闭，否则堵塞后死锁
	fmt.Println(j, ok)

	select {
	case j, ok := <-buffedChan1:
		fmt.Println("jjj", j, ok)
	default:
		fmt.Println("will not out")
	}
}
