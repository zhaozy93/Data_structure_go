package btree

// import (
// 	"fmt"
// 	"sync"
// )

// type BTree struct {
// 	step        int
// 	root        *Node
// 	compareFunc func(interface{}, interface{}) int
// 	locker      sync.Mutex
// }

// func NewBTree(step int, compareFunc func(interface{}, interface{}) int) *BTree {
// 	bt := &BTree{
// 		step:        step,
// 		compareFunc: compareFunc,
// 	}
// 	return bt
// }
// func (bt *BTree) SetRoot(n *Node) {
// 	bt.locker.Lock()
// 	defer bt.locker.Unlock()
// 	bt.root = n
// }
// func (bt *BTree) Search(val interface{}) bool {
// 	bt.locker.Lock()
// 	defer bt.locker.Unlock()
// 	return bt.root.Search(val)
// }
// func (bt *BTree) Insert(val interface{}) bool {
// 	bt.locker.Lock()
// 	defer bt.locker.Unlock()
// 	node := bt.root
// 	for {
// 		// 找元素本身有没有
// 		for i := 0; i < len(node.value); i++ {
// 			if bt.compareFunc(val, node.value[i]) == 0 {
// 				return false
// 			}
// 		}
// 		// 元素没有孩子 有没有找到就证明了没有找到
// 		if node.child == nil || len(node.child) == 0 {
// 			node.AddValue(val)
// 			return true
// 		}
// 		i := 0
// 		for ; i < len(node.child); i++ {
// 			before := bt.compareFunc(val, node.child[i].value[0])
// 			// node
// 			if before == 0 {
// 				return false
// 			}
// 			if before == -1 {
// 				break
// 			}
// 			after := bt.compareFunc(val, node.child[i].value[len(node.child[i].value)-1])
// 			if after == 1 {
// 				continue
// 			}
// 			if after == 0 {
// 				return false
// 			}
// 			if after == -1 {
// 				break
// 			}
// 		}
// 		if i == len(node.child) {
// 			i--
// 		}
// 		node = node.child[i]
// 	}
// 	return false
// }
// func (bt *BTree) Delete(val interface{}) bool {
// 	bt.locker.Lock()
// 	defer bt.locker.Unlock()
// 	node := bt.root
// 	for {
// 		// 找元素本身有没有
// 		for i := 0; i < len(node.value); i++ {
// 			if bt.compareFunc(val, node.value[i]) == 0 {
// 				break
// 			}
// 		}

// 		// 元素没有孩子 有没有找到就证明了没有找到
// 		if node.child == nil || len(node.child) == 0 {
// 			node.AddValue(val)
// 			return true
// 		}
// 		i := 0
// 		for ; i < len(node.child); i++ {
// 			before := bt.compareFunc(val, node.child[i].value[0])
// 			// node
// 			if before == 0 {
// 				return false
// 			}
// 			if before == -1 {
// 				break
// 			}
// 			after := bt.compareFunc(val, node.child[i].value[len(node.child[i].value)-1])
// 			if after == 1 {
// 				continue
// 			}
// 			if after == 0 {
// 				return false
// 			}
// 			if after == -1 {
// 				break
// 			}
// 		}
// 		if i == len(node.child) {
// 			i--
// 		}
// 		node = node.child[i]
// 	}
// 	return false
// }

// type Node struct {
// 	bt     *BTree
// 	parent *Node
// 	child  []*Node
// 	value  []interface{}
// }

