// Check whether the given binary trees are identical
package main

import "fmt"

// Node points to each node in our tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Identical checks whether two binary trees are equal
func Identical(m, n *Node) bool {
	if m == nil && n == nil {
		return true
	}

	if m != nil && n != nil {
		if m.Value == n.Value && Identical(m.Left, n.Left) && Identical(m.Right, n.Right) {
			return true
		}
		return false
	}
	return false
}

/*
Example: 1; Identical

   10              10
   /\      <=>     /\
 15   20         15   20

------------------------------------
   Example: 2; Non-identical

         18                   90
        /  \                 / \
      49    29     !=       19  21
     / \      \            /
   10   42     21         76
-------------------------------------

*/

func main() {
	n1 := &Node{
		Value: 10,
		Left:  &Node{Value: 15},
		Right: &Node{Value: 20},
	}
	n2 := &Node{
		Value: 10,
		Left:  &Node{Value: 15},
		Right: &Node{Value: 20},
	}

	if Identical(n1, n2) {
		fmt.Println("Both trees are identical")
	} else {
		fmt.Println("Both trees are not identical")
	}

	n3 := &Node{
		Value: 18,
		Left: &Node{
			Value: 49,
			Left:  &Node{Value: 10},
			Right: &Node{Value: 42},
		},
		Right: &Node{
			Value: 29,
			Right: &Node{Value: 21},
		},
	}
	n4 := &Node{
		Value: 90,
		Left: &Node{
			Value: 19,
			Left:  &Node{Value: 76},
			Right: &Node{Value: 21},
		},
	}
	if Identical(n3, n4) {
		fmt.Println("Both trees are identical")
	} else {
		fmt.Println("Both trees are not identical")
	}
}
