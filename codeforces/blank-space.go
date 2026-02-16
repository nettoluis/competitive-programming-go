package main

import "fmt"

func main() {
	var testCases int

	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var length int

		fmt.Scan(&length)

		largestSequence := 0
		counter := 0

		for j := 0; j < length; j++ {
			var value int

			fmt.Scan(&value)

			if value == 0 {
				counter++
			} else {
				if counter > largestSequence {
					largestSequence = counter
				}
				counter = 0
			}
		}

		if counter > largestSequence {
			largestSequence = counter
		}

		fmt.Println(largestSequence)
	}
}
