package main

import (
	"fmt"
	"bufio"
	"os"
)

const (
	um = "1"
	zero = "0"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int

	fmt.Fscan(reader, &n)
	
	for i := 0; i < n; i++ {
		var s string

		fmt.Fscan(reader, &s)
		indexes := make(map[string]int)

		for j := 0; j < len(s); j++ {
			indexes[string(s[j])]++
		}

		t := ""
		for j := 0; j < len(s); j++ {
			needsToBreak := false
			switch string(s[j]) {
				case um:
					if indexes[zero] > 0 {
						t += zero
						indexes[zero]--
					} else {
						needsToBreak = true
					}
				case zero:
					if indexes[um] > 0 {
						t += um
						indexes[um]--
					} else {
						needsToBreak = true
					}
			}

			if needsToBreak {
				break
			}
		}

		fmt.Fprintln(writer, len(s) - len(t))
	}
	
}
