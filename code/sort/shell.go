package main

import "fmt"

// ShellSort 增量序列折半的希尔排序
func ShellSort(list []int) {
	// 数组长度
	n := len(list)

	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		// 直接插入排序算法请看 《插入排序》章节

		// 插入排序开始
		for i := step; i <= n-1; i += step {
			deal := list[i] // 待排序的数
			j := i - step   // 待排序的数左边的最近一个数的位置

			// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
			if deal < list[j] {
				// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
				for ; j >= 0 && deal < list[j]; j -= step {
					list[j+step] = list[j] // 某数后移，给待排序留空位(注意移动是以step为步长)
				}
				list[j+step] = deal // 结束了，待排序的数插入空位
			}
		}
		// 插入排序结束

	}
}

func main() {
	list := []int{5}
	ShellSort(list)
	fmt.Println(list)

	list1 := []int{5, 9}
	ShellSort(list1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	ShellSort(list2)
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 2, 4, 23, 467, 85, 23, 567, 335, 677, 33, 56, 2, 5, 33, 6, 8, 3}
	ShellSort(list3)
	fmt.Println(list3)
}
