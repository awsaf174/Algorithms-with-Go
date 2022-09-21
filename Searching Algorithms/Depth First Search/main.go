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

func (n *Node) DFSInOrder() (list []int) {
	return n.TraverseInOrder(list)
}

func (n *Node) DFSPreOrder() (list []int) {
	return n.TraversePreOrder(list)
}

func (n *Node) DFSPostOrder() (list []int) {
	return n.TraversePostOrder(list)
}

func (n *Node) TraversePreOrder(list []int) []int {

	list = append(list, n.Key)

	if n.Left != nil {
		list = n.Left.TraversePreOrder(list)
	}

	if n.Right != nil {
		list = n.Right.TraversePreOrder(list)
	}

	return list

}

func (n *Node) TraversePostOrder(list []int) []int {
	if n.Left != nil {
		list = n.Left.TraversePostOrder(list)
	}

	if n.Right != nil {
		list = n.Right.TraversePostOrder(list)
	}

	list = append(list, n.Key)

	return list
}

func (n *Node) TraverseInOrder(list []int) []int {
	if n.Left != nil {
		list = n.Left.TraverseInOrder(list)
	}

	list = append(list, n.Key)

	if n.Right != nil {
		list = n.Right.TraverseInOrder(list)
	}

	return list
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(150)
	tree.Insert(110)
	tree.Insert(160)
	tree.Insert(50)
	tree.Insert(10)
	tree.Insert(60)

	fmt.Println("In Order: ", tree.DFSInOrder())
	fmt.Println("Pre Order: ", tree.DFSPreOrder())
	fmt.Println("Post Order: ", tree.DFSPostOrder())
}
