package main

import "testing"

func TestFindMaxCalories(t *testing.T) {
	elfCalories := []int{6000, 4000, 11000, 24000, 10000}
	got := findMaxCalories(elfCalories)
	want := 24000
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestFindTopKCaloriesSu(t *testing.T) {
	elfCalories := []int{6000, 4000, 11000, 24000, 10000}
	got := findTopKCaloriesSum(elfCalories, 3)
	want := 45000
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
