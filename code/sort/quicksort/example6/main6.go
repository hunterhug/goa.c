package main

import (
	"fmt"
	"sort"
)

func InnerSort() {
	list := []struct {
		Name string
		Age  int
	}{
		{"A", 75},
		{"B", 4},
		{"C", 5},
		{"D", 5},
		{"E", 2},
		{"F", 5},
		{"G", 5},
	}

	sort.SliceStable(list, func(i, j int) bool { return list[i].Age < list[j].Age })
	fmt.Println(list)

	list2 := []struct {
		Name string
		Age  int
	}{
		{"A", 75},
		{"B", 4},
		{"C", 5},
		{"D", 5},
		{"E", 2},
		{"F", 5},
		{"G", 5},
	}

	sort.Slice(list2, func(i, j int) bool { return list2[i].Age < list2[j].Age })
	fmt.Println(list2)
}

func main() {
	InnerSort()
}
