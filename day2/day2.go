package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var col1Map = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}

// part 1 uses col2 to indicate what hand to pick
var col2Map = map[string]string{
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

// part 2 uses col2 to indicate the desired result of the match instead
var col2ConditionMap = map[string]string{
	"X": "lose",
	"Y": "draw",
	"Z": "win",
}

// opponent hand => condition => appropriate hand for me
var handConditionPairMap = map[string]map[string]string{
	"rock": {
		"lose": "scissors",
		"draw": "rock",
		"win":  "paper",
	},
	"paper": {
		"lose": "rock",
		"draw": "paper",
		"win":  "scissors",
	},
	"scissors": {
		"lose": "paper",
		"draw": "scissors",
		"win":  "rock",
	},
}

var conditionValues = map[string]int{
	"lose": 0,
	"draw": 3,
	"win":  6,
}

var shapePointsMap = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

func main() {
	score := scoreFromInput("input.txt", true)
	fmt.Println(score)
}

func handPoints(theirHand string, myHand string) int {
	totalPoints := 0
	shapePoints := shapePointsMap[myHand]
	if theirHand == myHand {
		totalPoints = shapePoints + 3
	}
	if myHand == "rock" {
		switch theirHand {
		case "paper":
			totalPoints = shapePoints
		case "scissors":
			totalPoints = shapePoints + 6
		}
	}
	if myHand == "paper" {
		switch theirHand {
		case "scissors":
			totalPoints = shapePoints
		case "rock":
			totalPoints = shapePoints + 6
		}
	}
	if myHand == "scissors" {
		switch theirHand {
		case "rock":
			totalPoints = shapePoints
		case "paper":
			totalPoints = shapePoints + 6
		}
	}
	return totalPoints
}

func scoreFromInput(fileName string, col2gameConditions bool) int {
	score := 0
	allRounds := handsFromInput(fileName, col2gameConditions)
	for i := 0; i < len(allRounds); i++ {
		hands := allRounds[i]
		theirHand := hands[0]
		myHand := hands[1]
		score += handPoints(theirHand, myHand)
	}
	return score
}

func handsFromInput(fileName string, col2gameConditions bool) [][]string {
	var hands [][]string
	fileRows := readFile(fileName)
	fileString := string(fileRows)
	stringRows := strings.Split(fileString, "\n")
	for i := 0; i < len(stringRows); i++ {
		row := stringRows[i]
		rowChars := strings.Split(row, " ")
		theirLetter := rowChars[0]
		myLetter := rowChars[1]
		theirHand := col1Map[theirLetter]
		myHand := col2Map[myLetter]
		if col2gameConditions {
			gameCondition := col2ConditionMap[myLetter]
			myHand = handConditionPairMap[theirHand][gameCondition]
			// stuff := fmt.Sprintf("%s %s %s", theirHand, myHand, gameCondition)
			// fmt.Println(stuff)
		}
		hand := []string{theirHand, myHand}
		hands = append(hands, hand)
	}
	return hands
}

func readFile(fileName string) []byte {
	fileRows, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return fileRows
}
