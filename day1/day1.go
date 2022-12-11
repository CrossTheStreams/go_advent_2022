package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	elfCalories := caloriesArrayFromInput("input.txt")
	fmt.Println("Part 1:")
	fmt.Println(findMaxCalories(elfCalories))
	fmt.Println("Part 2:")
	fmt.Println(findTopKCaloriesSum(elfCalories, 3))
}

func findMaxCalories(elfCalories []int) int {
	maxCalories := 0
	for i := 0; i < len(elfCalories); i++ {
		if elfCalories[i] > maxCalories {
			maxCalories = elfCalories[i]
		}
	}
	return maxCalories
}

func findTopKCaloriesSum(elfCalories []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))
	caloriesSum := 0
	for i := 0; i < k; i++ {
		caloriesSum += elfCalories[i]
	}
	return caloriesSum
}

func caloriesArrayFromInput(fileName string) []int {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	text := string(fileContent)
	calorieStrings := strings.Split(text, "\n")
	elfCalories := []int{}
	caloriesSum := 0
	for i := 0; i < len(calorieStrings); i++ {
		calorieString := calorieStrings[i]
		if calorieString == "" {
			elfCalories = append(elfCalories, caloriesSum)
			caloriesSum = 0
			continue
		}
		calorieValue, err := strconv.Atoi(calorieString)
		if err != nil {
			fmt.Print("Well that was wild, there's invalid data in input.txt")
			os.Exit(1)
		} else {
			caloriesSum += calorieValue
		}
	}
	return elfCalories
}
