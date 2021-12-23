# 分治法和递归

在计算机科学中，分治法是一种很重要的算法。

字面上的解释是`分而治之`，就是把一个复杂的问题分成两个或更多的相同或相似的子问题。

直到最后子问题可以简单的直接求解，原问题的解即子问题的解的合并。

分治法一般使用递归来求问题的解。

## 一、递归

递归就是不断地调用函数本身。

比如我们求阶乘 `1 * 2 * 3 * 4 * 5 *...* N`：

```go
package main
import "fmt"

func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}

	return n * Rescuvie(n-1)
}

func main() {
    fmt.Println(Rescuvie(5))
}
```

会反复进入一个函数，它的过程如下:

```go
Rescuvie(5)
{5 * Rescuvie(4)}
{5 * {4 * Rescuvie(3)}}
{5 * {4 * {3 * Rescuvie(2)}}}
{5 * {4 * {3 * {2 * Rescuvie(1)}}}}
{5 * {4 * {3 * {2 * 1}}}}
{5 * {4 * {3 * 2}}}
{5 * {4 * 6}}
{5 * 24}
120
```

函数不断地调用本身，并且还乘以一个变量：`n * Rescuvie(n-1)`，这是一个递归的过程。

很容易看出, 因为递归式使用了运算符，每次重复的调用都使得运算的链条不断加长，系统不得不使用栈进行数据保存和恢复。

如果每次递归都要对越来越长的链进行运算，那速度极慢，并且可能栈溢出，导致程序奔溃。

所以有另外一种写法，叫尾递归：

```go
package main
import "fmt"

func RescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}

	return RescuvieTail(n-1, a*n)
}


func main() {
    fmt.Println(RescuvieTail(5, 1))
}
```

他的递归过程如下:

```go
RescuvieTail(5, 1)
RescuvieTail(4, 1*5)=RescuvieTail(4, 5)
RescuvieTail(3, 5*4)=RescuvieTail(3, 20)
RescuvieTail(2, 20*3)=RescuvieTail(2, 60)
RescuvieTail(1, 60*2)=RescuvieTail(1, 120)
120
```

尾部递归是指递归函数在调用自身后直接传回其值，而不对其再加运算，效率将会极大的提高。

如果一个函数中所有递归形式的调用都出现在函数的末尾，我们称这个递归函数是尾递归的。当递归调用是整个函数体中最后执行的语句且它的返回值不属于表达式的一部分时，这个递归调用就是尾递归。尾递归函数的特点是在回归过程中不用做任何操作，这个特性很重要，因为大多数现代的编译器会利用这种特点自动生成优化的代码。-- 来自百度百科。

尾递归函数，部分高级语言编译器会进行优化，减少不必要的堆栈生成，使得程序栈维持固定的层数，不会出现栈溢出的情况。

我们将会举多个例子说明。

## 二、例子：斐波那契数列

斐波那契数列是指，后一个数是前两个数的和的一种数列。如下：

```go
1 1 2 3 5 8 13 21 ... N-1 N 2N-1
```

尾递归的求解为：

````go
package main
import "fmt"


func F(n int, a1, a2 int) int {
	if n == 0 {
		return a1
	}

	return F(n-1, a2, a1+a2)

}

func main() {
	fmt.Println(F(1, 1, 1))
	fmt.Println(F(2, 1, 1))
	fmt.Println(F(3, 1, 1))
	fmt.Println(F(4, 1, 1))
	fmt.Println(F(5, 1, 1))
}
````

输出：

```go
1
2
3
5
8
```

当 `n=5` 的递归过程如下:

```go
F(5,1,1)
F(4,1,1+1)=F(4,1,2)
F(3,2,1+2)=F(3,2,3)
F(2,3,2+3)=F(2,3,5)
F(1,5,3+5)=F(1,5,8)
F(0,8,5+8)=F(0,8,13)
8
```

## 三、例子：二分查找

在一个已经排好序的数列，找出某个数，如：

```go
1 5 9 15 81 89 123 189 333
```

从上面排好序的数列中找出数字 `189`。

二分查找的思路是，先拿排好序数列的中位数与目标数字 `189` 对比，如果刚好匹配目标，结束。

如果中位数比目标数字大，因为已经排好序，所以中位数右边的数字绝对都比目标数字大，那么从中位数的左边找。

如果中位数比目标数字小，因为已经排好序，所以中位数左边的数字绝对都比目标数字小，那么从中位数的右边找。

这种分而治之，一分为二的查找叫做二分查找算法。

递归解法：

```go
package main

import "fmt"

// 二分查找递归解法
func BinarySearch(array []int, target int, l, r int) int {
	if l > r {
		// 出界了，找不到
		return -1
	}

	// 从中间开始找
	mid := (l + r) / 2
	middleNum := array[mid]

	if middleNum == target {
		return mid // 找到了
	} else if middleNum > target {
		// 中间的数比目标还大，从左边找
		return BinarySearch(array, target, 0, mid-1)
	} else {
		// 中间的数比目标还小，从右边找
		return BinarySearch(array, target, mid+1, r)
	}

}

func main() {
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 500
	result := BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)

	target = 189
	result = BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)
}
```

输出：

```go
500 -1
189 7
```

可以看到，`189` 这个数字在数列的下标 `7` 处，而 `500` 这个数找不到。

当然，递归解法都可以转化为非递归，如：

```go
package main

import "fmt"

// 二分查找非递归解法
func BinarySearch2(array []int, target int, l, r int) int {
	ltemp := l
	rtemp := r

	for {
		if ltemp > rtemp {
			// 出界了，找不到
			return -1
		}

		// 从中间开始找
		mid := (ltemp + rtemp) / 2
		middleNum := array[mid]

		if middleNum == target {
			return mid // 找到了
		} else if middleNum > target {
			// 中间的数比目标还大，从左边找
			rtemp = mid - 1
		} else {
			// 中间的数比目标还小，从右边找
			ltemp = mid + 1
		}
	}
}

func main() {
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 500
	result := BinarySearch2(array, target, 0, len(array)-1)
	fmt.Println(target, result)

	target = 189
	result = BinarySearch2(array, target, 0, len(array)-1)
	fmt.Println(target, result)
}
```

很多计算机问题都可以用递归来简化求解，理论上，所有的递归方式都可以转化为非递归的方式，只不过使用递归，代码的可读性更高。
