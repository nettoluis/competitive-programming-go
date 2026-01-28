package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)

		diferencas := make(map[int]int)
		var contaPares int64 

		for j := 0; j < n; j++ {
			var num int
			fmt.Fscan(reader, &num)

			val := num - j
			contaPares += int64(diferencas[val])
			diferencas[val]++
		}

		fmt.Fprintln(writer, contaPares)
	}
}
