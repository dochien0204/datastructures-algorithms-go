package bst

import "fmt"

type Tree struct {
	Root *Node
}

func NewTree(root *Node) *Tree {
	return &Tree{
		Root: root,
	}
}

type Interface interface {
	Less(than Interface) bool
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func NewNode(key int) *Node {
	return &Node{
		Key: key,
	}
}

// insert element in BST
func (t *Tree) Insert(data int) *Tree {
	if t.Root == nil {
		t.Root = &Node{
			Key: data,
		}
	} else {
		t.Root.insert(data)
	}

	return t
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data < n.Key {
		if n.Left == nil {
			n.Left = &Node{
				Key: data,
			}
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{
				Key: data,
			}
		} else {
			n.Right.insert(data)
		}
	}
}

func (t *Tree) InorderTraversal(root *Node) {
	if root == nil {
		return
	}
	t.InorderTraversal(root.Left)
	fmt.Print(root.Key, " ")
	t.InorderTraversal(root.Right)
}

func (t *Tree) PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Key, " ")
	t.PreOrderTraversal(root.Left)
	t.PreOrderTraversal(root.Right)
}

func (t *Tree) PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	t.PostOrderTraversal(root.Left)
	t.PostOrderTraversal(root.Right)
	fmt.Print(root.Key, " ")
}
