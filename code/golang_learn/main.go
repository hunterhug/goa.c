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
	p := true                             // bool
	a := 3                                // int
	b := 6.0                              // float64
	c := "hi"                             // string
	d := [3]string{"1", "2", "3"}         // array，基本不用到
	e := []int64{1, 2, 3}                 // slice
	f := map[string]int64{"a": 3, "b": 4} // map
	fmt.Printf("type:%T:%v\n", p, p)
	fmt.Printf("type:%T:%v\n", a, a)
	fmt.Printf("type:%T:%v\n", b, b)
	fmt.Printf("type:%T:%v\n", c, c)
	fmt.Printf("type:%T:%v\n", d, d)
	fmt.Printf("type:%T:%v\n", e, e)
	fmt.Printf("type:%T:%v\n", f, f)

	// 切片放值
	e[0] = 9
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

	s := make([]int64, 5)
	s1 := make([]int64, 0, 5)
	m1 := make(map[string]int64, 5)
	m2 := make(map[string]int64)
	fmt.Printf("%#v,cap:%#v,len:%#v\n", s, cap(s), len(s))
	fmt.Printf("%#v,cap:%#v,len:%#v\n", s1, cap(s1), len(s1))
	fmt.Printf("%#v,len:%#v\n", m1, len(m1))
	fmt.Printf("%#v,len:%#v\n", m2, len(m2))

	var ll []int64
	fmt.Printf("%#v\n", ll)
	ll = append(ll, 1)
	fmt.Printf("%#v\n", ll)
	ll = append(ll, 2, 3, 4, 5, 6)
	fmt.Printf("%#v\n", ll)
	ll = append(ll, []int64{7, 8, 9}...)
	fmt.Printf("%#v\n", ll)

	fmt.Println(ll[0:2])
	fmt.Println(ll[:2])
	fmt.Println(ll[0:])
	fmt.Println(ll[:])

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

	num := 4
	switch num {
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("not found")

	}
}
