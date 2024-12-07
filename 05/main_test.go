package main

import (
	"testing"
)

func TestP1(t *testing.T) {
	testinput := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

	result := part1(testinput)
	if result != 143 {
		t.Errorf("failed. got %d, should be 143", result)
	}
}

func TestP2(t *testing.T) {
	testinput := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

	result := part2(testinput)
	if result != 123 {
		t.Errorf("failed. got %d, should be 123", result)
	}
}
