package main

import (
	"fmt"
	"testing"
)

func TestSumHirarchyMaxValuePath(t *testing.T) {
	tests := []struct {
		hierarchy [][]int
		expected  int
	}{
		{
			hierarchy: hirarchy1,
			expected:  237,
		},
		{
			hierarchy: hirarchy2,
			expected:  7273,
		},
	}

	for _, test := range tests {
		result := sumHirarchyMaxValuePath(test.hierarchy)
		fmt.Printf("input = %v output = %d\n", test.hierarchy, result)
		if result != test.expected {
			t.Errorf("For hierarchy %v, expected %d but got %d", test.hierarchy, test.expected, result)
		}
	}
}
