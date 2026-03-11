package main

import "fmt"

func main() {
	var testCases int

	fmt.Scan(&testCases)

	for t := 0; t < testCases; t++ {
		var a, b, c, d int
		
		fmt.Scan(&a, &b, &c, &d)

		if a == b && b == c && c == d {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
