package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name string
	input []int
	output []int
}

func TestQuickSort(t *testing.T) {
	tests := []TestCase{
		{"empty array", []int{}, []int{}},
		{"single element", []int{5}, []int{5}},
		{"sorted slice", []int{1, 2, 3}, []int{1, 2, 3}},
		{"reversed slice", []int{3, 2, 1}, []int{1, 2, 3}},
		{"unsorted slice", []int{5, 2, 9, 1, 5, 6}, []int{1, 2, 5, 5, 6, 9}},
		{"duplicates", []int{4, 2, 2, 8, 3, 3, 1}, []int{1, 2, 2, 3, 3, 4, 8}},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			quicksort(&testCase.input, 0, len(testCase.input))
			if !reflect.DeepEqual(testCase.input, testCase.output) {
				t.Errorf("what I got %v, what I want %v", testCase.input, testCase.output)
			}
		})
	}
}