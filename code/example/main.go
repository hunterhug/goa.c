package main

import "fmt"

func changeTwo(a, b int) {
	a = 6
	b = 8
}

func main() {
	a, b := 1, 2
	fmt.Println(a, b)
	changeTwo(a, b)
	fmt.Println(a, b)
}
