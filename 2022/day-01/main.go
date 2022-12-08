package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read file content
	filename := os.Getenv("FILENAME")

	partOne(filename)
	partTwo(filename)
}

func partOne(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var max = 0
	var current = 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			number, _ := strconv.Atoi(line)
			current += number
			if current > max {
				max = current
			}
		} else {
			current = 0
		}
	}

	fmt.Println(max)
}

func partTwo(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var currentSum = 0
	leaderboard := &IntHeap{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			number, _ := strconv.Atoi(line)
			currentSum += number
		} else {
			heap.Push(leaderboard, currentSum)
			currentSum = 0
		}
	}
	heap.Push(leaderboard, currentSum)

	var top3 = 0

	for i := 0; i < 3; i++ {
		top3 += heap.Pop(leaderboard).(int)
	}
	fmt.Println(top3)

}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
