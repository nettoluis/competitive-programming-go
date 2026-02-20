package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var length int
		var binary string

		fmt.Scan(&length)
		fmt.Scan(&binary)

		first, second := "1", "1"
		oneAppeared := false
		for j := 1; j < len(binary); j++ {
			switch string(binary[j]) {
				case "2":
					if !oneAppeared {
						first += "1"
						second += "1"
					} else {
						first += "0"
						second += "2"
					}
				case "0":
					first += "0"
					second += "0"
				default:
					if !oneAppeared {
						first += "1"
						second += "0"
					} else {
						first += "0"
						second += "1"
					}

					oneAppeared = true
			}
		}

		fmt.Printf("%s\n%s\n", first, second)
	}
}
