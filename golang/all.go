// Golang程序入口的包名必须为 main
package main // import "golang"

// 导入其他地方的包，包通过 go mod 机制寻找
import (
	"fmt"
	"golang/diy"
)

// init函数在main函数之前执行
func init() {
	// 声明并初始化三个值
	var i, j, k = 1, 2, 3
	// 使用格式化包打印
	fmt.Println("init hello world")
	fmt.Println(i, j, k)
}

// 函数，两个数相加
func sum(a, b int64) int64 {
	return a + b
}

// 程序入口必须为 main 函数
func main() {
	// 未使用的变量，不允许声明
	//cannot := 6

	fmt.Println("hello world")

	// 定义基本数据类型
	a := 3                                // int
	b := 6.0                              // float64
	c := "hi"                             // string
	d := [3]int64{1, 2, 3}                // array，基本不用到
	e := []int64{1, 2, 3}                 // slice
	f := map[string]int64{"a": 3, "b": 4} // map
	fmt.Printf("type:%T:%v\n", a, a)
	fmt.Printf("type:%T:%v\n", b, b)
	fmt.Printf("type:%T:%v\n", c, c)
	fmt.Printf("type:%T:%v\n", d, d)
	fmt.Printf("type:%T:%v\n", e, e)
	fmt.Printf("type:%T:%v\n", f, f)

	// 切片增加值
	e = append(e, 3)

	// 增加map键值
	f["f"] = 5

	// 查找map键值
	v, ok := f["f"]
	fmt.Println(v, ok)
	v, ok = f["ff"]
	fmt.Println(v, ok)

	// 判断语句
	if a > 0 {
		fmt.Println("a>0")
	} else {
		fmt.Println("a<=0")
	}

	// 循环语句
	for {
		if true {
			fmt.Println("for")
			// 退出循环
			break
		}
	}

	// 循环语句
	for i := 9; i <= 10; i++ {
		fmt.Printf("i=%d\n", i)
	}

	// 循环切片
	for k, v := range e {
		fmt.Println(k, v)
	}

	// 循环map
	for k, v := range f {
		fmt.Println(k, v)
	}

	// 定义 int64 变量
	var h, i int64 = 4, 6

	// 使用函数
	sum := sum(h, i)
	fmt.Printf("sum(h+i),h=%v,i=%v,%v\n", h, i, sum)

	// 新建结构体，值
	g := diy.Diy{
		A: 2,
		//b: 4.0, // 小写成员不能导出
	}

	// 打印类型，值
	fmt.Printf("type:%T:%v\n", g, g)

	// 小写方法不能导出
	//g.set(1,1)
	g.Set(1, 1)
	fmt.Printf("type:%T:%v\n", g, g) // 结构体值变化

	g.Set2(3, 3)
	fmt.Printf("type:%T:%v\n", g, g) // 结构体值未变化

	// 新建结构体，引用
	k := &diy.Diy{
		A: 2,
	}
	fmt.Printf("type:%T:%v\n", k, k)
	k.Set(1, 1)
	fmt.Printf("type:%T:%v\n", k, k) // 结构体值变化
	k.Set2(3, 3)
	fmt.Printf("type:%T:%v\n", k, k) // 结构体值未变化

	// 新建结构体，引用
	m := new(diy.Diy)
	m.A = 2
	fmt.Printf("type:%T:%v\n", m, m)
}
