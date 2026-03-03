package main

import "fmt"

const MINUTES_A_DAY = 1440
func main() {
	var testCases int

	fmt.Scan(&testCases)

	for t := 0; t < testCases; t++ {
		var hours, minutes int

		fmt.Scan(&hours, &minutes)

		fmt.Println(MINUTES_A_DAY - ((60 * hours) + minutes))
	}
}
