package main

import "fmt"

func main() {
	var qtd int

	fmt.Scan(&qtd)

	for i := 0; i < qtd; i++ {
		var n int
		var s string

		fmt.Scan(&n)
		fmt.Scan(&s)

		resposta := 0
		abertos := 0
		for i := 0; i < len(s); i++ {
			if string(s[i]) == "(" {
				abertos++
			} else {
				abertos--
				if (abertos < 0) {
					abertos = 0
					resposta++
				}
			}
		}

		fmt.Println(resposta)
	}
}
