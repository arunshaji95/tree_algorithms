package main

import "fmt"

// Node will be used to store a vertex of a tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Mirror(x, y *Node) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	if x.Value == y.Value {
		if Mirror(x.Left, y.Right) && Mirror(x.Right, y.Left) {
			return true
		}
	}
	return false
}

func main() {

	/* Input 1
			90            90
	       /  \          /  \
		12    35        35   12
		/                     \
	   4                       4

	*/
	t1 := &Node{
		Value: 90,
		Left: &Node{
			Value: 12,
			Left:  &Node{Value: 4},
		},
		Right: &Node{
			Value: 35,
		},
	}

	t2 := &Node{
		Value: 90,
		Left:  &Node{Value: 35},
		Right: &Node{
			Value: 12,
			Right: &Node{Value: 4},
		},
	}
	if Mirror(t1, t2) {
		fmt.Println("t1 & t2 are mirror images")
	}
}
