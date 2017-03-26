package demo_test

import (
	"testing"

	sample "github.com/klrkdekira/monkey/sample"
)

func TestAdd(t *testing.T) {
	tests := [][]int{
		{1, 1, 2},
		{2, 2, 4},
	}

	for _, test := range tests {
		result := sample.Add(test[0], test[1])
		if result != test[2] {
			t.Fatalf("Expected %d, got %d", test[2], result)
		}
	}
}
