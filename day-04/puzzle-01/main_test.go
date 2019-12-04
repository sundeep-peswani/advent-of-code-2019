package main

import "testing"

func Test_IsValid(t *testing.T) {
	tests := []struct{
		input int
		expected bool
	}{
		{111111, true},
		{11111, false},
		{223450, false},
		{123789, false},
		{123788, true},
	}

	for i, test := range tests {
		actual := isValid(test.input)
		if actual != test.expected {
			t.Errorf("Test %d: expected %v for %d, got %v\n", i+1, test.expected, test.input, actual)
		}
	}
}