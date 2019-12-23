package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := make(map[int]int, 10000)
	for i := 0; i < 10000; i++ {
		a[i] = i
	}

	runtime.GC()
	printMemStats("添加1万个键值对后")

	fmt.Println("删除前数组长度：", len(a))
	for i := 0; i < 10000-1; i++ {
		delete(a, i)
	}
	fmt.Println("删除后数组长度：", len(a))

	// 再次进行手动GC回收
	runtime.GC()
	printMemStats("删除1万个键值对后")

	// 设置为nil进行回收
	runtime.GC()
	a = nil
	printMemStats("设置为nil后")
}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v：分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}
