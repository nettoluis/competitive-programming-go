package main

import "fmt"

func main() {
	var testCases int

	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var n int

		fmt.Scan(&n)

		fmt.Println(n-1)
	}
}
