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
	fmt.Printf("%s,%d,%s,%v,%#v\n", "cat", 2, "3", map[int]string{1: "s"}, map[int]string{1: "s"})
```

输出：

```go
4 5 6
2
cat
"cat",33
cat,2,3,map[1:s],map[int]string{1:"s"}
```

函数 `Printf` 使用到了另外一个函数 `Fprintf`，而函数 `Fprintf` 内部又调用了其他的结构体方法。

对于我们经常使用的 `func Printf(format string, a ...interface{})`，我们传入 `format` 和许多变量 `a ...interface{}`，就可以在控制台打印出我们想要的结果。如：

```go
fmt.Printf("%s,%d,%s,%v,%#v\n", "cat", 2, "3", map[int]string{1: "s"}, map[int]string{1: "s"})
```

其中 `%` 是占位符，表示后面的变量逐个占位。占位符后面的小写字母表示占位的类型，`%s` 表示字符串的占位，`%d` 表示数字类型的占位, `%v` 或 `%#v` 表示未知类型的占位，会自动判断类型后打印，加 `#` 会打印得更详细一点。因为该打印不会换行，我们需要使用 `\n` 换行符来换行。


在某些时候，我们可以使用官方库或别人写的库，毕竟轮子重造需要时间。

同时，如果想开发速度提高，建议安装 `IDE`，也就是 `Integrated Development Environment` (集成开发环境)，如 `Goland`(原生支持 `Golang`) 或 `IDEA` 软件(需安装插件)。

## 二、总结

我们只学习了 `Golang` 语言的一个子集，想更详细的学习，可以安装 `docker` 后，打开终端：

```
# 拉镜像
docker pull hunterhug/gotourzh

# 后台运行
docker run -d -p 9999:9999 hunterhug/gotourzh
```

打开浏览器输入：[http://127.0.0.1:9999](http://127.0.0.1:9999) 更全面地学习。

后面的算法分析和实现，会使用 `Golang` 来举例。
