package main

import (
	"testing"
)

func Test_CalculateFuel(t *testing.T) {
	tests := []struct {
		reactions    []string
		expectedOre  int
		expectedFuel int
	}{
		{[]string{"157 ORE => 5 NZVS", "165 ORE => 6 DCFZ", "44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL", "12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ", "179 ORE => 7 PSHF", "177 ORE => 5 HKGWZ", "7 DCFZ, 7 PSHF => 2 XJWVT", "165 ORE => 2 GPVTF", "3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT"}, 13312, 82892753},
		{[]string{"2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG", "17 NVRVD, 3 JNWZP => 8 VPVL", "53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL", "22 VJHF, 37 MNCFX => 5 FWMGM", "139 ORE => 4 NVRVD", "144 ORE => 7 JNWZP", "5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC", "5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV", "145 ORE => 6 MNCFX", "1 NVRVD => 8 CXFTF", "1 VJHF, 6 MNCFX => 4 RFSQX", "176 ORE => 6 VJHF"}, 180697, 5586022},
		{[]string{"171 ORE => 8 CNZTR", "7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL", "114 ORE => 4 BHXH", "14 VRPVC => 6 BMBT", "6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL", "6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT", "15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW", "13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW", "5 BMBT => 4 WPTQ", "189 ORE => 9 KTJDG", "1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP", "12 VRPVC, 27 CNZTR => 2 XDBXC", "15 KTJDG, 12 BHXH => 5 XCVML", "3 BHXH, 2 VRPVC => 7 MZWV", "121 ORE => 7 VRPVC", "7 XCVML => 6 RJRHP", "5 BHXH, 4 VRPVC => 5 LTCX"}, 2210736, 460664},
	}

	for i, test := range tests {
		n := newNanofactory()
		for _, r := range test.reactions {
			n.addReaction(r)
		}

		actualOre := n.calculateOre(c("FUEL", 1))
		if test.expectedOre != actualOre {
			t.Errorf("Test %d: expected %d ore, actual %d\n", i+1, test.expectedOre, actualOre)
		}

		actualFuel := n.calculateFuel(1000000000000, 0, 1000000000000)
		if test.expectedFuel != actualFuel {
			t.Errorf("Test %d: expected %d fuel, actual %d\n", i+1, test.expectedFuel, actualFuel)
		}
	}

}

func Test_CalculateOre(t *testing.T) {
	tests := []struct {
		reactions []string
		expected  int
	}{
		{[]string{"10 ORE => 10 A", "1 ORE => 1 B", "7 A, 1 B => 1 C", "7 A, 1 C => 1 D", "7 A, 1 D => 1 E", "7 A, 1 E => 1 FUEL"}, 31},
		{[]string{"9 ORE => 2 A", "8 ORE => 3 B", "7 ORE => 5 C", "3 A, 4 B => 1 AB", "5 B, 7 C => 1 BC", "4 C, 1 A => 1 CA", "2 AB, 3 BC, 4 CA => 1 FUEL"}, 165},
		{[]string{"157 ORE => 5 NZVS", "165 ORE => 6 DCFZ", "44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL", "12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ", "179 ORE => 7 PSHF", "177 ORE => 5 HKGWZ", "7 DCFZ, 7 PSHF => 2 XJWVT", "165 ORE => 2 GPVTF", "3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT"}, 13312},
		{[]string{"2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG", "17 NVRVD, 3 JNWZP => 8 VPVL", "53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL", "22 VJHF, 37 MNCFX => 5 FWMGM", "139 ORE => 4 NVRVD", "144 ORE => 7 JNWZP", "5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC", "5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV", "145 ORE => 6 MNCFX", "1 NVRVD => 8 CXFTF", "1 VJHF, 6 MNCFX => 4 RFSQX", "176 ORE => 6 VJHF"}, 180697},
		{[]string{"171 ORE => 8 CNZTR", "7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL", "114 ORE => 4 BHXH", "14 VRPVC => 6 BMBT", "6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL", "6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT", "15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW", "13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW", "5 BMBT => 4 WPTQ", "189 ORE => 9 KTJDG", "1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP", "12 VRPVC, 27 CNZTR => 2 XDBXC", "15 KTJDG, 12 BHXH => 5 XCVML", "3 BHXH, 2 VRPVC => 7 MZWV", "121 ORE => 7 VRPVC", "7 XCVML => 6 RJRHP", "5 BHXH, 4 VRPVC => 5 LTCX"}, 2210736},
	}

	for i, test := range tests {
		n := newNanofactory()
		for _, r := range test.reactions {
			n.addReaction(r)
		}

		actual := n.calculateOre(c("FUEL", 1))
		if test.expected != actual {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, test.expected, actual)
		}
	}
}

func Test_Parse(t *testing.T) {
	tests := []struct {
		input  string
		inputs []*chemical
		output *chemical
	}{
		{"9 ORE => 2 A", []*chemical{c("ORE", 9)}, c("A", 2)},
		{"8 ORE => 3 B", []*chemical{c("ORE", 8)}, c("B", 3)},
		{"7 ORE => 5 C", []*chemical{c("ORE", 7)}, c("C", 5)},
		{"3 A, 4 B => 1 AB", []*chemical{c("A", 3), c("B", 4)}, c("AB", 1)},
		{"5 B, 7 C => 1 BC", []*chemical{c("B", 5), c("C", 7)}, c("BC", 1)},
		{"4 C, 1 A => 1 CA", []*chemical{c("C", 4), c("A", 1)}, c("CA", 1)},
		{"2 AB, 3 BC, 4 CA => 1 FUEL", []*chemical{c("AB", 2), c("BC", 3), c("CA", 4)}, c("FUEL", 1)},
	}

	n := newNanofactory()
	for i, test := range tests {
		inputs, output := n.parse(test.input)

		if !test.output.equals(output) {
			t.Errorf("Test %d: expected %s, actual %s\n", i+1, test.output, output)
		}

		if len(test.inputs) != len(inputs) {
			t.Errorf("Test %d: expected %d inputs, got %d\n", i+1, len(test.inputs), len(inputs))
		} else {
			for j, input := range inputs {
				if !test.inputs[j].equals(input) {
					t.Errorf("Test %d: expected %s, actual %s at %d\n", i+1, test.inputs[j], input, j)
				}
			}
		}
	}
}
