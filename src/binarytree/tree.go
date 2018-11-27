package tree

type Node struct {
	parent *Node
	lchild *Node
	rchild *Node
	value  interface{}
}

func NewBinaryNode(val interface{}) *Node {
	return &Node{value: val}
}
func (n *Node) SetValue(val interface{}) {
	n.value = val
}
func (n *Node) GetValue() interface{} {
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
}

func NewBinaryTree(r *Node) *Tree {
	return &Tree{r}
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
