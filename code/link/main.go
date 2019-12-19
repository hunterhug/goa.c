package main

import (
	"fmt"
	"sync"
)


func main() {
	array := [5]int64{}
	fmt.Println(array)
	array[0] = 8
	array[1] = 9
	array [2] = 7
	fmt.Println(array)
	fmt.Println(array[2])

	ArrayLink()
	SetTest()
}

func ArrayLink() {
	type Value struct {
		Data      string
		NextIndex int64
	}

	var array [5]Value          // 五个节点的数组
	array[0] = Value{"I", 3}    // 下一个节点的下标为3
	array[1] = Value{"Army", 4} // 下一个节点的下标为4
	array[2] = Value{"You", 1}  // 下一个节点的下标为1
	array[3] = Value{"Love", 2} // 下一个节点的下标为2
	array[4] = Value{"!", -1}   // -1表示没有下一个节点
	node := array[0]
	for {
		fmt.Println(node.Data)
		if node.NextIndex == -1 {
			break
		}
		node = array[node.NextIndex]
	}

}

// 集合结构体
type Set struct {
	m            map[int]bool // 用字典来实现，因为字段键不能重复
	len          int          // 集合的大小
	sync.RWMutex              // 锁，实现并发安全
}

// 新建一个空集合
func New() *Set {
	return &Set{
		m: map[int]bool{},
	}
}

// 增加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true  // 实际往字典添加这个键
	s.len = s.len + 1 // 集合大小增加
}

// 移除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	s.Unlock()
	delete(s.m, item) // 实际从字典删除这个键
	s.len = s.len - 1 // 集合大小减少
}

// 查看是否存在元素
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// 查看集合大小
func (s *Set) Len() int {
	return s.len
}

// 清除集合所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{} // 字典置空
	s.len = 0            // 大小归零
}

// 集合是够为空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// 将集合转化为列表
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func SetTest() {
	// 初始化
	s := New()

	s.Add(1)
	s.Add(1)
	s.Add(2)

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("empty")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		fmt.Println("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items", s.List())
}
