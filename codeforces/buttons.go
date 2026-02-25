package main

import "fmt"

func main() {
	var testCases int

	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var a, b, c int

		fmt.Scan(&a, &b, &c)
		if c % 2 == 0 {
				if a > b {
					fmt.Println("First")
				} else {
					fmt.Println("Second")
				}
		} else {
			if b > a {
				fmt.Println("Second")
			} else {
				fmt.Println("First")
			}
		}
	}
}
