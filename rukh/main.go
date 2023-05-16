package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func newNode(val int) *Node {
	return &Node{val: val, left: nil, right: nil}
}

func Insert(node *Node, val int) {
	if val < node.val {
		if node.left == nil {
			node.left = newNode(val)
			return
		}

		Insert(node.left, val)
	} else {
		if node.right == nil {
			node.right = newNode(val)
			return
		}

		Insert(node.right, val)
	}
}

func inOrderTraversal(node *Node) {

	if node == nil {
		return
	}

	inOrderTraversal(node.left)

	fmt.Printf("%d \n", node.val)

	inOrderTraversal(node.right)
}

func preOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d \n", node.val)

	preOrderTraversal(node.left)

	preOrderTraversal(node.right)
}

func postOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	postOrderTraversal(node.left)

	postOrderTraversal(node.right)

	fmt.Printf("%d \n", node.val)
}

func main() {
	rootNode := newNode(5)

	Insert(rootNode, 2)

	Insert(rootNode, 6)

	Insert(rootNode, 4)

	Insert(rootNode, 10)

	fmt.Println("Inorder Traversal")
	inOrderTraversal(rootNode)

	fmt.Println("preorder Traversal")
	preOrderTraversal(rootNode)

	fmt.Println("postorder Traversal")
	postOrderTraversal(rootNode)
}
