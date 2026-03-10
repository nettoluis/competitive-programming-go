package main

import (
	"fmt"
)

func main() {
	calories := make(map[string]int)
	for c := 1; c <= 4; c++ {
		var value int

		fmt.Scan(&value)

		calories[fmt.Sprintf("%d", c)] = value
	}

	var sequence string

	fmt.Scan(&sequence)

	sum := 0
	for i := range sequence {
		sum += calories[string(sequence[i])]
	}

	fmt.Println(sum)
}
