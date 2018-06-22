package main

import "fmt"

// could be off by one, can't remember exact test cases.

func longestCylce(input []int) int {
	largestCycle := 1
	for i := range input {
		num := input[i]
		currentCycle := 1
		for num != i {
			num = input[num]
			currentCycle++
		}
		if currentCycle > largestCycle {
			largestCycle = currentCycle
		}
	}
	return largestCycle
}

func main() {
	testCycle := []int{2, 3, 4, 1, 5, 0}
	fmt.Println(longestCylce(testCycle))
	return
}
