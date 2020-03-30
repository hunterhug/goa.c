package main

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
	"math"
	"sync"
)

const (
	// 扩容因子
	expandFactor = 1.0 / 2.0
)

// 哈希表
type HashMap struct {
	array        []*keyPairs // 哈希表数组，每个元素是一个键值对
	capacity     int         // 数组容量
	len          int         // 已添加键值对元素数量
	capacityMask int         // 掩码，等于 capacity-1
	// 增删键值对时，需要考虑并发安全
	lock sync.Mutex
}

// 键值对，连成一个链表
type keyPairs struct {
	key   string      // 键
	value interface{} // 值
	next  *keyPairs   // 下一个键值对
}

// 创建大小为 capacity 的哈希表
func NewHashMap(capacity int) *HashMap {
	// 默认大小为 16
	defaultCapacity := 1 << 4
	if capacity <= defaultCapacity {
		// 如果传入的大小小于默认大小，那么使用默认大小16
		capacity = defaultCapacity
	} else {
		// 否则，实际大小为大于 capacity 的第一个 2^k
		capacity = 1 << (int(math.Ceil(math.Log2(float64(capacity)))))
	}

	// 新建一个哈希表
	m := new(HashMap)
	m.array = make([]*keyPairs, capacity, capacity)
	m.capacity = capacity
	m.capacityMask = capacity - 1
	return m
}

// 返回哈希表已添加元素数量
func (m *HashMap) Len() int {
	return m.len
}

// 返回哈希表目前的容量
func (m *HashMap) Capacity() int {
	return m.capacity
}

// 求 key 的哈希值
var hashAlgorithm = func(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

// 对键进行哈希求值，并计算下标
func (m *HashMap) hashIndex(key string, mask int) int {
	// 求哈希
	hash := hashAlgorithm([]byte(key))
	// 求下标
	index := hash & uint64(mask)
	return int(index)
}

// 哈希表添加键值对
func (m *HashMap) Put(key string, value interface{}) {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()

	// 键值对要放的哈希表数组下标
	index := m.hashIndex(key, m.capacityMask)

	// 哈希表数组下标的元素
	element := m.array[index]

	// 元素为空，表示空链表，没有哈希冲突，直接赋值
	if element == nil {
		m.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		// 链表最后一个键值对
		var lastPairs *keyPairs

		// 遍历链表查看元素是否存在，存在则替换值，否则找到最后一个键值对
		for element != nil {
			// 键值对存在，那么更新值并返回
			if element.key == key {
				element.value = value
				return
			}

			lastPairs = element
			element = element.next
		}

		// 找不到键值对，将新键值对添加到链表尾端
		lastPairs.next = &keyPairs{
			key:   key,
			value: value,
		}
	}

	// 新的哈希表数量
	newLen := m.len + 1

	// 如果超出扩容因子，需要扩容
	if float64(newLen)/float64(m.capacity) >= expandFactor {
		// 新建一个原来两倍大小的哈希表
		newM := new(HashMap)
		newM.array = make([]*keyPairs, 2*m.capacity, 2*m.capacity)
		newM.capacity = 2 * m.capacity
		newM.capacityMask = 2*m.capacity - 1

		// 遍历老的哈希表，将键值对重新哈希到新哈希表
		for _, pairs := range m.array {
			for pairs != nil {
				// 直接递归Put
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}

		// 替换老的哈希表
		m.array = newM.array
		m.capacity = newM.capacity
		m.capacityMask = newM.capacityMask
	}

	m.len = newLen
}

// 哈希表获取键值对
func (m *HashMap) Get(key string) (value interface{}, ok bool) {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()

	// 键值对要放的哈希表数组下标
	index := m.hashIndex(key, m.capacityMask)

	// 哈希表数组下标的元素
	element := m.array[index]

	// 遍历链表查看元素是否存在，存在则返回
	for element != nil {
		if element.key == key {
			return element.value, true
		}

		element = element.next
	}

	return
}

// 哈希表删除键值对
func (m *HashMap) Delete(key string) {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()

	// 键值对要放的哈希表数组下标
	index := m.hashIndex(key, m.capacityMask)

	// 哈希表数组下标的元素
	element := m.array[index]

	// 空链表，不用删除，直接返回
	if element == nil {
		return
	}

	// 链表的第一个元素就是要删除的元素
	if element.key == key {
		// 将第一个元素后面的键值对链上
		m.array[index] = element.next
		m.len = m.len - 1
		return
	}

	// 下一个键值对
	nextElement := element.next
	for nextElement != nil {
		if nextElement.key == key {
			// 键值对匹配到，将该键值对从链中去掉
			element.next = nextElement.next
			m.len = m.len - 1
			return
		}

		element = nextElement
		nextElement = nextElement.next
	}
}

// 哈希表遍历
func (m *HashMap) Range() {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, pairs := range m.array {
		for pairs != nil {
			fmt.Printf("'%v'='%v',", pairs.key, pairs.value)
			pairs = pairs.next
		}
	}

	fmt.Println()
}

func main() {
	// 新建一个哈希表
	hashMap := NewHashMap(16)

	// 放35个值
	for i := 0; i < 35; i++ {
		hashMap.Put(fmt.Sprintf("%d", i), fmt.Sprintf("v%d", i))
	}
	fmt.Println("cap:", hashMap.Capacity(), "len:", hashMap.Len())

	// 打印全部键值对
	hashMap.Range()

	key := "4"
	value, ok := hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}

	// 删除键
	hashMap.Delete(key)
	fmt.Println("after delete cap:", hashMap.Capacity(), "len:", hashMap.Len())
	value, ok = hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}
}
