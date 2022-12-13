package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part1:")
	fmt.Println(countContainingSections("input.txt"))
	fmt.Println("Part2:")
	fmt.Println(countOverlappingSections("input.txt"))
}

func countContainingSections(fileName string) int {
	count := 0
	fileRows := readFile(fileName)
	fileString := string(fileRows)
	stringRows := strings.Split(fileString, "\n")
	for i := 0; i < len(stringRows); i++ {
		rowSectionStrings := strings.Split(stringRows[i], ",")
		tupleA := sectionIDToTuple(rowSectionStrings[0])
		tupleB := sectionIDToTuple(rowSectionStrings[1])
		if sliceContainsOther(tupleA, tupleB) {
			count += 1
		}
	}
	return count
}

func countOverlappingSections(fileName string) int {
	count := 0
	fileRows := readFile(fileName)
	fileString := string(fileRows)
	stringRows := strings.Split(fileString, "\n")
	for i := 0; i < len(stringRows); i++ {
		rowSectionStrings := strings.Split(stringRows[i], ",")
		tupleA := sectionIDToTuple(rowSectionStrings[0])
		tupleB := sectionIDToTuple(rowSectionStrings[1])
		if slicesOverlap(tupleA, tupleB) {
			count += 1
		}
	}
	return count
}

func readFile(fileName string) []byte {
	fileRows, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return fileRows
}

func sectionIDToTuple(sectionID string) []int {
	idStrings := strings.Split(sectionID, "-")
	tuple := []int{}
	for i := 0; i < len(idStrings); i++ {
		char := string(idStrings[i])
		n, err := strconv.Atoi(char)
		if err != nil {
			break
		}
		tuple = append(tuple, n)
	}
	return tuple
}

func sliceContainsOther(a []int, b []int) bool {
	aContainsB := (a[0] <= b[0] && a[1] >= b[1])
	bContainsA := (b[0] <= a[0] && b[1] >= a[1])
	return aContainsB || bContainsA
}

func slicesOverlap(a []int, b []int) bool {
	return (a[1] >= b[0] && a[0] <= b[1]) || (b[1] >= a[0] && b[0] <= a[1])
}
