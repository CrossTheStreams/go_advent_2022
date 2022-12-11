package main

import "testing"

func TestRockPaper(t *testing.T) {
	got := handPoints("rock", "paper")
	want := 8
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPaperRock(t *testing.T) {
	got := handPoints("paper", "rock")
	want := 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestScissorsPaper(t *testing.T) {
	got := handPoints("scissors", "paper")
	want := 2
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestScoreFromInputPartOne(t *testing.T) {
	got := scoreFromInput("small_input.txt", false)
	want := 19
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestScoreFromInputPartTwo(t *testing.T) {
	got := scoreFromInput("small_input.txt", true)
	want := 19
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
