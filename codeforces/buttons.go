package main

import "fmt"

func main() {
	var testCases int

	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var a, b, c int

		fmt.Scan(&a, &b, &c)

		switch (a + b + c) % 2{
			case 0: 
				fmt.Println("Second")
			default:
				fmt.Println("First")
		}
	}
}
