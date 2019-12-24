package main

import "fmt"

/*
	// growslice handles slice growth during append.

	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		if old.len < 1024 {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = cap
			}
		}
	}

*/
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
