package main

import "fmt"

// InsertSort 改进：当数组规模小时使用直接插入排序
func InsertSort(list []int) {
	n := len(list)
	// 进行 N-1 轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		if deal < list[j] {
			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j] // 某数后移，给待排序留空位
			}
			list[j+1] = deal // 结束了，待排序的数插入空位
		}
	}
}

func QuickSort1(array []int, begin, end int) {
	if begin < end {
		// 当数组小于 4 时使用直接插入排序
		if end-begin <= 4 {
			InsertSort(array[begin : end+1])
			return
		}

		// 进行切分
		loc := partition(array, begin, end)
		// 对左部分进行快排
		QuickSort1(array, begin, loc-1)
		// 对右部分进行快排
		QuickSort1(array, loc+1, end)
	}
}

// 切分函数，并返回切分元素的下标
func partition(array []int, begin, end int) int {
	i := begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
	j := end       // array[end]是数组的最后一位

	// 没重合之前
	for i < j {
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i] // 交换
			j--
		} else {
			i++
		}
	}

	/* 跳出 for 循环后，i = j。
	 * 此时数组被分割成两个部分   -->  array[begin+1] ~ array[i-1] < array[begin]
	 *                        -->  array[i+1] ~ array[end] > array[begin]
	 * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
	 * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件就退出！
	 */
	if array[i] >= array[begin] { // 这里必须要取等“>=”，否则数组元素由相同的值组成时，会出现错误！
		i--
	}

	array[begin], array[i] = array[i], array[begin]
	return i
}

func main() {
	list := []int{5}
	QuickSort1(list, 0, len(list)-1)
	fmt.Println(list)

	list1 := []int{5, 9}
	QuickSort1(list1, 0, len(list1)-1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1}
	QuickSort1(list2, 0, len(list2)-1)
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort1(list3, 0, len(list3)-1)
	fmt.Println(list3)
}
