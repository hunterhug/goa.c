# 使用标准库

## 一、避免重复造轮子

官方提供了很多库给我们用，是封装好的轮子，比如包 `fmt`，我们多次使用它来打印数据。

我们可以查看到其里面的实现：

```go
package fmt

func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintln(a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```

函数 `Println` 是直接打印并换行，`Printf` 的作用是格式化输出，如：

```go
	// 打印一行空行
	fmt.Println()
	// 打印 4 5 6
	fmt.Println(4, 5, 6)
	// 占位符 %d 打印数字，\n换行
	fmt.Printf("%d\n", 2)
	// 占位符 %s 打印字符串，\n换行
	fmt.Printf("%s\n", "cat")
	// 占位符 %v或者%#v 打印任何类型，\n换行
	fmt.Printf("%#v,%v\n", "cat", 33)
	// 更多示例
	fmt.Printf("%s,%d,%s,%#v\n", "cat", 2, "3", map[int]string{1: "s"})
```

输出：

```go
4 5 6
2
cat
"cat",33
cat,2,3,map[int]string{1:"s"}
```

函数 `Printf` 使用到了另外一个函数，而函数又使用了另外的函数。

我们在某些时候，可以使用官方库或别人写的库，毕竟轮子重造需要时间。

同时，如果想开发速度提高，建议安装 IDE(集成开发环境)，如 `Goland`，请自行百度。

## 二、总结

我们只学习了 `Golang` 语言的一个子集，很多关键特征如 `interface` 还有 `go func()` 以及 `chan` 都没有讲解。

想更详细的学习，可以安装 `docker` 后，打开终端：

```
# 拉镜像
docker pull hunterhug/gotourzh

# 后台运行
docker run -d -p 9999:9999 hunterhug/gotourzh
```

打开浏览器输入：[127.0.0.1:9999](http://127.0.0.1:9999) 更全面地学习。

后面的算法分析和实现，会使用 `Golang` 来举例。