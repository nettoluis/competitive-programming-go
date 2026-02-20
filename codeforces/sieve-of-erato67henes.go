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

	var n int

	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var length int

		fmt.Fscan(reader, &length)
		nums := make([]int, length)
		encontrado := make(map[int]bool)

		for j := 0; j < len(nums); j++ {
			fmt.Fscan(reader, &nums[j])

			if nums[j] == 67 || nums[j] == 1 {
				encontrado[nums[j]] = true
			}

		}

		saida := ""
		switch (encontrado[67] && encontrado[1]) || encontrado[67] {
			case true:
				saida = "YES"
			case false:
				saida = "NO"
		}

		fmt.Fprintln(writer, saida)


	}
}
