package main

import "fmt"

func main() {
	var testCases int
	
	fmt.Scan(&testCases)

	for t := 0; t < testCases; t++ {
		var n int

		fmt.Scan(&n)

		saida := ""
		for i := 0; i < n; i ++ {
			if i == 0 {
				saida += fmt.Sprintf("%d", i+1)
			} else {
				saida += fmt.Sprintf(" %d", i+1)
			}
		}

		fmt.Println(saida)
	}
}
