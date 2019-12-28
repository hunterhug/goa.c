package main

import "fmt"

func main() {
	// 新建一个容量为4的字典 map
	m := make(map[string]int64, 4)

	// 放三个键值对
	m["dog"] = 1
	m["cat"] = 2
	m["hen"] = 3

	fmt.Println(m)

	// 查找 hen
	which := "hen"
	v, ok := m[which]
	if ok {
		// 找到了
		fmt.Println("find:", which, "value:", v)
	} else {
		// 找不到
		fmt.Println("not find:", which)
	}

	// 查找 ccc
	which = "ccc"
	v, ok = m[which]
	if ok {
		// 找到了
		fmt.Println("find:", which, "value:", v)
	} else {
		// 找不到
		fmt.Println("not find:", which)
	}

	m1 := make(map[int]int, 9)
	fmt.Printf("len:%v", len(m1))
	m1[5] = 9
	fmt.Printf("len:%v", len(m1))
}
