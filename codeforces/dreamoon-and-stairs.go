package main

import (
	"fmt"
)

func ceil(a, b int) int {
	if a % b == 0 {
		return a / b
	} else {
		return (a / b) + 1
	}
}

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	for i := ceil(n, 2); i <= int(n); i++ {
		if i % m == 0 {
		fmt.Println(i)
		return
		}
	}

	fmt.Println(-1)
}
