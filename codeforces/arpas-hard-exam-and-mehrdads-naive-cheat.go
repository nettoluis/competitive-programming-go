package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	var saida int

	if n == 0 {
		fmt.Println(1)
		return
	}

	switch n % 4 {
		case 0:
			saida = 6
		case 1:
			saida = 8
		case 2:
			saida = 4
		case 3:
			saida = 2
	}

	fmt.Println(saida)
}
