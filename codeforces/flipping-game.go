package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	ones := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
		if numbers[i] == 1 {
			ones++
		}
	}
	
	if ones == n {
		fmt.Println(n - 1)
		return
	}

	maxGain := -1        
	currentGain := 0     

	for i := 0; i < n; i++ {
		val := 0
		if numbers[i] == 0 {
			val = 1  
		} else {
			val = -1 
		}

		currentGain += val

		if currentGain < 0 {
			currentGain = 0
		}

		if currentGain > maxGain {
			maxGain = currentGain
		}
	}
	fmt.Println(ones + maxGain)
}
