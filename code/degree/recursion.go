package main

import "fmt"

func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}

	return n * Rescuvie(n-1)
}

// 尾部递归是指递归函数在调用自身后直接传回其值，而不对其再加运算。
func RescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}

	return RescuvieTail(n-1, a*n)
}

func main() {
	fmt.Println(Rescuvie(5))
	fmt.Println()
	fmt.Println(RescuvieTail(5, 1))
}
