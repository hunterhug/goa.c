package main

import "fmt"

// 普通红黑树实现，参考 Java TreeMap，更强壮。
// red-black tree

// 定义颜色
const (
	RED   = true
	BLACK = false
)

// 普通红黑树
type RBTree struct {
	Root *RBTNode // 树根节点
}

// 新建一棵空树
func NewRBTree() *RBTree {
	return &RBTree{}
}

// 普通红黑树节点
type RBTNode struct {
	Value  int64    // 值
	Times  int64    // 值出现的次数
	Left   *RBTNode // 左子树
	Right  *RBTNode // 右子树
	Parent *RBTNode // 父节点
	Color  bool     // 父亲指向该节点的链接颜色
}

// 节点的颜色
func IsRed(node *RBTNode) bool {
	if node == nil {
		return false
	}
	return node.Color == RED
}

// 返回节点的父亲节点
func ParentOf(node *RBTNode) *RBTNode {
	if node == nil {
		return nil
	}

	return node.Parent
}

// 返回节点的左子节点
func LeftOf(node *RBTNode) *RBTNode {
	if node == nil {
		return nil
	}

	return node.Left
}

// 返回节点的右子节点
func RightOf(node *RBTNode) *RBTNode {
	if node == nil {
		return nil
	}

	return node.Right
}

// 设置节点颜色
func SetColor(node *RBTNode, color bool) {
	if node != nil {
		node.Color = color
	}
}

// 左旋转
func RotateLeft(h *RBTNode) *RBTNode {
	if h == nil {
		return nil
	}

	// 看图理解
	x := h.Right
	h.Right = x.Left
	x.Left = h
	return x
}

// 右旋转
func RotateRight(h *RBTNode) *RBTNode {
	if h == nil {
		return nil
	}

	// 看图理解
	x := h.Left
	h.Left = x.Right
	x.Right = h
	return x
}

// 普通红黑树添加元素
func (tree *RBTree) Add(value int64) {
	// 根节点为空
	if tree.Root == nil {
		// 根节点都是黑色
		tree.Root = &RBTNode{
			Value: value,
			Color: BLACK,
		}
		return
	}

	// 辅助变量 t，表示新元素要插入到该子树，t是该子树的根节点
	t := tree.Root

	// 插入元素后，插入元素的父亲节点
	var parent *RBTNode

	// 辅助变量，为了知道元素最后要插到左边还是右边
	var cmp int64 = 0

	for {
		parent = t

		cmp = value - t.Value
		if cmp < 0 {
			// 比当前节点小，往左子树插入
			t = t.Left
		} else if cmp > 0 {
			// 比当前节点节点大，往右子树插入
			t = t.Right
		} else {
			// 已经存在值了，更新出现的次数
			t.Times = t.Times + 1
			return
		}

		// 终于找到要插入的位置了
		if t == nil {
			break // 这时叶子节点是 parent，要插入到 parent 的下面，跳到外层去
		}
	}

	// 新节点，它要插入到 parent下面
	newNode := &RBTNode{
		Value:  value,
		Parent: parent,
	}
	if cmp < 0 {
		// 知道要从左边插进去
		parent.Left = newNode
	} else {
		// 知道要从右边插进去
		parent.Right = newNode
	}

	// 插入新节点后，可能破坏了红黑树特征，需要修复，核心函数
	tree.fixAfterInsertion(newNode)
}

// 调整新插入的节点，自下而上
func (tree *RBTree) fixAfterInsertion(node *RBTNode) {
	// 插入的新节点一定要是红色
	node.Color = RED

	// 节点不能是空，不能是根节点，父亲的颜色必须为红色（如果是黑色，那么直接插入不破坏平衡，不需要调整了）
	for node != nil && node != tree.Root && node.Parent.Color == RED {
		// 父亲在祖父的左边
		if ParentOf(node) == LeftOf(ParentOf(ParentOf(node)))
	}

	// 根节点永远为黑
	tree.Root.Color = BLACK
}

// 普通红黑树删除元素
func (tree *RBTree) Delete(value int64) {
	// 当找不到值时直接返回
	if tree.Find(value) == nil {
		return
	}

	if !IsRed(tree.Root.Left) && !IsRed(tree.Root.Right) {
		// 左右子树都是黑节点，那么先将根节点变为红节点，方便后面的红色左移或右移
		tree.Root.Color = RED
	}

	tree.Root = tree.Root.Delete(value)

	// 最后，如果根节点非空，永远都要为黑节点，赋值黑色
	if tree.Root != nil {
		tree.Root.Color = BLACK
	}
}

func main() {
	tree := NewRBTree()
	values := []int64{2, 3, 7, 10, 10, 10, 10, 23, 9, 102, 109, 111, 112, 113, 115, 18}
	for _, v := range values {
		tree.Add(v)
	}

	// 找到最大值或最小值的节点
	fmt.Println("find min value:", tree.FindMinValue())
	fmt.Println("find max value:", tree.FindMaxValue())

	// 查找不存在的99
	node := tree.Find(99)
	if node != nil {
		fmt.Println("find it 99!")
	} else {
		fmt.Println("not find it 99!")
	}

	// 查找存在的9
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}

	tree.MidOrder()

	tree.Delete(9)

	// 查找存在的9
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}
}
