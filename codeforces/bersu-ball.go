package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	saida := 0
	switch {
		case a < 0:
			saida = -a
		case a == 0:
			saida = 0
		case a > 0:
			saida = a
	}

	return saida
}

func main() {
	var num_boys int

	fmt.Scan(&num_boys)
	boys := make([]int, num_boys)

	for i := 0; i < len(boys); i++ {
		fmt.Scan(&boys[i])
	}

	var num_girls int

	fmt.Scan(&num_girls)
	girls := make([]int, num_girls)

	for i := 0; i < len(girls); i++ {
		fmt.Scan(&girls[i])
	}

	sort.Ints(boys)
	sort.Ints(girls)


	pairs := 0
	for i := 0; i < len(girls); i++ {
		for j := 0; j < len(boys); j++ {
			if abs(girls[i] - boys[j]) <= 1 {
				boys[j] = 102
				pairs++
				break
			}
		}
	}

	fmt.Println(pairs)
}
