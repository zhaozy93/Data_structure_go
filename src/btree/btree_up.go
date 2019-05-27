package btree

import (
	"fmt"
	"sync"
)

type BTree struct {
	step        int
	root        *Node
	len         int64
	compareFunc func(interface{}, interface{}) int
	locker      sync.Mutex
}

func NewBTree(step int, compareFunc func(interface{}, interface{}) int) *BTree {
	bt := &BTree{
		step:        step,
		compareFunc: compareFunc,
	}
	return bt
}
func (bt *BTree) SetRoot(n *Node) {
	bt.locker.Lock()
	defer bt.locker.Unlock()
	bt.root = n
}
func (bt *BTree) Search(val interface{}) bool {
	bt.locker.Lock()
	defer bt.locker.Unlock()
	return bt.root.Search(val)
}
func (bt *BTree) Insert(val interface{}) bool {
	bt.locker.Lock()
	defer bt.locker.Unlock()
	if bt.root == nil {
		root := NewNode(bt)
		root.AddValue(val)
		bt.root = root
		bt.len++
		return true
	}
	node := bt.root
	for {
		// 找元素本身有没有
		for i := 0; i < len(node.value); i++ {
			if bt.compareFunc(val, node.value[i]) == 0 {
				return false
			}
		}
		// 元素没有孩子 有没有找到就证明了没有找到
		if node.child == nil || len(node.child) == 0 {
			node.AddValue(val)
			return true
		}
		i := 0
		for ; i < len(node.child); i++ {
			before := bt.compareFunc(val, node.child[i].value[0])
			// node
			if before == 0 {
				return false
			}
			if before == -1 {
				break
			}
			after := bt.compareFunc(val, node.child[i].value[len(node.child[i].value)-1])
			if after == 1 {
				continue
			}
			if after == 0 {
				return false
			}
			if after == -1 {
				break
			}
		}
		if i == len(node.child) {
			i--
		}
		node = node.child[i]
	}
	return false
}
func (bt *BTree) Delete(val interface{}) bool {
	bt.locker.Lock()
	defer bt.locker.Unlock()
	node := bt.root
	for {
		// 找元素本身有没有
		for i := 0; i < len(node.value); i++ {
			if bt.compareFunc(val, node.value[i]) == 0 {
				break
			}
		}

		// 元素没有孩子 有没有找到就证明了没有找到
		if node.child == nil || len(node.child) == 0 {
			node.AddValue(val)
			return true
		}
		i := 0
		for ; i < len(node.child); i++ {
			before := bt.compareFunc(val, node.child[i].value[0])
			// node
			if before == 0 {
				return false
			}
			if before == -1 {
				break
			}
			after := bt.compareFunc(val, node.child[i].value[len(node.child[i].value)-1])
			if after == 1 {
				continue
			}
			if after == 0 {
				return false
			}
			if after == -1 {
				break
			}
		}
		if i == len(node.child) {
			i--
		}
		node = node.child[i]
	}
	return false
}

type Node struct {
	bt       *BTree
	parent   *Node
	child    []*Node
	value    []interface{}
	bitValue []byte // 1表示value 2表示child
}

func NewNode(bt *BTree) *Node {
	bitValue := make([]byte, bt.step*3)
	return &Node{bt: bt, bitValue: bitValue}
}
func (n *Node) SetParent(p *Node) {
	n.parent = p
}
func (n *Node) GetParent() *Node {
	return n.parent
}
func (n *Node) IsValue(i int) bool {
	return n.bitValue[i] == 1
}

func (n *Node) AddChild(c *Node) {
	if n.value == nil || len(n.value) == 0 {
		fmt.Print("illeage addChild")
		return
	}
	i := 0
	valIndex := 0
	childIndex := 0
	cVal := c.value[0]
	for ; i < len(n.bitValue) && n.bitValue[i] != 0; i++ {
		if n.IsValue(i) {
			com := n.bt.compareFunc(cVal, n.value[valIndex])
			if com == 1 {
				valIndex++
			} else {
				break
			}
		} else {
			val := n.child[childIndex].value[0]
			com := n.bt.compareFunc(cVal, val)
			if com == 1 {
				childIndex++
			} else {
				break
			}
		}
		continue
	}
	copy(n.bitValue[i+1:], n.bitValue[i:])
	n.bitValue[i] = 2
	if n.child == nil || len(n.child) == 0 {
		n.child = []*Node{c}
		return
	}
	n.child = append(n.child, nil)
	if childIndex == len(n.child)-1 {
		n.child[childIndex] = c
	} else {
		copy(n.child[childIndex+1:], n.child[childIndex:])
		n.child[childIndex] = c
	}
}

