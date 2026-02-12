package main

import (
	"fmt"
)

func main() {
	var velasAcesas, quantidadeParaNovaVela int
	
	fmt.Scan(&velasAcesas, &quantidadeParaNovaVela)

	horas, velasApagadas := 0, 0
	for velasAcesas > 0 || velasApagadas >= quantidadeParaNovaVela {
		if velasAcesas > 0 {
			horas++
			velasAcesas--
			velasApagadas++
		} else if velasApagadas >= quantidadeParaNovaVela {
			velasAcesas++
			velasApagadas -= quantidadeParaNovaVela
		}
	}
	fmt.Println(horas)
}
