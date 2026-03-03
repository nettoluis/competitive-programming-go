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

	for t := 0; t < testCases; t++ {
		var n, k int

		fmt.Fscan(reader, &n, &k)

		nums := make([]int, n)
		for i := 0; i < len(nums); i++ {
			fmt.Fscan(reader, &nums[i])
		}

		saida := -1
		for _, v := range nums {
			if v % k == 0 {
				saida = 0
				break
			}
		}

		fmt.Fprintln(writer, saida)
	}
}
