package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := os.Getenv("FILENAME")
	partOneMap := map[string]int{
		"A X": 4,
		"B X": 1,
		"C X": 7,
		"A Y": 8,
		"B Y": 5,
		"C Y": 2,
		"A Z": 3,
		"B Z": 9,
		"C Z": 6,
	}

	partTwoMap := map[string]int{
		"A X": 3,
		"B X": 1,
		"C X": 2,
		"A Y": 4,
		"B Y": 5,
		"C Y": 6,
		"A Z": 8,
		"B Z": 9,
		"C Z": 7,
	}

	calculateScore(filename, partOneMap)
	calculateScore(filename, partTwoMap)
}

func calculateScore(filename string, m map[string]int) {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var score = 0

	for scanner.Scan() {
		line := scanner.Text()
		score += m[line]
	}

	fmt.Println(score)
}
