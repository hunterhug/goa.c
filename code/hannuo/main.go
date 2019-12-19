package main

import "fmt"

var total = 0

// 汉诺塔
// 一开始A杆上有N个盘子，B和C杆都没有盘子。
func main() {
	n := 4   // 64 个盘子
	a := "a" // 杆子A
	b := "b" // 杆子B
	c := "c" // 杆子C
	tower(n, a, b, c)

	// 当 n=1 时，移动次数为 1
	// 当 n=2 时，移动次数为 3
	// 当 n=3 时，移动次数为 7
	// 当 n=4 时，移动次数为 15
	fmt.Println(total)
}

// 表示将N个盘子，从 a 杆，借助 b 杆移到 c 杆
func tower(n int, a, b, c string) {
	if n == 1 {
		total = total + 1
		fmt.Println(a, "->", c)
		return
	}

	tower(n-1, a, c, b)
	total = total + 1
	fmt.Println(a, "->", c)
	tower(n-1, b, a, c)
}
