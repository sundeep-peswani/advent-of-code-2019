package main

import "testing"

func Test_IsValid(t *testing.T) {
	tests := []struct{
		input int
		expected bool
	}{
		{111111, false},
		{11111, false},
		{223450, false},
		{123789, false},
		{123788, true},		
		{112233, true},
		{123444, false},
		{111122, true},
		{788999, true},
		{22333, false},
		{223333, true},
		{222333, false},
	}

	for i, test := range tests {
		actual := isValid(test.input)
		if actual != test.expected {
			t.Errorf("Test %d: expected %v for %d, got %v\n", i+1, test.expected, test.input, actual)
		}
	}
}