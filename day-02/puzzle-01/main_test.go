package main

import "testing"

func Test_Intcode(t *testing.T) {
	tests := []struct{
		opcodes string
		expected []int
	}{
		{"1,9,10,3,2,3,11,0,99,30,40,50", []int{3500,9,10,70,2,3,11,0,99,30,40,50}},
		{"1,0,0,0,99", []int{2,0,0,0,99}},
		{"2,3,0,3,99", []int{2,3,0,6,99}},
		{"2,4,4,5,99,0", []int{2,4,4,5,99,9801}},
		{"1,1,1,4,99,5,6,0,99", []int{30,1,1,4,2,5,6,0,99}},
	}

	for i, test := range tests {
		var testGap intcode
		testGap.read(test.opcodes)
		testGap.run()

		if !equal(test.expected, testGap.opcodes) {
			t.Errorf("Test %d: expected %v, actual %v\n", i + 1, test.expected, testGap.opcodes)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, j := range a {
		if j != b[i] {
			return false
		}
	}

	return true
}
