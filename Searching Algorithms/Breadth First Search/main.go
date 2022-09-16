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

func (n *Node) BFSTraversal() []int {
	var currentNode = n
	var list = []int{}
	var queue = []*Node{}
	queue = append(queue, currentNode)

	for len(queue) > 0 {
		currentNode = queue[0]
		queue = append(queue[1:])
		list = append(list, currentNode.Key)
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
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

	fmt.Println(tree.BFSTraversal())

}
