package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase345 struct {
	input string
	expected string
}

func TestReverseVowels(t *testing.T) {
	tests := []TestCase345 {
		{"IceCreAm", "AceCreIm"},
		{"leetcode", "leotcede"},
	}

	for _, testCase := range tests {
		t.Run(testCase.input, func(t *testing.T) {
			fmt.Println(testCase.input)
			output := reverseVowels(testCase.input)
			if !reflect.DeepEqual(output, testCase.expected) {
				t.Errorf("expected: %s output: %s", testCase.expected, output)
			} else {
				fmt.Printf(":) %s test passed correctly\n", testCase.input)
			}
		})
	}
}
