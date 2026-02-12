package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int

	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, k int

		fmt.Fscan(reader, &n, &k)

		encontrado := false
		for j := 0; j < n; j++ {
			var valor int

			fmt.Fscan(reader, &valor)

			if valor == k {
				encontrado = true
			}
		}

		switch {
			case encontrado:
				fmt.Fprintln(writer, "YES")
			default:
				fmt.Fprintln(writer, "NO")
		}

		}
}

