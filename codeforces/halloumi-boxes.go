package main

import "fmt"

func main() {
	var testCases int

	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var n, k int

		fmt.Scan(&n, &k)

		nonDecreasing := true 
		lastValue := 1
		for j := 0; j < n; j++ {
			var value int
			
			fmt.Scan(&value)

			if value < lastValue {
				nonDecreasing = false
			}

			lastValue = value
		}

		if k == 1 && !nonDecreasing {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}

	
}
