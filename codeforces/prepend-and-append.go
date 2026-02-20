package main

import (
	"fmt"
	"strconv"
)

func main() {
	var testCases int

	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var length int
		var number string

		fmt.Scan(&length)
		fmt.Scan(&number)

		counter := 0
		for j := 0; j < length / 2; j++ {
			atual, err1 := strconv.Atoi(string(number[j]))
			oposto, err2 := strconv.Atoi(string(number[length-1-j]))
			if err1 != nil || err2 != nil {
				panic(err1)
				panic(err2)
			}
			if (atual ^ oposto) == 1 {
				counter += 2
			} else {
				break
			}
		}

		fmt.Println(length - counter)
	}
}
