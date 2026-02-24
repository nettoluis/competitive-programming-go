package main

import "fmt"

const VALOR_MAXIMO = 6

func retornaMaior(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func mdc(a, b int) int {
	for b != 0 {
		resto := a % b
		a = b
		b = resto
	}

	return a
}

func main() {
	var a, b int
	
	fmt.Scan(&a, &b)

	maior := retornaMaior(a, b)
	possibilidades := 6 - (maior - 1)
	divisor := mdc(possibilidades, VALOR_MAXIMO)

	fmt.Printf("%d/%d\n", possibilidades / divisor, VALOR_MAXIMO / divisor)
}
