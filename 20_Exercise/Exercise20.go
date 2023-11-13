package main

import (
	"fmt"
)

func main() {

	var x int = 10

	// initalization of a slice
	xi := []int{1, 2, 3, 4, 5}

	// initalization of a map
	m := map[string]int{
		"James Bond":  42,
		"Joel Lovely": 17,
	}
	// Making basic For-Loop
	for i := 0; i < x; i++ {
		fmt.Printf("Basic Commands of a For-Loop\n")
	}

	// Making basic For-Range-Loop
	for _, value := range xi {
		fmt.Printf("Basic Commands of using a slice with a For-Range-Loop %v\n", value)
	}

	for key, value := range m {
		fmt.Printf("Basic commands of using a map with a For-Range-Loop Notice this will be unordered %v with number %v\n", key, value)

	}

}
