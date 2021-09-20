package main

import "fmt"

// Node will be used to store a vertex of a tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// inOrder traversal
func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Printf(" %d,", root.Value)
	inOrder(root.Right)
}

// preOrder traversal
func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf(" %d,", root.Value)
	preOrder(root.Left)
	preOrder(root.Right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Printf(" %d,", root.Value)
}

/*
	Example Tree we use in code
            12
		  /    \
        46     75
	   /  \      \
	  18   5       39
            \
            20
*/
func main() {
	root := &Node{
		Value: 12,
		Left: &Node{
			Value: 46,
			Left:  &Node{Value: 18},
			Right: &Node{
				Value: 5,
				Right: &Node{Value: 20},
			},
		},
		Right: &Node{
			Value: 75,
			Right: &Node{Value: 39},
		},
	}
	fmt.Printf("In Order: ")
	inOrder(root)
	fmt.Printf("\nPre-Order: ")
	preOrder(root)
	fmt.Printf("\nPost-Order: ")
	postOrder(root)
	fmt.Println()
}
