package main

import "fmt"

// 普通红黑树实现
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
	Value int64    // 值
	Times int64    // 值出现的次数
	Left  *RBTNode // 左子树
	Right *RBTNode // 右子树
	Color bool     // 父亲指向该节点的链接颜色
}

// 节点的颜色
func IsRed(node *RBTNode) bool {
	if node == nil {
		return false
	}
	return node.Color == RED
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

// 红色左移
// 节点 h 是红节点，其左儿子和左儿子的左儿子都为黑节点，左移后使得其左儿子或左儿子的左儿子有一个是红色节点
func MoveRedLeft(h *RBTNode) *RBTNode {
	// 应该确保 isRed(h) && !isRed(h.left) && !isRed(h.left.left)
	ColorChange(h)

	// 右儿子有左红链接
	if IsRed(h.Right.Left) {
		// 对右儿子右旋
		h.Right = RotateRight(h.Right)
		// 再左旋
		h = RotateLeft(h)
		ColorChange(h)
	}

	return h
}

// 红色右移
// 节点 h 是红节点，其右儿子和右儿子的左儿子都为黑节点，右移后使得其右儿子或右儿子的右儿子有一个是红色节点
func MoveRedRight(h *RBTNode) *RBTNode {
	// 应该确保 isRed(h) && !isRed(h.right) && !isRed(h.right.left);
	ColorChange(h)

	// 左儿子有左红链接
	if IsRed(h.Left.Left) {
		// 右旋
		h = RotateRight(h)
		// 变色
		ColorChange(h)
	}

	return h
}

// 颜色变换
func ColorChange(h *RBTNode) {
	if h == nil {
		return
	}
	h.Color = !h.Color
	h.Left.Color = !h.Left.Color
	h.Right.Color = !h.Right.Color
}

// 普通红黑树添加元素
func (tree *RBTree) Add(value int64) {
	// 跟节点开始添加元素，因为可能调整，所以需要将返回的节点赋值回根节点
	tree.Root = tree.Root.Add(value)
	// 根节点的链接永远都是黑色的
	tree.Root.Color = BLACK
}

// 往节点添加元素
func (node *RBTNode) Add(value int64) *RBTNode {
	// 插入的节点为空，将其链接颜色设置为红色，并返回
	if node == nil {
		return &RBTNode{
			Value: value,
			Color: RED,
		}
	}

	// 插入的元素重复
	if value == node.Value {
		node.Times = node.Times + 1
	} else if value > node.Value {
		// 插入的元素比节点值大，往右子树插入
		node.Right = node.Right.Add(value)
	} else {
		// 插入的元素比节点值小，往左子树插入
		node.Left = node.Left.Add(value)
	}

	// 辅助变量
	nowNode := node

	// 右链接为红色，那么进行左旋，确保树是左倾的
	// 这里做完操作后就可以结束了，因为插入操作，新插入的右红链接左旋后，nowNode节点不会出现连续两个红左链接，因为它只有一个左红链接
	if IsRed(nowNode.Right) && !IsRed(nowNode.Left) {
		nowNode = RotateLeft(nowNode)
	} else {
		// 连续两个左链接为红色，那么进行右旋
		if IsRed(nowNode.Left) && IsRed(nowNode.Left.Left) {
			nowNode = RotateRight(nowNode)
		}

		// 旋转后，可能左右链接都为红色，需要变色
		if IsRed(nowNode.Left) && IsRed(nowNode.Right) {
			ColorChange(nowNode)
		}
	}

	return nowNode
}

// 找出最小值的节点
func (tree *RBTree) FindMinValue() *RBTNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.FindMinValue()
}

func (node *RBTNode) FindMinValue() *RBTNode {
	// 左子树为空，表面已经是最左的节点了，该值就是最小值
	if node.Left == nil {
		return node
	}

	// 一直左子树递归
	return node.Left.FindMinValue()
}

// 找出最大值的节点
func (tree *RBTree) FindMaxValue() *RBTNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.FindMaxValue()
}

func (node *RBTNode) FindMaxValue() *RBTNode {
	// 右子树为空，表面已经是最右的节点了，该值就是最大值
	if node.Right == nil {
		return node
	}

	// 一直右子树递归
	return node.Right.FindMaxValue()
}

// 查找指定节点
func (tree *RBTree) Find(value int64) *RBTNode {
	return tree.Root.Find(value)
}

