package main

import (
	"fmt"
)

const VALOR_MAXIMO = 50

var isPrime[VALOR_MAXIMO+1]bool
var primos []int = crivoEratostenes(VALOR_MAXIMO)

func crivoEratostenes(n int) []int {
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	isPrime[0] = false
	isPrime[1] = false

	for p := 2; p*p <= VALOR_MAXIMO; p++ {
		if isPrime[p] {
		for i := p * p; i <= VALOR_MAXIMO; i += p {
			isPrime[i] = false
			}
		}
	}

	primos := make([]int, VALOR_MAXIMO + 1)
	id := 0
	for i := 0; i <= VALOR_MAXIMO; i++ {
		if isPrime[i] {
			primos[id] = i
			id++
		}
	}

	return primos
}

func ehProximoPrimo(primos []int, n, m int) bool {
	for i := range primos {
		if primos[i] == n {
			return primos[i+1] == m
		}
	}

	return false
}

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	saida := "NO"
	if ehProximoPrimo(primos, n, m)	{
		saida = "YES"
	}

	fmt.Println(saida)
}
