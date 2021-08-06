package main

import (
	"fmt"
)

func uniqueNames(a, b []string) []string {
	var merged = append(a, b...)

	var result []string

	allKeys := make(map[string]bool)
	for _, item := range merged {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			result = append(result, item)
		}
	}

	return result
}

func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}