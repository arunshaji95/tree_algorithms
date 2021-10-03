package main

import "testing"

func Test_Height(t *testing.T) {
	testCases := []struct {
		Name     string
		Root     *Node
		Expected int
	}{
		{
			Name:     "When there are no nodes",
			Root:     nil,
			Expected: 0,
		}, {
			Name: "When there are more nodes in left subtree",
			Root: &Node{
				Value: 15,
				Left: &Node{
					Value: 28,
					Left: &Node{
						Value: 89,
					},
				},
				Right: &Node{
					Value: 23,
				},
			},
			Expected: 3,
		}, {
			Name: "When there are more nodes in right subtree",
			Root: &Node{
				Value: 15,
				Right: &Node{
					Value: 28,
					Left: &Node{
						Value: 89,
					},
				},
				Left: &Node{
					Value: 23,
				},
			},
			Expected: 3,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			got := Height(tt.Root)
			if got != tt.Expected {
				t.Errorf("Expected: %v for height, got: %v", tt.Expected, got)
			}
		})
	}
}
