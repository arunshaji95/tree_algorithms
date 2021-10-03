package main

import "fmt"

// Program to display the depth of a binary tree
// Node stores each vertex of the tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Algoriithm to calculate the height of a binary tree
func Height(root *Node) int {
	if root == nil {
		return 0
	} else {
		heightLeft := Height(root.Left)
		heightRight := Height(root.Right)
		if heightLeft > heightRight {
			return 1 + heightLeft
		} else {
			return 1 + heightRight
		}
	}
}

func main() {
	/*
		    Example:
			          90
			        /   \
			      15     201
			     /  \
			   8     -1
			    \
			     56
	*/
	tree := &Node{
		Value: 90,
		Left: &Node{
			Value: 15,
			Left: &Node{
				Value: 8,
				Right: &Node{
					Value: 56,
				},
			},
			Right: &Node{
				Value: -1,
			},
		},
		Right: &Node{
			Value: 201,
		},
	}
	fmt.Printf("Height of the tree is %d\n", Height(tree))
}
