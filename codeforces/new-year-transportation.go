package main

import "fmt"

func main() {
	var n, t int

	fmt.Scan(&n, &t)
	celulas := make([]int, n)

	for i := 0; i < n - 1; i++ {
		fmt.Scan(&celulas[i])
	}
	
	for pos := 1; pos <= t;  {
		if pos == t {
			fmt.Println("YES")
			return
		}
		pos += celulas[pos-1]
	}

	fmt.Println("NO")
}
