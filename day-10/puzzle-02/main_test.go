package main

import (
	"math"
	"testing"
)

func Test_Vaporize(t *testing.T) {
	sm := newSpaceMap(".#....#####...#..\n##...##.#####..##\n##...#...#.#####.\n..#.....X...###..\n..#.#.....#....##")

	expectedOrder := []coord{
		c(8, 1), c(9, 0), c(9, 1), c(10, 0), c(9, 2), c(11, 1), c(12, 1), c(11, 2), c(15, 1),
		c(12, 2), c(13, 2), c(14, 2), c(15, 2), c(12, 3), c(16, 4), c(15, 4), c(10, 4), c(4, 4),
	}
	actualOrder := sm.vaporize(len(expectedOrder), coord{8, 3, 0.0, 0.0})

	for i, v := range expectedOrder {
		actual := actualOrder[i]
		if !actual.Equals(v) {
			t.Errorf("Vaporization %d: expected %s, got %s\n", i+1, v, actual)
		}
	}
}

func c(x, y int) coord {
	return coord{x, y, 0.0, 0.0}
}

func Test_ClockwisePolar(t *testing.T) {
	tests := []struct {
		input, expected float64
	}{
		{90.0, 0},
		{0.0, 90.0},
		{270.0, 180.0},
		{180.0, 270.0},
		{1.0, 89.0},
		{359.0, 91.0},
	}

	for i, test := range tests {
		actual := clockwisePolar(test.input)
		if !isAcceptable(actual - test.expected) {
			t.Errorf("Test %d: for %.5f expected %.5f, got %.5f, diff %.5f\n", i+1, test.input, test.expected, actual, test.expected-actual)
		}
	}
}

func isAcceptable(delta float64) bool { return math.Abs(delta) < 0.005 }
