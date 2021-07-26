package main

import "fmt"

func main() {
	array := make([]int, 0)
	newCap := cap(array)
	for {
		if newCap > 1032000 {
			return
		}
		array = append(array, 1)

		if cap(array) > newCap {
			oldCap := newCap
			newCap = cap(array)
			if len(array) > 1024 {
				fmt.Printf("1.25 scale:%d+%d/4=%d,len:%d,new cap:%d,difference:%d-%d=%d\n", oldCap, oldCap, oldCap/4*5, len(array), newCap, newCap, oldCap/4*5, newCap-oldCap/4*5)
			} else {
				fmt.Printf("small 1024,len:%d,new cap:%d\n", len(array), newCap)
			}
		}
	}
}
