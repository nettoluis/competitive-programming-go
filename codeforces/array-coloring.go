package main

import "fmt"

func main() {
	var t int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int

		fmt.Scan(&n)

		pares, impares := 0, 0
		for j := 0; j < n; j++ {
			var valor int

			fmt.Scan(&valor)

			switch {
				case valor % 2 == 0:
					pares++
				default:
					impares++
			}
		}

		switch {
			case impares % 2 == 0:
				fmt.Println("YES")
			default:
				fmt.Println("NO")
		}
	}
}
