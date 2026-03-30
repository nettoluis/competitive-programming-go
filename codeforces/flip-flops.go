package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var testCases int

	fmt.Fscan(reader, &testCases)

	for t := 0; t < testCases; t++ {
		var n, c, k int

		fmt.Fscan(reader, &n, &c, &k)

		monsters := make([]int, n)
		for i := range monsters {
			fmt.Fscan(reader, &monsters[i])
		}
		sort.Ints(monsters)

		for i := range monsters {
			if monsters[i] < c && k > 0 {
				difference := c - monsters[i]
				monsters[i] = min(c, monsters[i] + k)
				k = max(0, k - difference)
			}

			if monsters[i] > c {
				break
			}
			c += monsters[i]
		}

		fmt.Fprintln(writer, c)
	}
}
