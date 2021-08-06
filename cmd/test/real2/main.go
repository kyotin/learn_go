package main

import (
	"fmt"
)

func findMaxSum(numbers []int) int {
	first := -1
	second := -1

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > first {
			second = first
			first = numbers[i]
		} else if numbers[i] > second {
			second = numbers[i]
		}
	}

	return first + second
}

func main() {
	list := []int { 5, 9, 7, 11 }
	fmt.Println(findMaxSum(list))
}