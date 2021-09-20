package main

import "testing"

func TestIdentical(t *testing.T) {
	testCases := []struct {
		name     string
		t1       *Node
		t2       *Node
		expected bool
	}{
		{
			name:     "Case 1: both trees are null",
			expected: true,
		}, {
			name:     "Case 2: both trees are equal",
			t1:       &Node{Value: 12},
			t2:       &Node{Value: 12},
			expected: true,
		}, {
			name: "Case 3: both trees are not equal",
			t1: &Node{
				Value: 15,
				Left:  &Node{Value: 21},
				Right: &Node{Value: 44},
			},
			t2: &Node{
				Value: 15,
				Left:  &Node{Value: 21},
				Right: &Node{
					Value: 44,
					Left:  &Node{Value: 18},
				},
			},
			expected: false,
		},
	}

	for _, tt := range testCases {
		got := Identical(tt.t1, tt.t2)
		if got != tt.expected {
			t.Errorf("%s; got %v, expected %v", tt.name, got, tt.expected)
		}
	}
}
