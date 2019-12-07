package main

import "testing"

func TestCountOrbits(t *testing.T) {
	input := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	tests := []struct{
		obj string
		expected int
	}{
		{"D", 3},
		{"L", 7},
		{"COM", 0},
	}

	u := newUniversalOrbitMap()
	u.load(input)

	for _, test := range tests {
		obj := u.get(test.obj)
		actual := obj.countOrbits()

		if actual != test.expected {
			t.Errorf("Expected %d for %s, actual %d\n", test.expected, test.obj, actual)
		}
	}
}

func TestCountOrbitalTransfers(t *testing.T) {
	input := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	tests := []struct{
		start, dest string
		expected int
	}{
		{"D", "COM", 2},
		{"E", "I", 0},
		{"L", "K", 1},
		{"L", "COM", 6},
		{"YOU", "SAN", 4},
		{"H", "E", 3},
	}

	u := newUniversalOrbitMap()
	u.load(input)

	for _, test := range tests {
		actual := u.countOrbitalTransfers(test.start, test.dest)
		if actual != test.expected {
			t.Errorf("Expected %d for %s to %s, actual %d\n", test.expected, test.start, test.dest, actual)
		}
	}
}