// func NewNode(bt *BTree) *Node {
// 	value := make([]interface{}, 0)
// 	child := make([]*Node, 0)
// 	return &Node{bt: bt, child: child, value: value}
// }
// func (n *Node) SetParent(p *Node) {
// 	n.parent = p
// }
// func (n *Node) GetParent() *Node {
// 	return n.parent
// }
// func (n *Node) AddChild(c *Node) {
// 	min := c.value[0]
// 	i := 0
// 	for i < len(n.child) {
// 		if n.bt.compareFunc(min, n.child[i].value[0]) == -1 {
// 			break
// 		}
// 		i++
// 	}
// 	oldChild := n.child
// 	n.child = make([]*Node, len(oldChild)+1)
// 	copy(n.child[:i], oldChild)
// 	n.child[i] = c
// 	copy(n.child[i+1:], oldChild[i:])
// }
// func (n *Node) AddValue(val interface{}) {
// 	//找到value所属的位置
// 	i := 0
// 	for i < len(n.value) {
// 		if n.bt.compareFunc(val, n.value[i]) == 1 {
// 			i++
// 			continue
// 		}
// 		break
// 	}
// 	oldVal := n.value
// 	n.value = make([]interface{}, 0)
// 	n.value = append(n.value, oldVal[:i]...)
// 	n.value = append(n.value, val)
// 	n.value = append(n.value, oldVal[i:]...)
// 	if len(n.value) < n.bt.step {
// 		return
// 	}
// 	// 保存的数据大于步长 则需要分裂
// 	// 首先切割val到自身和peerNode(值一定大于自身)
// 	splitIndex := n.bt.step / 2
// 	oldVal = n.value
// 	n.value = make([]interface{}, splitIndex)
// 	copy(n.value, oldVal)
// 	peer := NewNode(n.bt)
// 	peer.value = make([]interface{}, n.bt.step-splitIndex-1)
// 	copy(peer.value, oldVal[splitIndex+1:])
// 	// val是要之后添加到parent的
// 	val = oldVal[splitIndex]
// 	// 继续切割之前的child到自身和peer
// 	if n.child != nil && len(n.child) != 0 {
// 		index := 0
// 		for {
// 			if index == len(n.child) {
// 				break
// 			}
// 			if n.bt.compareFunc(val, n.child[index].value[len(n.child[index].value)-1]) == 1 {
// 				index++
// 				continue
// 			}
// 			break
// 		}
// 		oldChild := n.child
// 		n.child = make([]*Node, index)
// 		copy(n.child, oldChild[:index])

// 		peer.child = make([]*Node, len(oldChild)-index)
// 		copy(peer.child, oldChild[index:])
// 	}
// 	if n.parent == nil {
// 		p := NewNode(n.bt)
// 		p.AddChild(n)
// 		p.AddChild(peer)
// 		p.AddValue(val)
// 		n.SetParent(p)
// 		peer.SetParent(p)
// 		n.bt.root = p
// 		return
// 	} else {
// 		n.parent.AddChild(peer)
// 		n.parent.AddValue(val)
// 	}
// }
// func (n *Node) Search(val interface{}) bool {
// 	// 找元素本身有没有
// 	for i := 0; i < len(n.value); i++ {
// 		if n.bt.compareFunc(val, n.value[i]) == 0 {
// 			return true
// 		}
// 	}
// 	// 元素没有孩子 有没有找到就证明了没有找到
// 	if n.child == nil || len(n.child) == 0 {
// 		return false
// 	}
// 	for i := 0; i < len(n.child); i++ {
// 		before := n.bt.compareFunc(val, n.child[i].value[0])
// 		if before == 0 {
// 			return true
// 		}
// 		if before == -1 {
// 			return n.child[i].Search(val)
// 		}
// 		after := n.bt.compareFunc(val, n.child[i].value[len(n.child[i].value)-1])
// 		if after == -1 {
// 			return n.child[i].Search(val)
// 		}
// 		if after == 0 {
// 			return true
// 		}
// 	}
// 	return n.child[len(n.child)-1].Search(val)
// }
// func (n *Node) Print(tag int, ftag int) {
// 	fmt.Print("FTag: ", ftag, ". MyTag: ", tag, ". Value: ", n.value, "\n")
// 	if n.child != nil {
// 		for i := 0; i < len(n.child); i++ {
// 			n.child[i].Print(tag+i+1, tag)
// 		}
// 	}
// }
