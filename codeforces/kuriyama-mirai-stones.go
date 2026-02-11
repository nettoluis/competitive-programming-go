package main

import (
	"fmt"
	"os"
	"bufio"
)

func soma(s []int) []uint64 {
	somas := make([]uint64, len(s)+1)

	for i := 0; i < len(s); i++ {
		somas[i+1] = somas[i] + uint64(s[i])
	}

	return somas
}

func mergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	primeiraMetade := s[:len(s) / 2]
	segundaMetade := s[len(s) / 2:]

	primeiraMetade = mergeSort(primeiraMetade)
	segundaMetade = mergeSort(segundaMetade)

	return merge(primeiraMetade, segundaMetade)
}

func merge(a, b []int) []int {
	c := make([]int, len(a) + len(b))
	idC := 0

	for len(a) != 0 && len(b) != 0 {
		if a[0] < b[0] {
			c[idC] = a[0]
			a = a[1:]
		} else {
			c[idC] = b[0]
			b = b[1:]
		}
			idC++
	}

	for len(a) != 0 {
		c[idC] = a[0]
		a = a[1:]
		idC++
	}

	for len(b) != 0 {
		c[idC] = b[0]
		b = b[1:]
		idC++
	}

	return c
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var	n int

	fmt.Fscan(reader,&n)
	pedras := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader,&pedras[i])
	}

	pedrasOrdenadas := mergeSort(pedras)

	prefixSumsOriginal := soma(pedras)
	prefixSumsOrdenada := soma(pedrasOrdenadas)

	var m int

	fmt.Fscan(reader,&m)

	for j := 0; j < m; j++ {
		var tipo, l, r int

		fmt.Fscan(reader,&tipo, &l, &r)

		switch {
			case tipo == 1:
				fmt.Fprintln(writer, prefixSumsOriginal[r] - prefixSumsOriginal[l-1])
			default:
				fmt.Fprintln(writer, prefixSumsOrdenada[r] - prefixSumsOrdenada[l-1])
		}
	}
}
