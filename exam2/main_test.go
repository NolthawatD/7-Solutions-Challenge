package main

import (
	"fmt"
	"testing"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		encoded  string
		expected string
	}{
		{
			encoded:  "LLRR=",
			expected: "210122",
		},
		{
			encoded:  "==RLL",
			expected: "000210",
		},
		{
			encoded:  "=LLRR",
			expected: "221012",
		},
		{
			encoded:  "RRL=R",
			expected: "012001",
		},
	}

	for _, test := range tests {
		resultString, _ := decodeString(test.encoded)
		fmt.Printf("input = %s output = %s\n", test.encoded, resultString)
		if resultString != test.expected {
			t.Errorf("For encoded %s, expected %s but got %s", test.encoded, test.expected, resultString)
		}
	}
}
