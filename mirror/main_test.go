package main

import "testing"

func Test_Mirror(t *testing.T) {
	testCases := []struct {
		name     string
		t1       *Node
		t2       *Node
		expected bool
	}{
		{
			name: "both trees are mirror images",
			t1: &Node{
				Value: 15,
				Left:  &Node{Value: 85},
				Right: &Node{Value: 1},
			},
			t2: &Node{
				Value: 15,
				Left:  &Node{Value: 1},
				Right: &Node{Value: 85},
			},
			expected: true,
		}, {
			name: "both trees are not mirror images and one tree is having more depth",
			t1: &Node{
				Value: 8,
				Left: &Node{
					Value: 26,
					Right: &Node{Value: 46},
				},
				Right: &Node{
					Value: 15,
				},
			},
			t2: &Node{
				Value: 8,
				Left: &Node{
					Value: 15,
				},
				Right: &Node{
					Value: 26,
				},
			},
			expected: false,
		},
	}

	for _, tt := range testCases {
		got := Mirror(tt.t1, tt.t2)
		if got != tt.expected {
			t.Errorf("%s: expected: %v, got %v", tt.name, tt.expected, got)
		}
	}
}