func (n *Node) AddValue(val interface{}) {
	if val == interface{}(6) {
		fmt.Println("xxx")
		fmt.Println(n.parent)
		fmt.Println(n.value)
		fmt.Println(n.child)
	}

	if n.value == nil {
		n.value = []interface{}{val}
		n.bitValue[0] = 1
		return
	}
	i := 0
	valIndex := 0
	childIndex := 0
	for ; i < len(n.bitValue) && n.bitValue[i] != 0; i++ {
		if n.IsValue(i) {
			com := n.bt.compareFunc(val, n.value[valIndex])
			if com == 1 {
				valIndex++
			} else if com == 0 {
				return
			} else if com == -1 {
				break
			}
		} else {
			cVal := n.child[childIndex].value[0]
			com := n.bt.compareFunc(val, cVal)
			if com == 1 {
				childIndex++
			} else {
				break
			}
		}
	}
	n.value = append(n.value, nil)
	if valIndex == len(n.value)-1 {
		n.value[valIndex] = val
	} else {
		copy(n.value[valIndex+1:], n.value[valIndex:])
		n.value[valIndex] = val
	}
	copy(n.bitValue[i+1:], n.bitValue[i:])
	n.bitValue[i] = 1
	if len(n.value) < n.bt.step {
		return
	}

	middle := n.bt.step / 2
	i = 0
	valIndex = 0
	childIndex = 0
	for ; i < len(n.bitValue) && valIndex < middle; i++ {
		if n.IsValue(i) {
			valIndex++
		} else {
			childIndex++
		}
	}
	val = n.value[valIndex-1]
	peer := NewNode(n.bt)
	peer.value = make([]interface{}, len(n.value)-valIndex)
	copy(peer.value, n.value[valIndex:])
	copy(peer.bitValue, n.bitValue[i:])
	n.value = n.value[:valIndex-1]
	bitValue := make([]byte, n.bt.step*3)
	copy(bitValue, n.bitValue[:i-1])
	n.bitValue = bitValue

	if n.child != nil && len(n.child) != 0 {
		peer.child = make([]*Node, len(n.child)-childIndex)
		copy(peer.child, n.child[childIndex:])
		n.child = n.child[:childIndex]
	}

	if val == interface{}(6) {
		fmt.Println("xxx")
		fmt.Println(n.parent)
		fmt.Println(n.value)
		fmt.Println(n.child)
	}
	if n.parent == nil {
		p := NewNode(n.bt)
		p.AddValue(val)
		p.AddChild(n)
		p.AddChild(peer)
		n.SetParent(p)
		peer.SetParent(p)
		n.bt.root = p
	} else {
		n.parent.AddChild(peer)
		n.parent.AddValue(val)
	}
	return
}

func (n *Node) Search(val interface{}) bool {
	// 找元素本身有没有
	for i := 0; i < len(n.value); i++ {
		if n.bt.compareFunc(val, n.value[i]) == 0 {
			return true
		}
	}
	// 元素没有孩子 有没有找到就证明了没有找到
	if n.child == nil || len(n.child) == 0 {
		return false
	}
	for i := 0; i < len(n.child); i++ {
		before := n.bt.compareFunc(val, n.child[i].value[0])
		if before == 0 {
			return true
		}
		if before == -1 {
			return n.child[i].Search(val)
		}
		after := n.bt.compareFunc(val, n.child[i].value[len(n.child[i].value)-1])
		if after == -1 {
			return n.child[i].Search(val)
		}
		if after == 0 {
			return true
		}
	}
	return n.child[len(n.child)-1].Search(val)
}
func (n *Node) Print(tag int, ftag int) {
	fmt.Print("FTag: ", ftag, ". MyTag: ", tag, ". Value: ", n.value, "\n")
	if n.child != nil {
		for i := 0; i < len(n.child); i++ {
			n.child[i].Print(tag+i+1, tag)
		}
	}
}
