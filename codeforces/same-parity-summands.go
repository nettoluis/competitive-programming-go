package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

const (
	MENOR_PAR = 2
	MENOR_IMPAR = 1
)

func construirSaida(n, k int) string {
	DIFERENCA_POR_PAR := n - (MENOR_PAR * (k - 1))
	DIFERENCA_POR_IMPAR := n - (MENOR_IMPAR * (k - 1))
	saida := ""

	if DIFERENCA_POR_IMPAR > 0 && DIFERENCA_POR_IMPAR % 2 != 0 {
		saida = strings.Repeat("1 ", k - 1)
		saida += fmt.Sprintf("%d", DIFERENCA_POR_IMPAR)
	}

	
	if DIFERENCA_POR_PAR > 0 && DIFERENCA_POR_PAR % 2 == 0 {
		saida = strings.Repeat("2 ", k - 1)
		saida += fmt.Sprintf("%d", DIFERENCA_POR_PAR)
	}

	return saida
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var testCases int
	
	fmt.Fscan(reader, &testCases)

	for i := 0; i < testCases; i++ {
		var n, k int

		fmt.Fscan(reader, &n, &k)

		switch n % 2 == 0 {
			case true:
				if (n >= (k * MENOR_PAR)) || ((n >= (k * MENOR_IMPAR) && k % 2 == 0)) {
					fmt.Fprintln(writer, "YES")
					saida := construirSaida(n, k)
					fmt.Fprintln(writer, saida)
				} else {
					fmt.Fprintln(writer, "NO")
				}
			case false:
				if (n >= k * MENOR_IMPAR) && k % 2 != 0 {
					fmt.Fprintln(writer, "YES")
					saida := construirSaida(n, k)
					fmt.Fprintln(writer, saida)
				} else {
					fmt.Fprintln(writer, "NO")
				}
		}
	}
}
