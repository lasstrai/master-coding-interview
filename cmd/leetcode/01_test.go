package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	name string
	nums []int
	target int
	expected []int	
}

func TestTwoSums(t *testing.T) {
	tests := []TestCase{
		{"nums = [2, 7, 11, 15] | target = 9 | [0, 1]", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"nums = [3, 2, 4] | target = 6 | [1, 2]", []int{3, 2, 4}, 6, []int{1, 2}},
		{"nums = [3, 3] | target = 6 | [0, 1]", []int{3, 3}, 6, []int{0, 1}},
	}
	
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			fmt.Println(testCase.name)
			output := twoSums(testCase.nums, testCase.target)
			if !reflect.DeepEqual(testCase.expected, output) {
				t.Errorf("expected: %v output: %v", testCase.expected, output)
			} else {
				fmt.Printf(":) test passed\n")
			}
		})
	}
}

// Input: nums = [3,3], target = 6
// Output: [0,1]