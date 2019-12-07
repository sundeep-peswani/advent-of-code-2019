package main

import "testing"

func TestCountRelationships(t *testing.T) {
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
	}
	tests := []struct{
		object string
		expected int
	}{
		{"D", 3},
		{"L", 7},
		{"COM", 0},
	}

	u := newUniversalOrbitMap()
	u.load(input)

	actual := u.countRelationships()
	expected := 42
	if actual != expected {
		t.Errorf("Expected %d, actual %d\n", expected, actual)
	}

	for _, test := range tests {
		obj := u.get(test.object)
		if obj == nil {
			t.Errorf("Unable to find object labelled %s\n", test.object)
		} else {
			actual = obj.countOrbits()
			if actual != test.expected {
				t.Errorf("Expected %d for %s, actual %d\n", test.expected, test.object, actual)
			}
		}
	}
}