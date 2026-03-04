package main

import (
	"fmt"
	"bufio"
	"os"
)

func constroiSaida(a []int) string {
	saida := ""

	for i := range a {
		if i == 0 {
			saida += fmt.Sprintf("%d", a[i])
		} else {
			saida += fmt.Sprintf(" %d", a[i])
		}
	}

	return saida
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var testCases int

	fmt.Fscan(reader, &testCases)

	for t := 0; t < testCases; t++ {
		var n, m, h int

		fmt.Fscan(reader, &n, &m, &h)

		nums := make([]int, n)
		numsOriginal := make([]int, n)
		for i := 0; i < len(nums); i++ {
			fmt.Fscan(reader, &nums[i])
		}
		copy(numsOriginal, nums)

		for k := 0; k < m; k++ {
			var i, value int

			fmt.Fscan(reader, &i, &value)

			if nums[i-1] + value <= h {
				nums[i-1] += value
			} else {
				copy(nums, numsOriginal)
			}
		}

		fmt.Fprintln(writer, constroiSaida(nums))
	}
}
