// 包名
package diy

// 结构体
type Diy struct {
	A int64   // 大写导出成员
	b float64 // 小写不可以导出
}

// 引用结构体的方法，引用传递，会改变原有结构体的值
func (diy *Diy) Set(a int64, b float64) {
	diy.A = a
	diy.b = b
	return
}

// 值结构体的方法，值传递，不会改变原有结构体的值
func (diy Diy) Set2(a int64, b float64) {
	diy.A = a
	diy.b = b
	return
}

// 小写方法，不能导出
func (diy Diy) set(a int64, b float64) {
	diy.A = a
	diy.b = b
	return
}

// 小写函数，不能导出，只能在同一包下使用
func sum(a, b int64) int64 {
	return a + b
}
