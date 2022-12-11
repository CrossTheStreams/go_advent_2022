package main

import "testing"

func TestRuckSackFromLine(t *testing.T) {
	line := "vJrwpWtwJgWrhcsFMMfFFhFp"
	rucksack := ruckSackFromLine(line)
	expectedRucksack := Rucksack{
		compartment1: "vJrwpWtwJgWr",
		compartment2: "hcsFMMfFFhFp",
	}
	if rucksack != expectedRucksack {
		t.Errorf("got %+v, wanted %+v", rucksack, expectedRucksack)
	}
}

func TestFindItem(t *testing.T) {
	rucksack := Rucksack{
		compartment1: "abc",
		compartment2: "cde",
	}
	got := findItem(rucksack)
	want := "c"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestPrioritySumFromInput(t *testing.T) {
	got := prioritySumFromInput("small_input.txt")
	want := 157
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
