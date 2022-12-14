package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var priorityLookup = map[string]int{}

type Rucksack struct {
	compartment1 string
	compartment2 string
}

func findItem(rucksack Rucksack) string {
	sideA := rucksack.compartment1
	sideB := rucksack.compartment2
	lookupMap := map[string]string{}
	shared := ""
	for i := 0; i < len(sideA); i++ {
		char := sideA[i : i+1]
		lookupMap[char] = char
	}
	for i := 0; i < len(sideB); i++ {
		char := sideB[i : i+1]
		if match, ok := lookupMap[char]; ok {
			shared = match
		}
	}
	return shared
}

func findGroupSharedItem(group []string) string {
	var sharedItem string
	itemsInLines := map[string]map[int]bool{}
	for i := 0; i < len(group); i++ {
		line := group[i]
		for j := 0; j < len(line); j++ {
			item := string(line[j])
			if itemsInLines[item] == nil {
				itemsInLines[item] = map[int]bool{}
			}
			itemsInLines[item][i] = true
			inAllThreeGroups := (itemsInLines[item][0] && itemsInLines[item][1] && itemsInLines[item][2])
			if inAllThreeGroups {
				sharedItem = item
				break
			}
		}
		if sharedItem != "" {
			break
		}
	}
	return sharedItem
}

func ruckSackFromLine(line string) Rucksack {
	midPoint := (len(line) / 2)
	firstHalf := line[0:midPoint]
	secondHalf := line[midPoint:]
	return Rucksack{
		compartment1: firstHalf,
		compartment2: secondHalf,
	}
}

func createPriorityLookupMap() map[string]int {
	priorityLookup := map[string]int{}
	var ch byte
	value := 0
	for ch = 'a'; ch <= 'z'; ch++ {
		value++
		priorityLookup[string(ch)] = value

	}
	for ch = 'A'; ch <= 'Z'; ch++ {
		value++
		priorityLookup[string(ch)] = value
	}
	return priorityLookup
}

func readFile(fileName string) []byte {
	fileRows, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return fileRows
}

func prioritySumFromInput(fileName string) int {
	priorityMap := createPriorityLookupMap()
	sum := 0
	fileRows := readFile(fileName)
	fileString := string(fileRows)
	stringRows := strings.Split(fileString, "\n")
	for i := 0; i < len(stringRows); i++ {
		stringRow := stringRows[i]
		rucksack := ruckSackFromLine(stringRow)
		char := findItem(rucksack)
		itemPriority := priorityMap[char]
		sum += itemPriority
	}
	return sum
}

func prioritySumForGroupsOfThree(fileName string) int {
	priorityMap := createPriorityLookupMap()
	sum := 0
	fileRows := readFile(fileName)
	fileString := string(fileRows)
	stringRows := strings.Split(fileString, "\n")
	groupLines := []string{}
	for i := 0; i < len(stringRows); i++ {
		stringRow := stringRows[i]
		groupLines = append(groupLines, stringRow)
		if len(groupLines) == 3 {
			char := findGroupSharedItem(groupLines)
			itemPriority := priorityMap[char]
			sum += itemPriority
			groupLines = []string{}
		}
	}
	return sum
}

func main() {
	fmt.Println("Part1:")
	fmt.Println(prioritySumFromInput("small_input.txt"))
	fmt.Println("Part2:")
	fmt.Println(prioritySumForGroupsOfThree("small_input.txt"))
}
