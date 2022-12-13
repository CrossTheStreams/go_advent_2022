package main

import (
	"testing"
)

func TestCountContainingSections(t *testing.T) {
	got := countContainingSections("small_input.txt")
	want := 2
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCountOverlappingSections(t *testing.T) {
	got := countOverlappingSections("small_input.txt")
	want := 4
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSectionIDToTuple(t *testing.T) {
	sectionID := "2-4"
	got := sectionIDToTuple(sectionID)
	want := []int{2, 4}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("got[%d] != wanted[%d], with got[%d] == %+v and want[%d] == %+v", i, i, i, got[i], i, want[i])
		}
	}
}

func TestSlicesContainsOther(t *testing.T) {
	a := []int{1, 4}
	b := []int{2, 4}
	got := sliceContainsOther(a, b)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestSlicesContainsOtherOverlap(t *testing.T) {
	a := []int{1, 4}
	b := []int{4, 7}
	got := sliceContainsOther(a, b)
	want := false
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestSlicesContainsOtherNoOverlap(t *testing.T) {
	a := []int{1, 4}
	b := []int{5, 8}
	got := sliceContainsOther(a, b)
	want := false
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestSlicesDontOverlap(t *testing.T) {
	a := []int{1, 4}
	b := []int{5, 8}
	got := slicesOverlap(a, b)
	want := false
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestSlicesOverlap(t *testing.T) {
	a := []int{1, 4}
	b := []int{4, 8}
	got := slicesOverlap(a, b)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestSlicesOverlapMore(t *testing.T) {
	a := []int{1, 7}
	b := []int{3, 6}
	got := slicesOverlap(a, b)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
