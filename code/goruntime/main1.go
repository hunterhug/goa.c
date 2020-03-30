package main

import (
	"fmt"
	"time"
)

func Hu(ch chan int) {
	// 使用睡眠模仿一些耗时
	time.Sleep(2 * time.Second)
	fmt.Println("after 2 second hu!!!")

	// 执行语句后，通知主协程已经完成操作
	ch <- 1000
}

func main() {
	// 新建一个没有缓冲的信道
	ch := make(chan int)

	// 将信道传入函数，开启协程
	go Hu(ch)
	fmt.Println("start hu, wait...")

	// 从空缓冲的信道读取 int，将会堵塞，直到有消息到来
	v := <-ch
	fmt.Println("receive:", v)
}
