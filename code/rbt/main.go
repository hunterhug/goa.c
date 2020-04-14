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

// 对某节点左旋转
func (tree *RBTree) RotateLeft(h *RBTNode) {
	if h != nil {

		// 看图理解
		x := h.Right
		h.Right = x.Left

		if x.Left != nil {
			x.Left.Parent = h
		}

		x.Parent = h.Parent
		if h.Parent == nil {
			tree.Root = x
		} else if h.Parent.Left == h {
			h.Parent.Left = x
		} else {
			h.Parent.Right = x
		}
		x.Left = h
		h.Parent = x
	}
}

// 对某节点右旋转
func (tree *RBTree) RotateRight(h *RBTNode) {
	if h != nil {

		// 看图理解
		x := h.Left
		h.Left = x.Right

		if x.Right != nil {
			x.Right.Parent = h
		}

		x.Parent = h.Parent
		if h.Parent == nil {
			tree.Root = x
		} else if h.Parent.Right == h {
			h.Parent.Right = x
		} else {
			h.Parent.Left = x
		}
		x.Right = h
		h.Parent = x
	}
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

// 调整新插入的节点，自底而上
// 可以看图理解
func (tree *RBTree) fixAfterInsertion(node *RBTNode) {
	// 插入的新节点一定要是红色
	node.Color = RED

	// 节点不能是空，不能是根节点，父亲的颜色必须为红色（如果是黑色，那么直接插入不破坏平衡，不需要调整了）
	for node != nil && node != tree.Root && node.Parent.Color == RED {
		// 父亲在祖父的左边
		if ParentOf(node) == LeftOf(ParentOf(ParentOf(node))) {
			// 叔叔节点
			uncle := RightOf(ParentOf(ParentOf(node)))

			// 图例3左边部分，叔叔是红节点，祖父变色，也就是父亲和叔叔变黑，祖父变红
			if IsRed(uncle) {
				SetColor(ParentOf(node), BLACK)
				SetColor(uncle, BLACK)
				SetColor(ParentOf(ParentOf(node)), RED)
				// 还要向上递归
				node = ParentOf(ParentOf(node))
			} else {
				// 图例4左边部分，叔叔是黑节点，并且插入的节点在父亲的右边，需要对父亲左旋
				if node == RightOf(ParentOf(node)) {
					node = ParentOf(node)
					tree.RotateLeft(node)
				}

				// 变色，并对祖父进行右旋
				SetColor(ParentOf(node), BLACK)
				SetColor(ParentOf(ParentOf(node)), RED)
				tree.RotateRight(ParentOf(ParentOf(node)))
			}
		} else {
			// 父亲在祖父的右边，与父亲在祖父的左边相似
			// 叔叔节点
			uncle := LeftOf(ParentOf(ParentOf(node)))

			// 图例3右边部分，叔叔是红节点，祖父变色，也就是父亲和叔叔变黑，祖父变红
			if IsRed(uncle) {
				SetColor(ParentOf(node), BLACK)
				SetColor(uncle, BLACK)
				SetColor(ParentOf(ParentOf(node)), RED)
				// 还要向上递归
				node = ParentOf(ParentOf(node))
			} else {
				// 图例4右边部分，叔叔是黑节点，并且插入的节点在父亲的左边，需要对父亲右旋
				if node == LeftOf(ParentOf(node)) {
					node = ParentOf(node)
					tree.RotateLeft(node)
				}

				// 变色，并对祖父进行左旋
				SetColor(ParentOf(node), BLACK)
				SetColor(ParentOf(ParentOf(node)), RED)
				tree.RotateLeft(ParentOf(ParentOf(node)))
			}
		}
	}

	// 根节点永远为黑
	tree.Root.Color = BLACK
}

// 普通红黑树删除元素
func (tree *RBTree) Delete(value int64) {
	// 查找元素是否存在，不存在则退出
	p := tree.Find(value)
	if p == nil {
		return
	}

	// 删除该节点
	tree.delete(p)
}

// 删除节点核心函数
// 找最小后驱节点来补位，删除内部节点转为删除叶子节点
func (tree *RBTree) delete(node *RBTNode) {
	// 如果左右子树都存在，那么从右子树的左边一直找一直找，就找能到最小后驱节点
	if node.Left != nil && node.Right != nil {
		s := node.Right
		for s.Left != nil {
			s = s.Left
		}

		// 删除的叶子节点找到了，删除内部节点转为删除叶子节点
		node.Value = s.Value
		node.Times = s.Times
		node = s
	} else if node.Left == nil && node.Right == nil {
		// 没有子树，要删除的节点就是叶子节点。
	} else {
		// 只有一棵子树，因为红黑树的特征，该子树就只有一个节点
		// 找到该唯一节点
		replacement := node.Left
		if node.Left == nil {
			replacement = node.Right
		}

		// 替换开始，子树的唯一节点替代被删除的内部节点
		replacement.Parent = node.Parent
		if node.Parent == nil {
			// 要删除的节点的父亲为空，表示要删除的节点为根节点，唯一子节点成为树根
			tree.Root = replacement
			// 根节点永远都是黑色
			tree.Root.Color = BLACK
		} else if node == node.Parent.Left {
			// 子树的唯一节点替代被删除的内部节点
			node.Parent.Left = replacement
		} else {
			// 子树的唯一节点替代被删除的内部节点
			node.Parent.Right = replacement
		}

		// 子树的该唯一节点一定是一个红节点，不然破坏红黑树特征，所以替换后可以直接返回
		return
	}

	// 要删除的叶子节点没有父亲，那么它是根节点，直接置空，返回
	if node.Parent == nil {
		tree.Root = nil
		return
	}

	// 要删除的叶子节点，是一个黑节点，删除后会破坏平衡，需要进行调整，调整成可以删除的状态
	if !IsRed(node) {
		// 核心函数
		tree.fixAfterDeletion(node)
	}

	// 现在可以删除叶子节点了
	if node == node.Parent.Left {
		node.Parent.Left = nil
	} else if node == node.Parent.Right {
		node.Parent.Right = nil
	}

}

// 调整删除的叶子节点，自底向上
// 可以看图理解
func (tree *RBTree) fixAfterDeletion(node *RBTNode) {
	// 如果不是递归到根节点，且节点是黑节点，那么继续递归
	for tree.Root != node && !IsRed(node) {
		// 要删除的节点在父亲左边，对应图例1，2
		if node == LeftOf(ParentOf(node)) {
			// 找出兄弟
			brother := RightOf(ParentOf(node))

			// 兄弟是红色的，对应图例1，那么兄弟变黑，父亲变红，然后对父亲左旋，进入图例23
			if IsRed(brother) {
				SetColor(brother, BLACK)
				SetColor(ParentOf(node), RED)
				tree.RotateLeft(ParentOf(node))
				brother = RightOf(ParentOf(node)) // 图例1调整后进入图例23，兄弟此时变了
			}

			// 兄弟是黑色的，对应图例21，22，23
			// 兄弟的左右儿子都是黑色，进入图例23，将兄弟设为红色，父亲所在的子树作为整体，当作删除的节点，继续向上递归
			if !IsRed(LeftOf(brother)) && !IsRed(RightOf(brother)) {
				SetColor(brother, RED)
				node = ParentOf(node)
			} else {
				// 兄弟的右儿子是黑色，进入图例22，将兄弟设为红色，兄弟的左儿子设为黑色，对兄弟右旋，进入图例21
				if !IsRed(RightOf(brother)) {
					SetColor(LeftOf(brother), BLACK)
					SetColor(brother, RED)
					tree.RotateRight(brother)
					brother = RightOf(ParentOf(node)) // 图例22调整后进入图例21，兄弟此时变了
				}

				// 兄弟的右儿子是红色，进入图例21，将兄弟设置为父亲的颜色，兄弟的右儿子以及父亲变黑，对父亲左旋
				SetColor(brother, IsRed(ParentOf(node)))
				SetColor(ParentOf(node), BLACK)
				SetColor(RightOf(brother), BLACK)
				tree.RotateLeft(ParentOf(node))

				// 可以返回删除叶子节点了
				return
			}
		} else {
			// 要删除的节点在父亲右边，对应图例3，4
			// 找出兄弟
			brother := RightOf(ParentOf(node))

			// 兄弟是红色的，对应图例3，那么兄弟变黑，父亲变红，然后对父亲右旋，进入图例43
			if IsRed(brother) {
				SetColor(brother, BLACK)
				SetColor(ParentOf(node), RED)
				tree.RotateRight(ParentOf(node))
				brother = LeftOf(ParentOf(node)) // 图例3调整后进入图例43，兄弟此时变了
			}

			// 兄弟是黑色的，对应图例41，42，43
			// 兄弟的左右儿子都是黑色，进入图例43，将兄弟设为红色，父亲所在的子树作为整体，当作删除的节点，继续向上递归
			if !IsRed(LeftOf(brother)) && !IsRed(RightOf(brother)) {
				SetColor(brother, RED)
				node = ParentOf(node)
			} else {
				// 兄弟的左儿子是黑色，进入图例42，将兄弟设为红色，兄弟的右儿子设为黑色，对兄弟左旋，进入图例41
				if !IsRed(LeftOf(brother)) {
					SetColor(RightOf(brother), BLACK)
					SetColor(brother, RED)
					tree.RotateLeft(brother)
					brother = LeftOf(ParentOf(node)) // 图例42调整后进入图例41，兄弟此时变了
				}

				// 兄弟的左儿子是红色，进入图例41，将兄弟设置为父亲的颜色，兄弟的左儿子以及父亲变黑，对父亲右旋
				SetColor(brother, IsRed(ParentOf(node)))
				SetColor(ParentOf(node), BLACK)
				SetColor(LeftOf(brother), BLACK)
				tree.RotateRight(ParentOf(node))

				// 可以返回删除叶子节点了
				return
			}
		}
	}

	// 树根节点永远为黑
	tree.Root.Color = BLACK
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
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

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

// 验证是不是棵红黑树
func (tree *RBTree) IsRBTree() bool {
	if tree == nil || tree.Root == nil {
		return true
	}

	// 判断树是否是一棵二分查找树
	if !tree.Root.IsBST() {
		return false
	}

	// 判断树是否遵循2-3-4树，也就是不能有连续的两个红链接
	if !tree.Root.Is234() {
		return false
	}

	// 判断树是否平衡，也就是任意一个节点到叶子节点，经过的黑色链接数量相同
	// 先计算根节点到最左边叶子节点的黑链接数量
	blackNum := 0
	x := tree.Root
	for x != nil {
		if !IsRed(x) { // 是黑色链接
			blackNum = blackNum + 1
		}
		x = x.Left
	}

	if !tree.Root.IsBalanced(blackNum) {
		return false
	}
	return true
}

// 节点所在的子树是否是一棵二分查找树
func (node *RBTNode) IsBST() bool {
	if node == nil {
		return true
	}

	// 左子树非空，那么根节点必须大于左儿子节点
	if node.Left != nil {
		if node.Value > node.Left.Value {
		} else {
			fmt.Printf("father:%#v,lchild:%#v,rchild:%#v\n", node, node.Left, node.Right)
			return false
		}
	}

	// 右子树非空，那么根节点必须小于右儿子节点
	if node.Right != nil {
		if node.Value < node.Right.Value {
		} else {
			fmt.Printf("father:%#v,lchild:%#v,rchild:%#v\n", node, node.Left, node.Right)
			return false
		}
	}

	// 左子树也要判断是否是平衡查找树
	if !node.Left.IsBST() {
		return false
	}

	// 右子树也要判断是否是平衡查找树
	if !node.Right.IsBST() {
		return false
	}

	return true
}

// 节点所在的子树是否遵循2-3-4树
func (node *RBTNode) Is234() bool {
	if node == nil {
		return true
	}

	// 不允许连续两个左红链接
	if IsRed(node) && IsRed(node.Left) {
		fmt.Printf("father:%#v,lchild:%#v\n", node, node.Left)
		return false
	}

	if IsRed(node) && IsRed(node.Right) {
		fmt.Printf("father:%#v,rchild:%#v\n", node, node.Right)
		return false
	}

	// 左子树也要判断是否遵循2-3-4树
	if !node.Left.Is234() {
		return false
	}

	// 右子树也要判断是否是遵循2-3-4树
	if !node.Right.Is234() {
		return false
	}

	return true
}

// 节点所在的子树是否平衡，是否有 blackNum 个黑链接
func (node *RBTNode) IsBalanced(blackNum int) bool {
	if node == nil {
		return blackNum == 0
	}

	if !IsRed(node) {
		blackNum = blackNum - 1
	}

	if !node.Left.IsBalanced(blackNum) {
		fmt.Println("node.Left to leaf black link is not ", blackNum)
		return false
	}

	if !node.Right.IsBalanced(blackNum) {
		fmt.Println("node.Right to leaf black link is not ", blackNum)
		return false
	}

	return true
}

func main() {
	tree := NewRBTree()
	values := []int64{2, 3, 7, 10, 10, 10, 10, 23, 9, 102, 109, 111, 112, 113}
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

	// 删除存在的9后，再查找9
	tree.Delete(9)
	tree.Delete(10)
	tree.Delete(2)
	tree.Delete(3)
	tree.Add(4)
	tree.Add(3)
	tree.Add(10)
	tree.Delete(111)
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}

	if tree.IsRBTree() {
		fmt.Println("is a rb tree")
	} else {
		fmt.Println("is not rb tree")
	}

	tree.Delete(3)
	tree.Delete(4)
	tree.Delete(7)
	tree.Delete(10)
	tree.Delete(23)
	tree.Delete(102)
	tree.Delete(109)
	tree.Delete(112)
	tree.Delete(112)
	tree.MidOrder()
}