func (node *RBTNode) Find(value int64) *RBTNode {
	if value == node.Value {
		// 如果该节点刚刚等于该值，那么返回该节点
		return node
	} else if value < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil {
			// 左子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Left.Find(value)
	} else {
		// 如果查找的值大于节点值，从节点的右子树开始找
		if node.Right == nil {
			// 右子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Right.Find(value)
	}
}

// 中序遍历
func (tree *RBTree) MidOrder() {
	tree.Root.MidOrder()
}

func (node *RBTNode) MidOrder() {
	if node == nil {
		return
	}

	// 先打印左子树
	node.Left.MidOrder()

	// 按照次数打印根节点
	for i := 0; i <= int(node.Times); i++ {
		fmt.Println(node.Value)
	}

	// 打印右子树
	node.Right.MidOrder()
}

// 修复普通红黑树特征
func (node *RBTNode) FixUp() *RBTNode {
	// 辅助变量hunterhugxx-6833066
	nowNode := node

	// 红链接在右边，左旋恢复，让红链接只出现在左边
	if IsRed(nowNode.Right) {
		nowNode = RotateLeft(nowNode)
	}

	// 连续两个左链接为红色，那么进行右旋
	if IsRed(nowNode.Left) && IsRed(nowNode.Left.Left) {
		nowNode = RotateRight(nowNode)
	}

	// 旋转后，可能左右链接都为红色，需要变色
	if IsRed(nowNode.Left) && IsRed(nowNode.Right) {
		ColorChange(nowNode)
	}

	return nowNode
}

// 对该节点所在的子树删除最小元素
func (node *RBTNode) DeleteMin() *RBTNode {
	// 辅助变量
	nowNode := node

	// 没有左子树，那么删除它自己
	if nowNode.Left == nil {
		return nil
	}

	// 判断是否需要红色左移，因为最小元素在左子树中
	if !IsRed(nowNode.Left) && !IsRed(nowNode.Left.Left) {
		nowNode = MoveRedLeft(nowNode)
	}

	// 递归从左子树删除
	nowNode.Left = nowNode.Left.DeleteMin()

	// 修复普通红黑树特征
	return nowNode.FixUp()
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

// 对该节点所在的子树删除元素
func (node *RBTNode) Delete(value int64) *RBTNode {
	// 辅助变量
	nowNode := node
	// 删除的元素比子树根节点小，需要从左子树删除
	if value < nowNode.Value {
		// 因为从左子树删除，所以要判断是否需要红色左移
		if !IsRed(nowNode.Left) && !IsRed(nowNode.Left.Left) {
			// 左儿子和左儿子的左儿子都不是红色节点，那么没法递归下去，先红色左移
			nowNode = MoveRedLeft(nowNode)
		}

		// 现在可以从左子树中删除了
		nowNode.Left = nowNode.Left.Delete(value)
	} else {
		// 删除的元素等于或大于树根节点

		// 左节点为红色，那么需要右旋，方便后面可以红色右移
		if IsRed(nowNode.Left) {
			nowNode = RotateRight(nowNode)
		}

		// 值相等，且没有右孩子节点，那么该节点一定是要被删除的叶子节点，直接删除
		// 为什么呢，反证，它没有右儿子，但有左儿子，因为普通红黑树的特征，那么左儿子一定是红色，但是前面的语句已经把红色左儿子右旋到右边，不应该出现右儿子为空。
		if value == nowNode.Value && nowNode.Right == nil {
			return nil
		}

		// 因为从右子树删除，所以要判断是否需要红色右移
		if !IsRed(nowNode.Right) && !IsRed(nowNode.Right.Left) {
			// 右儿子和右儿子的左儿子都不是红色节点，那么没法递归下去，先红色右移
			nowNode = MoveRedRight(nowNode)
		}

		// 删除的节点找到了，它是中间节点，需要用最小后驱节点来替换它，然后删除最小后驱节点
		if value == nowNode.Value {
			minNode := nowNode.Right.FindMinValue()
			nowNode.Value = minNode.Value
			nowNode.Times = minNode.Times

			// 删除其最小后驱节点
			nowNode.Right = nowNode.Right.DeleteMin()
		} else {
			// 删除的元素比子树根节点大，需要从右子树删除
			nowNode.Right = nowNode.Right.Delete(value)
		}
	}

	// 最后，删除叶子节点后，需要恢复普通红黑树特征
	return nowNode.FixUp()
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
