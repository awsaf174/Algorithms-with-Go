package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if k > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if k < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}

}

func (n *Node) Search(k int) bool {

	if n == nil {
		return false
	}

	if k > n.Key {
		return n.Right.Search(k)
	} else if k < n.Key {
		return n.Left.Search(k)
	}

	return true
}

func main() {

	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(10)
	tree.Insert(150)

	fmt.Println(tree.Search(51))
}
