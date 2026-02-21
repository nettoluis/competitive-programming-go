package main

import (
	"fmt"
	"bufio"
	"os"
)

const codeforces = "codeforces"

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int

	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var word string

		fmt.Fscan(reader, &word)

		counter := 0
		for j := 0; j < len(word); j++ {
			if word[j] != codeforces[j] {
				counter++
			}
		}
		
		fmt.Println(counter)
	}
}
