package main

import "testing"

func TestFindMaxCalories(t *testing.T) {
	elfCalories := []int{6000, 4000, 11000, 24000, 10000}
	got := findMaxCalories(elfCalories)
	want := 2400
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestFindTopKCalories(t *testing.T) {
	elfCalories := []int{6000, 4000, 11000, 24000, 10000}
	got := findTopKCalories(elfCalories, 3)
	want := 2400
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
