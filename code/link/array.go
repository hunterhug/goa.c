package main

import "fmt"


//  打印出：
//  [0 0 0 0 0]
//  [8 9 7 0 0]
//  7
func main() {
	array := [5]int64{}
	fmt.Println(array)
	array[0] = 8
	array[1] = 9
	array[2] = 7
	fmt.Println(array)
	fmt.Println(array[2])
}