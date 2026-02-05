package main

import "fmt"

const (
	NUM_OF_PAIRS = 4
	NUM_OF_DIFFERENT_COLORS_REQUIRED = 4
)
func main() {
	cores := make(map[int]bool)

	for i := 0; i < NUM_OF_PAIRS; i++ {
		var cor int

		fmt.Scan(&cor)

		cores[cor] = true
	}

	fmt.Println(NUM_OF_DIFFERENT_COLORS_REQUIRED - len(cores))
}
