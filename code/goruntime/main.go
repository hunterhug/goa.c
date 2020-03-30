package main

import (
	"fmt"
	"time"
)

func Hu() {
	// 使用睡眠模仿一些耗时
	time.Sleep(2 * time.Second)
	fmt.Println("after 2 second hu!!!")
}

func main() {

	// 将会堵塞
	//Hu()

	// 开启新的协程，不会堵塞
	go Hu()

	fmt.Println("start hu, wait...")

	// 必须死循环，不然主协程退出了，程序就结束了
	for {
		time.Sleep(1 * time.Second)
	}

}
