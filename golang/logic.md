# 流程控制语句

`Golang` 只有一种判断和一种循环：`if` 和 `for`。

判断语句如：

```go
	// 判断语句
	if a > 0 {
		fmt.Println("a>0")
	} else {
		fmt.Println("a<=0")
	}
```

判断 `a > 0` 条件不需要加小括号。

循环语句：

```go
	// 循环语句
	for i := 9; i <= 10; i++ {
		fmt.Printf("i=%d\n", i)
	}
```

其中 `i` 是局部变量，循环第一次前被赋予了值 `9`，然后判断是否满足 `i<=10` 条件，如果满足那么进入循环打印，每一次循环后会加`1`，也就是 `i++`，然后继续判断是否满足条件。

你也可以死循环：

```go
	// 死循环语句
	a = 0
	for {
		if a >= 10 {
			fmt.Println("out")
			// 退出循环
			break
		}

		a = a + 1
		if a > 5 {
			continue
		} 
		
        fmt.Println(a)
	}
```

死循环直接 `for`，后面不需要加条件，然后当 `a>=10` 时跳出循环可以使用 `break`，表示跳出 `for {}`，对于 `a > 5`，我们不想打印出值，可以使用 `continue` 跳过后面的语句 `fmt.Println(a)`，提前再一次进入循环。

`切片` 和 `字典` 都可以使用循环来遍历数据:

```go

	e := []int64{1, 2, 3}                 // slice
	f := map[string]int64{"a": 3, "b": 4} // map

	// 循环切片
	for k, v := range e {
		fmt.Println(k, v)
	}

	// 循环map
	for k, v := range f {
		fmt.Println(k, v)
	}
```

切片遍历出来的结果为：数据下标，数据，字典遍历出来的结果为：数据的键，数据的值:

```
0 1
1 2
2 3
3 3

a 3
b 4
f 5
```