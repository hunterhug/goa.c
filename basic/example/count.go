package main

import "fmt"

func sum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total = total + i
	}
	return total
}

func sum2(n int) int {
	total := ((1 + n) * n) / 2
	return total
}

func main() {
	fmt.Println(sum(100))
	fmt.Println(sum2(100))
}
