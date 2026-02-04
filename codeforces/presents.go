package main

import (
	"fmt"
	"strconv"
)

func main() {
	var qtd int

	fmt.Scan(&qtd)

	presentes := make([]int, qtd)
	for i := 0; i < qtd; i++ {
		var destinatario int

		fmt.Scan(&destinatario)

		presentes[destinatario - 1] = i + 1
	}

	saida := ""
	for i := 0; i < len(presentes); i++ {
		numeroString := strconv.Itoa(presentes[i])

		if i != 0 {
			saida += " " + numeroString
		} else {
			saida += numeroString
		}
	}
	fmt.Println(saida)
}
