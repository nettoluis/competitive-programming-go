package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	saida := ""
	for i := 0; i < n; i++ {
		if i != 0 {
			saida += " that "
			if i % 2 == 0 {
				saida += "I hate"
			} else {
				saida += "I love"
			} 
		} else {
			saida += "I hate"
		}
	}

	saida += " it"
	fmt.Println(saida)
}
