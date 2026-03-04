package main

import "fmt"

func main() {
	var testCases int

	fmt.Scan(&testCases)

	for t := 0; t < testCases; t++ {
		var length int

		fmt.Scan(&length)

		largestNumber := -1
		for i := 0; i < length; i++ {
			var value int
			fmt.Scan(&value)

			if value >= largestNumber {
				largestNumber = value
			}
		}

		fmt.Println(largestNumber * length)
	}
}
