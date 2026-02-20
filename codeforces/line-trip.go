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

	var testCases int

	fmt.Fscan(reader, &testCases)

	for i := 0; i < testCases; i++ {
		var n, k int

		fmt.Fscan(reader, &n, &k)

		lastValue, maxDiff, currentDiff := 0, 0, 0
		for j := 0; j < n; j++ {
			var value int

			fmt.Fscan(reader, &value)

			currentDiff = value - lastValue

			if currentDiff > maxDiff {
				maxDiff = currentDiff
			}
			lastValue = value
		}

		fmt.Fprintln(writer, max(maxDiff, (k - lastValue) * 2))
	}

	
}
