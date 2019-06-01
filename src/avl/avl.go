package avl

type Node struct {
	val   int
	bf    int
	left  *Node
	right *Node
}

func R_rotate(p *Node) *Node {
	l := p.left
	p.left = l.right
	l.right = p
	return l
}

func L_rotate(p *Node) *Node {
	r := p.right
	p.right = r.left
	r.left = p
	return r
}

type AVL struct {
	root *Node
}

func (avl *AVL) Insert(e int) bool {
	if avl.root == nil {
		avl.root = &Node{e, 0, nil, nil}
		return
	}
	_, ok := avl.insert(avl.root, e)
	return ok
}

// (bool,bool)  (taller,insert)
func (avl *AVL) insert(root *Node, e int) (bool, bool) {
	if e == root.val {
		return false, false
	} else if e < root.val {
		if root.left != nil {
			taller, ok := avl.insert(root.left, e)
			if !ok {
				return false, false
			}
			if taller {
				if root.bf == -1 {
					root.bf = 0
					return false, true
				} else if root.bf == 0 {
					root.bf = 1
					return true, true
				} else {
					LeftBance(root)
					return false, true
				}
			}
		} else {
			root.left = &Node{e, 0, nil, nil}
			if root.right == nil {
				root.bf = 1
				return true, true
			} else {
				root.bf = 0
				return false, true
			}
		}
	} else {
		if root.right != nil {
			taller, ok := avl.insert(root.right, e)
			if !ok {
				return false, false
			}
			if taller {
				if root.bf == 0 {
					root.bf = -1
					return true, true
				} else if root.bf == 1 {
					root.bf = 0
					return false, true
				} else {
					RightBalance(root)
					return false, true
				}
			}
		} else {
			root.right = &Node{e, 0, nil, nil}
			if root.left == nil {
				root.bf = -1
				return true, true
			} else {
				root.bf = 0
				return false, true
			}
		}

	}
}

func LeftBalance(root *Node) {
	rootptr := &root
	l := root.left
	if l.bf == 1 {
		root.bf = 0
		l.bf = 0
		*rootptr = R_rotate(root)
	} else if l.bf == -1 {
		if l.right.bf == 1 {
			root.bf = -1
			l.bf = 0
		} else if l.right.bf == 0 {
			root.bf = 0
			l.bf = 0
		} else {
			root.bf = 0
			l.bf = 1
		}
		l.bf = 0
		root.left = L_rotate(root.left)
		*rootptr = R_rotate(root)
	}
}

func RightBalance(root *Node) {
	rootptr := &root
	r := root.right
	if r.bf == -1 {
		root.bf = 0
		r.bf = 0
		*rootptr = L_rotate(root)
	} else if r.bf == 1 {
		if r.left.bf == -1 {
			root.bf = 1
			r.bf = 0
		} else if r.right.bf == 0 {
			root.bf = 0
			r.bf = 0
		} else {
			root.bf = 0
			r.bf = -1
		}
		root.right = R_rotate(root.right)
		*rootptr = L_rotate(root)
	}
}
