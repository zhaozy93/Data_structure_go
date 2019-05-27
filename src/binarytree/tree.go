package tree

// 不支持并发
type Value interface {
	Less(Value) int
}
type Node struct {
	parent *Node
	lchild *Node
	rchild *Node
	value  Value
}

func NewBinaryNode(val Value) *Node {
	return &Node{value: val}
}
func (n *Node) SetValue(val Value) {
	n.value = val
}
func (n *Node) GetValue() Value {
	return n.value
}
func (n *Node) SetParent(p *Node) {
	n.parent = p
}
func (n *Node) GetParent() *Node {
	return n.parent
}
func (n *Node) SetLeftChild(c *Node) {
	n.lchild = c
}
func (n *Node) GetLeftChild() *Node {
	return n.lchild
}
func (n *Node) SetRightChild(c *Node) {
	n.rchild = c
}
func (n *Node) GetRightChild() *Node {
	return n.rchild
}

func (n *Node) preTraversal(res *[]interface{}) {
	*res = append(*res, n.GetValue())
	if n.lchild != nil {
		n.lchild.preTraversal(res)
	}
	if n.rchild != nil {
		n.rchild.preTraversal(res)
	}
}
func (n *Node) inTraversal(res *[]interface{}) {
	if n.lchild != nil {
		n.lchild.inTraversal(res)
	}
	*res = append(*res, n.GetValue())
	if n.rchild != nil {
		n.rchild.inTraversal(res)
	}
}
func (n *Node) postTraversal(res *[]interface{}) {
	if n.lchild != nil {
		n.lchild.postTraversal(res)
	}
	if n.rchild != nil {
		n.rchild.postTraversal(res)
	}
	*res = append(*res, n.GetValue())
}

// 每一个节点其实都是一个root二叉树
// 用Tree对node进行一次封装而已
type Tree struct {
	root *Node
	len  int64
}

func NewBinaryTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(val Value) bool {
	if t.root == nil {
		r := NewBinaryNode(val)
		t.root = r
		t.len++
		return true
	}
	preNode := t.root
	node := preNode
	less := 0
	for {
		if node == nil {
			break
		}
		preNode = node
		less = node.GetValue().Less(val)
		if less == 0 {
			return false
		}
		if less < 0 {
			node = node.GetRightChild()
		} else {
			node = node.GetLeftChild()
		}
		continue
	}
	if less < 0 {
		preNode.SetRightChild(NewBinaryNode(val))
		t.len++
	} else {
		preNode.SetLeftChild(NewBinaryNode(val))
		t.len++
	}
	return false
}

func (t *Tree) Delete(val Value) bool {
	node := t.root
	less := 0
	for {
		if node == nil {
			break
		}
		less = node.GetValue().Less(val)
		if less == 0 {
			break
		}
		if less < 0 {
			node = node.GetRightChild()
		}
		node = node.GetLeftChild()
		continue
	}
	if node != nil {
		t.len--
		parent := node.GetParent()
		left := node.GetLeftChild()
		right := node.GetRightChild()
		if left == nil && right == nil {
			if parent.GetLeftChild() == node {
				parent.SetLeftChild(nil)
			} else {
				parent.SetRightChild(nil)
			}
			node.SetParent(nil)
			return true
		}
		if left == nil {
			node.GetRightChild().SetParent(parent)
			if parent.GetLeftChild() == node {
				parent.SetLeftChild(node.GetRightChild())
			} else {
				parent.SetRightChild(node.GetRightChild())
			}
			node.SetParent(nil)
			node.SetRightChild(nil)
			node.SetRightChild(nil)
			return true
		}
		if right == nil {
			node.GetLeftChild().SetParent(parent)
			if parent.GetLeftChild() == node {
				parent.SetLeftChild(node.GetLeftChild())
			} else {
				parent.SetRightChild(node.GetLeftChild())
			}
			node.SetParent(nil)
			node.SetRightChild(nil)
			node.SetRightChild(nil)
			return true
		}
		// 寻找左子树 最大的节点， 然后顶替当前节点的位置
		for {
			right := left.GetRightChild()
			if right == nil {
				break
			}
			left = right
		}
		left = left.GetRightChild()
		if left.GetLeftChild() != nil {
			left.GetLeftChild().SetParent(left.GetParent())
			left.GetParent().SetRightChild(left.GetLeftChild())
		}
		if parent.GetLeftChild() == node {
			parent.SetLeftChild(left)
		} else {
			parent.SetRightChild(left)
		}
		left.SetParent(parent)
		left.SetLeftChild(node.GetLeftChild())
		left.SetRightChild(node.GetRightChild())
		node.SetParent(nil)
		node.SetRightChild(nil)
		node.SetRightChild(nil)
		return true
	}
	return false
}

func (t *Tree) Find(val Value) *Node {
	node := t.root
	less := 0
	for {
		if node == nil {
			break
		}
		less = node.GetValue().Less(val)
		if less == 0 {
			return node
		}
		if less < 0 {
			node = node.GetRightChild()
		}
		node = node.GetLeftChild()
		continue
	}
	return nil
}

func (t *Tree) PreTraversal() []interface{} {
	if t.root != nil {
		result := make([]interface{}, 0)
		t.root.preTraversal(&result)
		return result
	}
	return nil
}
func (t *Tree) InTraversal() []interface{} {
	if t.root != nil {
		result := make([]interface{}, 0)
		t.root.inTraversal(&result)
		return result
	}
	return nil
}
func (t *Tree) PostTraversal() []interface{} {
	if t.root != nil {
		result := make([]interface{}, 0)
		t.root.postTraversal(&result)
		return result
	}
	return nil
}
